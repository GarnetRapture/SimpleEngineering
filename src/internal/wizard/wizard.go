package wizard

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"simpleharness/src/internal/blueprint"
	"simpleharness/src/internal/catalog"
	"simpleharness/src/internal/console"
	"simpleharness/src/internal/generate"
	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/install"
	"simpleharness/src/internal/locale"
	"simpleharness/src/internal/system"
)

const outputFolder = "output"

var projectNamePattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*$`)
var skillPrefixPattern = regexp.MustCompile(`^[a-z0-9-]+$`)
var environmentVariablePattern = regexp.MustCompile(`^[A-Z_][A-Z0-9_]*$`)
var prefixInvalidCharacters = regexp.MustCompile(`[^a-z0-9-]+`)

var stepNumbers = []string{"①", "②", "③", "④", "⑤", "⑥", "⑦", "⑧", "⑨"}

var linuxGitCommands = map[string]string{
	"apt":    "sudo apt-get install -y git",
	"dnf":    "sudo dnf install -y git",
	"pacman": "sudo pacman -S git",
	"zypper": "sudo zypper install git",
}

var linuxInstallPatterns = map[string]string{
	"apt":    "sudo apt-get install -y <package>",
	"dnf":    "sudo dnf install -y <package>",
	"pacman": "sudo pacman -S <package>",
	"zypper": "sudo zypper install <package>",
}

func Run(input io.Reader, output, errorOutput io.Writer, template fs.FS) int {
	con := console.New(input, output)
	info := system.Detect()
	profile, err := ask(con, info)
	if errors.Is(err, console.ErrInputEnded) {
		con.Say(locale.InputEnded)
		return 1
	}
	if err != nil {
		fmt.Fprintln(errorOutput, err)
		return 1
	}
	if profile == nil {
		con.Say(locale.Cancelled)
		return 0
	}
	config := blueprint.Build(*profile)
	workingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(errorOutput, err)
		return 1
	}
	destination := filepath.Join(workingDirectory, outputFolder, profile.ProjectName)
	if _, err := os.Stat(destination); err == nil {
		con.Say(locale.OutputExists, destination)
		return 1
	}
	written, err := generate.Write(template, config, destination)
	if err != nil {
		con.Say(locale.GenerateFailed, err)
		return 1
	}
	con.Say(locale.Generated, len(written), destination)
	if err := offerTools(con, info, config, input, output, errorOutput); err != nil {
		if errors.Is(err, console.ErrInputEnded) {
			con.Say(locale.InputEnded)
			return 1
		}
		fmt.Fprintln(errorOutput, err)
		return 1
	}
	guidePath := filepath.Join(workingDirectory, outputFolder, profile.ProjectName+"-SETUP.md")
	guide := nextSteps(con.Language(), info, config, destination)
	con.Heading(locale.StepNext)
	con.Line(guide)
	if err := os.WriteFile(guidePath, []byte(guideDocument(con.Language(), config, guide)), 0o644); err != nil {
		fmt.Fprintln(errorOutput, err)
		return 1
	}
	con.Say(locale.NextGuideSaved, guidePath)
	return 0
}

func ask(con *console.Console, info system.Info) (*blueprint.Profile, error) {
	languageOptions := make([]console.Option, len(locale.Languages))
	for index, language := range locale.Languages {
		languageOptions[index] = console.Option{Label: locale.LanguageNames[language]}
	}
	con.Line("== SimpleHarness ==")
	languageIndex, err := con.Choose(locale.LanguageQuestion, languageOptions, 0)
	if err != nil {
		return nil, err
	}
	language := locale.Languages[languageIndex]
	con.SetLanguage(language)
	con.Heading(locale.Title)
	con.Say(locale.Intro)
	showEnvironment(con, info)

	con.Heading(locale.StepProject)
	name, err := con.Text(locale.ProjectNameQuestion.In(language), "my-project", projectNamePattern.MatchString, locale.ProjectNameInvalid)
	if err != nil {
		return nil, err
	}
	summary, err := con.OptionalText(locale.ProjectSummaryQuestion.In(language))
	if err != nil {
		return nil, err
	}
	if summary == "" {
		summary = locale.ProjectSummaryPlaceholder.Format(language, name)
	}
	prefix, err := con.Text(locale.SkillPrefixQuestion.In(language), suggestedPrefix(name), skillPrefixPattern.MatchString, locale.SkillPrefixInvalid)
	if err != nil {
		return nil, err
	}
	kindOptions := make([]console.Option, len(catalog.Kinds))
	for index, kind := range catalog.Kinds {
		kindOptions[index] = console.Option{Label: kind.Label.In(language)}
	}
	kindIndex, err := con.Choose(locale.KindQuestion.In(language), kindOptions, 0)
	if err != nil {
		return nil, err
	}
	kind := catalog.Kinds[kindIndex].ID
	languageChoices := make([]console.Option, len(catalog.Languages))
	defaultLanguage := -1
	for index, candidate := range catalog.Languages {
		option := console.Option{Label: candidate.Label}
		for _, recommended := range candidate.RecommendedKinds {
			if recommended == kind {
				option.Detail = locale.Recommended.In(language)
				if defaultLanguage < 0 {
					defaultLanguage = index
				}
			}
		}
		languageChoices[index] = option
	}
	projectLanguageIndex, err := con.Choose(locale.ProjectLanguageQuestion.In(language), languageChoices, defaultLanguage)
	if err != nil {
		return nil, err
	}
	projectLanguage := catalog.Languages[projectLanguageIndex].ID

	con.Heading(locale.StepAgents)
	agentOptions := make([]console.Option, len(catalog.MainAgents))
	for index, agent := range catalog.MainAgents {
		agentOptions[index] = console.Option{Label: agent.Label, Detail: agent.Detail.In(language)}
	}
	agentIndex, err := con.Choose(locale.MainAgentQuestion.In(language), agentOptions, 0)
	if err != nil {
		return nil, err
	}
	mainAgent := catalog.MainAgents[agentIndex].ID
	subagents, err := askSubagents(con, language, mainAgent)
	if err != nil {
		return nil, err
	}

	con.Heading(locale.StepRules)
	con.Say(locale.RulesIntro)
	var applicable []catalog.Rule
	for _, rule := range catalog.Rules {
		if rule.AppliesTo(projectLanguage) {
			applicable = append(applicable, rule)
		}
	}
	ruleOptions := make([]console.Option, len(applicable))
	ruleDefaults := make([]bool, len(applicable))
	for index, rule := range applicable {
		tag := locale.RuleSkillTag
		if rule.Hook {
			tag = locale.RuleHookTag
		}
		ruleOptions[index] = console.Option{Label: tag.In(language) + " " + rule.Label.In(language)}
		ruleDefaults[index] = true
	}
	enabled, err := con.ChooseMany("", ruleOptions, ruleDefaults)
	if err != nil {
		return nil, err
	}
	var rules []string
	for index, on := range enabled {
		if on {
			rules = append(rules, applicable[index].ID)
		}
	}

	profile := &blueprint.Profile{
		OutputLanguage:  language,
		ProjectName:     name,
		ProjectSummary:  summary,
		SkillPrefix:     prefix,
		ProjectKind:     kind,
		ProjectLanguage: projectLanguage,
		MainAgent:       mainAgent,
		Subagents:       subagents,
		Rules:           rules,
	}
	con.Heading(locale.StepGenerate)
	showSummary(con, profile)
	confirmed, err := con.Confirm(locale.ConfirmGenerate.In(language), true)
	if err != nil {
		return nil, err
	}
	if !confirmed {
		return nil, nil
	}
	return profile, nil
}

func suggestedPrefix(name string) string {
	prefix := strings.Trim(prefixInvalidCharacters.ReplaceAllString(strings.ToLower(name), "-"), "-")
	if prefix == "" {
		return "sh"
	}
	return prefix
}

func askSubagents(con *console.Console, language locale.Language, mainAgent string) ([]blueprint.SubagentChoice, error) {
	providers := catalog.ProvidersFor(mainAgent)
	options := make([]console.Option, len(providers))
	for index, provider := range providers {
		options[index] = console.Option{Label: provider.Label, Detail: provider.Detail.In(language)}
	}
	selected, err := con.ChooseMany(locale.SubagentQuestion.In(language), options, make([]bool, len(providers)))
	if err != nil {
		return nil, err
	}
	var choices []blueprint.SubagentChoice
	for index, on := range selected {
		if !on {
			continue
		}
		provider := providers[index]
		choice := blueprint.SubagentChoice{ProviderID: provider.ID}
		if len(provider.Regions) > 1 {
			regionOptions := make([]console.Option, len(provider.Regions))
			for regionIndex, region := range provider.Regions {
				regionOptions[regionIndex] = console.Option{Label: region.Label.In(language), Detail: region.BaseURL}
			}
			regionIndex, err := con.Choose(locale.ProviderRegionQuestion.Format(language, provider.Label), regionOptions, 0)
			if err != nil {
				return nil, err
			}
			choice.RegionID = provider.Regions[regionIndex].ID
		} else if len(provider.Regions) == 1 {
			choice.RegionID = provider.Regions[0].ID
		}
		if provider.ModelsURL != "" {
			model, err := con.Text(locale.ProviderModelQuestion.Format(language, provider.Label, provider.ModelsURL), provider.Model, func(value string) bool { return !strings.ContainsAny(value, " \t") }, locale.ProviderModelInvalid)
			if err != nil {
				return nil, err
			}
			choice.Model = model
		}
		if provider.TokenEnv != "" {
			tokenEnv, err := con.Text(locale.ProviderTokenQuestion.Format(language, provider.Label), provider.TokenEnv, environmentVariablePattern.MatchString, locale.EnvironmentVariableInvalid)
			if err != nil {
				return nil, err
			}
			choice.TokenEnv = tokenEnv
		}
		choices = append(choices, choice)
	}
	return choices, nil
}

func showSummary(con *console.Console, profile *blueprint.Profile) {
	language := profile.OutputLanguage
	con.Line(locale.SummaryHeading.In(language))
	kindLabel := ""
	for _, kind := range catalog.Kinds {
		if kind.ID == profile.ProjectKind {
			kindLabel = kind.Label.In(language)
		}
	}
	con.Line("  " + locale.SummaryLanguage.Format(language, locale.LanguageNames[language]))
	con.Line("  " + locale.SummaryProject.Format(language, profile.ProjectName, catalog.LanguageByID(profile.ProjectLanguage).Label, kindLabel))
	con.Line("  " + locale.SummaryPrefix.Format(language, profile.SkillPrefix))
	con.Line("  " + locale.SummaryMain.Format(language, catalog.AgentByID(profile.MainAgent).Label))
	subagents := locale.None.In(language)
	if len(profile.Subagents) > 0 {
		names := make([]string, len(profile.Subagents))
		for index, choice := range profile.Subagents {
			names[index] = catalog.ProviderByID(choice.ProviderID).Label
			if choice.Model != "" {
				names[index] += " (" + choice.Model + ")"
			}
		}
		subagents = strings.Join(names, ", ")
	}
	con.Line("  " + locale.SummarySubs.Format(language, subagents))
	rules := locale.None.In(language)
	if len(profile.Rules) > 0 {
		rules = strings.Join(profile.Rules, ", ")
	}
	con.Line("  " + locale.SummaryRules.Format(language, rules))
	con.Line("  " + locale.SummaryOutput.Format(language, filepath.Join(outputFolder, profile.ProjectName)))
}

func toolMark(installed bool, language locale.Language) string {
	if installed {
		return "✓ " + locale.Installed.In(language)
	}
	return "✗ " + locale.Missing.In(language)
}

func showEnvironment(con *console.Console, info system.Info) {
	language := con.Language()
	con.Heading(locale.EnvHeading)
	switch info.OS {
	case "windows":
		switch {
		case info.IsWindows11():
			con.Say(locale.EnvWindows11, info.WindowsBuild)
		case info.IsWindows10():
			con.Say(locale.EnvWindows10, info.WindowsBuild)
		default:
			con.Say(locale.EnvOther, "Windows")
		}
	case "linux":
		name := info.LinuxName
		if name == "" {
			name = "Linux"
		}
		if info.WSL {
			con.Say(locale.EnvLinuxWSL, name)
		} else {
			con.Say(locale.EnvLinux, name)
		}
	case "darwin":
		con.Say(locale.EnvMac, info.MacVersion)
	default:
		con.Say(locale.EnvOther, runtime.GOOS)
	}
	if info.PackageManager == "" {
		con.Say(locale.EnvNoPackageManager)
	} else {
		con.Say(locale.EnvPackageManager, info.PackageManager)
	}
	hasGit := system.Has("git")
	hasNode := system.Has("node")
	con.Line(fmt.Sprintf("  git  %s", toolMark(hasGit, language)))
	con.Line(fmt.Sprintf("  node %s", toolMark(hasNode, language)))
	if info.OS == "windows" {
		con.Line(fmt.Sprintf("  pwsh %s", toolMark(system.Has("pwsh"), language)))
	}
	step := 0
	next := func() string {
		number := stepNumbers[step]
		step++
		return number
	}
	switch {
	case info.IsWindows10():
		con.Say(locale.GuideWindows10Cmd)
		con.Say(locale.GuideShellChoice)
		con.Say(locale.GuideWindows10Winget)
		step = 1
		con.Say(locale.GuideWindowsPowerShell7, next())
		con.Say(locale.GuideWindows10Terminal)
		step = 3
		if !hasGit {
			con.Say(locale.GuideWindowsGit, next())
		}
		con.Say(locale.GuideWindows10DefaultTerminal, next())
		con.Say(locale.GuideWindowsGitBash)
		con.Say(locale.GuideWindowsExecutionPolicy)
		con.Say(locale.GuideCodexWindows10)
		con.Say(locale.GuideWindowsReopen)
	case info.IsWindows11():
		con.Say(locale.GuideWindows11Shell)
		con.Say(locale.GuideShellChoice)
		con.Say(locale.GuideWindowsPowerShell7, next())
		if !hasGit {
			con.Say(locale.GuideWindowsGit, next())
		}
		con.Say(locale.GuideWindows11DefaultProfile, next())
		con.Say(locale.GuideWindowsGitBash)
		con.Say(locale.GuideWindowsExecutionPolicy)
		con.Say(locale.GuideWindowsReopen)
	case info.OS == "linux":
		if info.WSL {
			con.Say(locale.GuideWSL)
			con.Say(locale.GuideWSLTerminal)
		}
		if command, found := linuxGitCommands[info.PackageManager]; found && !hasGit {
			con.Say(locale.GuideGitLinux, command)
		}
		if pattern, found := linuxInstallPatterns[info.PackageManager]; found {
			con.Say(locale.GuideLinuxInstall, pattern)
		}
		if !info.WSL {
			con.Say(locale.GuideLinuxTerminal)
		}
	case info.OS == "darwin":
		if info.PackageManager == "" {
			con.Say(locale.GuideMacBrew)
		}
		if !hasGit {
			con.Say(locale.GuideGitMac)
		}
		con.Say(locale.GuideMacShell)
	}
	if !hasNode {
		con.Say(locale.GuideNode)
	}
	con.Say(locale.GuideRunModes)
}

func offerTools(con *console.Console, info system.Info, config harness.Config, input io.Reader, output, errorOutput io.Writer) error {
	language := con.Language()
	con.Heading(locale.StepTools)
	con.Say(locale.ToolsIntro)
	items := install.Plan(info, config)
	var missing []install.Item
	for _, item := range items {
		con.Line(fmt.Sprintf("  %s  %-34s %s", toolMark(item.Installed, language), item.Tool.Label, item.Tool.Purpose.In(language)))
		if !item.Installed {
			missing = append(missing, item)
		}
	}
	if len(missing) == 0 {
		con.Say(locale.ToolsAllInstalled)
		return nil
	}
	var automatic, manual []install.Item
	for _, item := range missing {
		if item.Command != nil {
			automatic = append(automatic, item)
		} else {
			manual = append(manual, item)
		}
	}
	if len(automatic) > 0 {
		confirmed, err := con.Confirm(locale.ToolsInstallQuestion.Format(language, len(automatic), install.ManagerLabel(info.PackageManager)), false)
		if err != nil {
			return err
		}
		if confirmed {
			if err := install.Prepare(info.PackageManager, input, output, errorOutput); err != nil {
				con.Say(locale.ToolFailed, "apt-get update", err)
			}
			for _, item := range automatic {
				con.Say(locale.ToolInstalling, item.Tool.Label)
				if err := install.Run(item, input, output, errorOutput); err != nil {
					con.Say(locale.ToolFailed, item.Tool.Label, err)
					continue
				}
				con.Say(locale.ToolInstalled, item.Tool.Label)
			}
			con.Say(locale.ToolsRestartShell)
		} else {
			con.Say(locale.ToolsSkipped)
			for _, item := range automatic {
				con.Line("  " + strings.Join(item.Command, " "))
			}
		}
	} else if !install.Supported(info.PackageManager) {
		con.Say(locale.ToolsNoManager)
	}
	for _, item := range manual {
		con.Say(locale.ToolManual, item.Tool.Label, item.Manual)
	}
	return nil
}
