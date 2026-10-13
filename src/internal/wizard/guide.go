package wizard

import (
	"fmt"
	"strings"

	"simpleharness/src/internal/blueprint"
	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/locale"
	"simpleharness/src/internal/system"
)

func nextSteps(language locale.Language, info system.Info, config harness.Config, destination string) string {
	var steps []string
	add := func(text string) {
		steps = append(steps, fmt.Sprintf("%d. %s", len(steps)+1, text))
	}
	add(locale.NextCopy.Format(language, destination))
	add(locale.NextGitHooks.In(language))
	add(locale.NextGitignore.Format(language, config.Paths.Scratch))
	add(locale.NextConventions.Format(language, config.Paths.Conventions))
	switch config.Agents.Main {
	case "claude":
		add(locale.NextClaude.In(language))
	case "codex":
		add(locale.NextCodex.In(language))
	}
	var keyLines []string
	for _, subagent := range config.Subagents {
		switch {
		case subagent.TokenEnv != "" && info.OS == "windows":
			keyLines = append(keyLines, fmt.Sprintf("   PowerShell: [Environment]::SetEnvironmentVariable('%s', '<key>', 'User')", subagent.TokenEnv))
		case subagent.TokenEnv != "" && info.OS == "darwin":
			keyLines = append(keyLines, fmt.Sprintf("   zsh: echo 'export %s=<key>' >> ~/.zshrc", subagent.TokenEnv))
		case subagent.TokenEnv != "":
			keyLines = append(keyLines, fmt.Sprintf("   bash: echo 'export %s=<key>' >> ~/.bashrc", subagent.TokenEnv))
		case subagent.Provider == "ollama-local" && subagent.Model != "":
			keyLines = append(keyLines, "   ollama pull "+subagent.Model)
		}
	}
	if len(keyLines) > 0 {
		add(locale.NextSubagentKeys.In(language) + "\n" + strings.Join(keyLines, "\n"))
	}
	if len(config.Subagents) > 0 {
		add(locale.NextSubagentRun.Format(language, config.Paths.Scratch))
	}
	add(locale.NextStructure.In(language))
	add(locale.GuideRunModes.In(language))
	return strings.Join(steps, "\n")
}

func guideDocument(language locale.Language, config harness.Config, guide string) string {
	return "# " + locale.SetupGuideTitle.Format(language, config.Project.Name) + "\n\n" +
		guide + "\n\n" +
		"- " + blueprint.HookPath + "\n" +
		"- " + config.Paths.Conventions + "\n"
}
