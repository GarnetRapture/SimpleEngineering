package blueprint

import (
	"simpleharness/src/internal/catalog"
	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/locale"
)

type SubagentChoice struct {
	ProviderID string
	RegionID   string
	Model      string
	TokenEnv   string
}

type Profile struct {
	OutputLanguage  locale.Language
	ProjectName     string
	ProjectSummary  string
	SkillPrefix     string
	ProjectKind     string
	ProjectLanguage string
	MainAgent       string
	Subagents       []SubagentChoice
	Rules           []string
}

const (
	ScratchPath          = "tmp-agent"
	HookPath             = ".harness/hooks/run_project_hook.mjs"
	SubagentLauncherPath = ".harness/subagents/run_subagent.mjs"
)

var timeLabel = locale.Text{Ko: "현지 시각", En: "local time", Zh: "本地时间"}

const timeCommand = `node -e "const zone=Intl.DateTimeFormat().resolvedOptions().timeZone;process.stdout.write(new Intl.DateTimeFormat('sv-SE',{timeZone:zone,dateStyle:'short',timeStyle:'medium'}).format(new Date())+' '+zone+'\n')"`

var versionCommandUndefined = locale.Text{
	Ko: "(프로젝트 정의: 제품 버전을 바꾸는 명령)",
	En: "(project definition: the command that changes the product version)",
	Zh: "（项目定义：修改产品版本的命令）",
}

var subagentTrigger = "a task is delegated to a subagent or a subagent result is about to be used"
var namingTrigger = "a file of the primary language is created, renamed, or edited, or a symbol is named"
var inlineTrigger = "a file of the primary language is edited"

var namingRules = map[string]bool{"domain-structure": true, "domain-naming": true}

func Build(profile Profile) harness.Config {
	language := catalog.LanguageByID(profile.ProjectLanguage)
	agent := catalog.AgentByID(profile.MainAgent)
	stackExtensions := stackExtensions(language.Stack)
	versionCommand := language.VersionCommand
	if versionCommand == "" {
		versionCommand = versionCommandUndefined.In(profile.OutputLanguage)
	}
	config := harness.Config{
		Project: harness.Project{
			Name:     profile.ProjectName,
			Summary:  profile.ProjectSummary,
			Kind:     profile.ProjectKind,
			Language: profile.ProjectLanguage,
		},
		Output: harness.Output{Language: string(profile.OutputLanguage)},
		Agents: harness.Agents{Main: profile.MainAgent},
		Skills: harness.Skills{Prefix: profile.SkillPrefix, Root: agent.SkillRoot, Domain: []harness.DomainSkill{}},
		Paths: harness.Paths{
			Conventions:      "docs/conventions.md",
			Patterns:         "docs/agent-patterns.md",
			Tracking:         "docs/ai-tracking.md",
			Audit:            "docs/session-audit.md",
			Feedback:         "docs/user-feedback.md",
			Scratch:          ScratchPath,
			CommitGuard:      "tools/guard/staged_guard.mjs",
			Hook:             HookPath,
			SubagentLauncher: SubagentLauncherPath,
		},
		Time:           harness.Time{Label: timeLabel.In(profile.OutputLanguage), Command: timeCommand},
		Version:        harness.Version{Source: language.VersionSource, Command: versionCommand},
		Stacks:         catalog.Stacks,
		Rules:          append([]string{}, profile.Rules...),
		InlineRules:    []harness.InlineRule{},
		Observations:   []harness.Observation{},
		ProtectedPaths: []harness.ProtectedPath{},
		ForbiddenPaths: []harness.PatternReason{},
		DeniedCommands: []harness.PatternReason{},
		MCP: harness.MCP{
			QueryRoute:            "the registered tools of the C12 table",
			ExecutingToolPrefixes: []string{},
			HumanInputTools:       []string{},
			ObservationTools:      []string{},
			Servers:               []harness.MCPServer{},
		},
		Subagents:   []harness.Subagent{},
		Claude:      harness.ClaudeSettings{PermissionsAllow: []string{}},
		CommitGuard: harness.CommitGuard{BlockedExtensions: []string{"zip", "7z", "rar", "tar", "gz", "exe", "dll", "so", "dylib", "sqlite", "db", "mp4", "mov", "wav", "psd"}},
		Response:    catalog.Responses[profile.OutputLanguage],
	}
	hasNaming, hasInline := false, false
	for _, id := range profile.Rules {
		rule := catalog.RuleByID(id)
		config.InlineRules = append(config.InlineRules, rule.Inline...)
		if namingRules[id] {
			hasNaming = true
		} else {
			hasInline = true
		}
	}
	if hasNaming {
		config.Skills.Domain = append(config.Skills.Domain, harness.DomainSkill{Name: "naming", Trigger: namingTrigger, Paths: []string{}, Commands: []string{}, ToolPrefixes: []string{}, Extensions: stackExtensions})
	}
	if hasInline {
		config.Skills.Domain = append(config.Skills.Domain, harness.DomainSkill{Name: "inline", Trigger: inlineTrigger, Paths: []string{}, Commands: []string{}, ToolPrefixes: []string{}, Extensions: stackExtensions})
	}
	for _, protected := range append(append([]catalog.ProtectedPath{}, catalog.CommonProtected...), language.Protected...) {
		config.ProtectedPaths = append(config.ProtectedPaths, harness.ProtectedPath{
			Pattern:       protected.Pattern,
			ShellSegments: protected.ShellSegments,
			Target:        protected.Target.In(profile.OutputLanguage),
			Reason:        protected.Reason,
			Route:         protected.Route.In(profile.OutputLanguage),
		})
	}
	for _, choice := range profile.Subagents {
		config.Subagents = append(config.Subagents, buildSubagent(profile.MainAgent, choice))
	}
	if len(config.Subagents) > 0 {
		config.Skills.Domain = append(config.Skills.Domain, harness.DomainSkill{Name: "subagent", Trigger: subagentTrigger, Paths: []string{}, Commands: []string{SubagentLauncherPath}, ToolPrefixes: []string{}, Extensions: []string{}})
	}
	return config
}

func stackExtensions(stackName string) []string {
	for _, stack := range catalog.Stacks {
		if stack.Name == stackName {
			return append([]string{}, stack.Extensions...)
		}
	}
	panic("blueprint: unknown stack " + stackName)
}

func buildSubagent(mainAgent string, choice SubagentChoice) harness.Subagent {
	provider := catalog.ProviderByID(choice.ProviderID)
	switch provider.Runtime {
	case catalog.RuntimeCodex:
		return harness.Subagent{ID: provider.ID, Provider: provider.ID, Runtime: catalog.RuntimeCodex}
	case catalog.RuntimeOllama:
		if mainAgent == "codex" {
			return harness.Subagent{ID: provider.ID, Provider: provider.ID, Runtime: catalog.RuntimeCodex, LocalProvider: "ollama", Model: choice.Model}
		}
		return harness.Subagent{ID: provider.ID, Provider: provider.ID, Runtime: catalog.RuntimeClaude, BaseURL: provider.Regions[0].BaseURL, AuthVariable: "ANTHROPIC_AUTH_TOKEN", TokenValue: "ollama", Model: choice.Model}
	case catalog.RuntimeClaude:
		if len(provider.Regions) == 0 {
			return harness.Subagent{ID: provider.ID, Provider: provider.ID, Runtime: catalog.RuntimeClaude}
		}
		return harness.Subagent{
			ID:           provider.ID,
			Provider:     provider.ID,
			Runtime:      catalog.RuntimeClaude,
			BaseURL:      regionURL(provider, choice.RegionID),
			AuthVariable: provider.AuthVariable,
			TokenEnv:     choice.TokenEnv,
			Model:        choice.Model,
		}
	}
	panic("blueprint: unknown subagent runtime " + provider.Runtime)
}

func regionURL(provider catalog.Provider, regionID string) string {
	for _, region := range provider.Regions {
		if region.ID == regionID {
			return region.BaseURL
		}
	}
	panic("blueprint: unknown region " + regionID + " for " + provider.ID)
}
