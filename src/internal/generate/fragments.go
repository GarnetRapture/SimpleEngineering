package generate

import (
	"fmt"
	"io/fs"
	"strings"

	"simpleharness/src/internal/blueprint"
	"simpleharness/src/internal/catalog"
	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/locale"
	"simpleharness/src/internal/render"
)

var domainSkillLinks = map[string]string{
	"naming":   "namingSkill",
	"inline":   "inlineSkill",
	"subagent": "subagentSkill",
}

var domainSkillPurpose = map[string]locale.Text{
	"naming": {
		Ko: "C04·C05 도메인 구조·명명 규칙의 판정 절차와 좋은·나쁜 예시",
		En: "judgment procedure and good and bad examples of the C04 and C05 domain structure and naming rules",
		Zh: "C04、C05 领域结构与命名规则的判定流程及正反示例",
	},
	"inline": {
		Ko: "C04·C05 인라인 금지 규칙의 판정 절차와 좋은·나쁜 예시",
		En: "judgment procedure and good and bad examples of the C04 and C05 inline rules",
		Zh: "C04、C05 内联禁止规则的判定流程及正反示例",
	},
	"subagent": {
		Ko: "C12 서브 에이전트 위임·결과 검증 절차",
		En: "C12 subagent delegation and result verification procedure",
		Zh: "C12 子代理委派与结果验证流程",
	},
}

var englishLanguage = locale.Text{Ko: "영어", En: "English", Zh: "英语"}
var noneText = locale.Text{Ko: "없음", En: "none", Zh: "无"}
var handEdit = locale.Text{Ko: "손 편집·직접 변경", En: "hand edits and direct changes", Zh: "手动编辑与直接修改"}
var alwaysActive = locale.Text{Ko: "(항상 활성)", En: "(always active)", Zh: "（始终启用）"}
var hookEnforced = locale.Text{Ko: "훅이 편집 차단", En: "blocked by the hook at edit time", Zh: "钩子在编辑时拦截"}
var skillEnforced = locale.Text{Ko: "스킬·리뷰 절차로 판정", En: "judged through the skill and review procedure", Zh: "通过技能与评审流程判定"}
var observationSuffix = locale.Text{Ko: ", C12의 승인된 관찰 경로", En: ", and the authorized observation route of C12", Zh: "，以及 C12 授权的观察路径"}
var noChecks = locale.Text{Ko: "C12 스택 표에 등록된 검사와 C12의 승인된 관찰 경로", En: "the checks the C12 stack table registers, and the authorized observation route of C12", Zh: "C12 栈表登记的检查，以及 C12 授权的观察路径"}
var contextSourceDoc = locale.Text{
	Ko: "컨벤션 요약 표와 AGENTS 행동 요약의 원문 위치·해시",
	En: "the conventions summary table and the AGENTS behavior summary with their source locations and hashes",
	Zh: "约定摘要表与 AGENTS 行为摘要的原文位置与哈希",
}
var localOllama = locale.Text{Ko: "이 PC의 Ollama", En: "Ollama on this PC", Zh: "本机 Ollama"}
var userAccount = locale.Text{Ko: "사용자 계정 기본 설정", En: "the user's own account settings", Zh: "用户自己的账号设置"}

var hookRegistrationDoc = map[string]locale.Text{
	"claude": {
		Ko: "`.claude/settings.json`이 SessionStart·UserPromptSubmit·PreToolUse·PostToolUse·PostToolUseFailure·Stop에 등록한 `node -e … claude` 명령",
		En: "the `node -e … claude` command that `.claude/settings.json` registers on SessionStart, UserPromptSubmit, PreToolUse, PostToolUse, PostToolUseFailure, and Stop",
		Zh: "`.claude/settings.json` 在 SessionStart、UserPromptSubmit、PreToolUse、PostToolUse、PostToolUseFailure、Stop 上注册的 `node -e … claude` 命令",
	},
	"codex": {
		Ko: "`.codex/hooks.json`이 SessionStart·UserPromptSubmit·PreToolUse·PostToolUse·Stop에 등록한 `node -e … codex` 명령 (Codex는 프로젝트를 신뢰하고 `/hooks`에서 검토한 뒤 실행한다)",
		En: "the `node -e … codex` command that `.codex/hooks.json` registers on SessionStart, UserPromptSubmit, PreToolUse, PostToolUse, and Stop (Codex runs it once the project is trusted and the hooks are reviewed in `/hooks`)",
		Zh: "`.codex/hooks.json` 在 SessionStart、UserPromptSubmit、PreToolUse、PostToolUse、Stop 上注册的 `node -e … codex` 命令（Codex 在项目受信任并于 `/hooks` 中审查后运行）",
	},
}

func buildFragments(template fs.FS, config harness.Config, language locale.Language) (map[string]string, error) {
	prefix := config.Skills.Prefix
	fragments := map[string]string{
		"contextSource":         "the conventions summary table and the AGENTS behavior summary with their source identities",
		"contextSourceDoc":      contextSourceDoc.In(language),
		"skillLoadMethod":       skillLoadMethod(config),
		"skillRows":             skillRows(config),
		"skillOwnershipRows":    skillOwnershipRows(config, language),
		"checkTable":            checkTable(config),
		"checkCommands":         checkCommands(config, locale.English),
		"checkCommandsDoc":      checkCommands(config, language),
		"responseFieldList":     render.Join(config.Response.Fields, ", ", func(field string) string { return render.Code(field + ":") }),
		"basisLabelList":        render.Join(config.Response.BasisLabels, ", ", func(label string) string { return render.Code(label + ":") }),
		"mcpServers":            mcpServers(config, locale.English),
		"mcpServersDoc":         mcpServers(config, language),
		"stackRows":             stackRows(config, language),
		"stackSkillTable":       stackSkillTable(config),
		"protectedRows":         protectedRows(config, language),
		"mcpToolTable":          mcpToolTable(config),
		"harnessRows":           harnessRows(config),
		"hookRegistration":      hookRegistrationDoc[config.Agents.Main].In(locale.English),
		"hookRegistrationDoc":   hookRegistrationDoc[config.Agents.Main].In(language),
		"subagentRows":          subagentRows(config, language),
		"subagentTable":         subagentRows(config, locale.English),
		"ruleRows":              ruleRows(config, language),
		"inlineRuleTable":       inlineRuleTable(config),
		"outputLanguageName":    locale.LanguageNames[language],
		"outputLanguageEnglish": outputLanguageEnglish[language],
		"mainAgentName":         catalog.AgentByID(config.Agents.Main).Label,
		"primaryLanguage":       catalog.LanguageByID(config.Project.Language).Label,
		"projectKindDoc":        projectKind(config.Project.Kind).In(language),
		"skillPrefix":           prefix,
		"nodeHeading":           config.Response.NodeHeading,
		"trackingHeading":       config.Response.TrackingHeading,
		"todoHeading":           config.Response.TodoHeading,
		"openStatus":            config.Response.OpenStatus,
	}
	enabled := map[string]bool{}
	for _, id := range config.Rules {
		enabled[id] = true
	}
	for _, kind := range []string{"naming", "inline"} {
		key := kind + "Examples"
		if !hasDomainSkill(config, kind) {
			fragments[key] = ""
			continue
		}
		content, err := fs.ReadFile(template, "skills/"+kind+"/examples/"+config.Project.Language+".md")
		if err != nil {
			return nil, fmt.Errorf("examples for %s: %w", kind, err)
		}
		fragments[key] = enabledExamples(string(content), enabled)
	}
	return fragments, nil
}

func enabledExamples(content string, enabled map[string]bool) string {
	var kept []string
	keep := false
	for _, line := range strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n") {
		if id, found := strings.CutPrefix(line, "@rule "); found {
			keep = enabled[strings.TrimSpace(id)]
			continue
		}
		if keep {
			kept = append(kept, line)
		}
	}
	return strings.TrimSpace(strings.Join(kept, "\n"))
}

var outputLanguageEnglish = map[locale.Language]string{
	locale.Korean:  "Korean",
	locale.English: "English",
	locale.Chinese: "Simplified Chinese",
}

func projectKind(id string) locale.Text {
	for _, kind := range catalog.Kinds {
		if kind.ID == id {
			return kind.Label
		}
	}
	panic("generate: unknown project kind " + id)
}

func skillLoadMethod(config harness.Config) string {
	if config.Agents.Main == "codex" {
		return "by reading its `SKILL.md` file under `" + config.Skills.Root + "/`"
	}
	return "through the Skill tool"
}

func skillLink(config harness.Config, name, linkID string) string {
	return "[" + config.Skills.Prefix + "-" + name + "]({{link." + linkID + "}})"
}

func skillRows(config harness.Config) string {
	rows := []string{
		skillLink(config, "workflow", "workflowSkill") + " (shared workflow G1–G10 and the change decision D1–D7), " +
			skillLink(config, "fix-process", "fixSkill") + " (C06 diagnostic, FIX, and exception processes), and " +
			skillLink(config, "verify", "verifySkill") + " (C14 logic judgment and fixed checks) before any change, any run of project code or a tool, and any final response",
		skillLink(config, "structure", "structureSkill") + " when a file is added, moved, or renamed",
		skillLink(config, "stack", "stackSkill") + " when a file of an active stack changes or a stack check runs",
		skillLink(config, "mcp", "mcpSkill") + " when a registered MCP server is queried",
	}
	for _, skill := range config.Skills.Domain {
		rows = append(rows, skillLink(config, skill.Name, domainSkillLinks[skill.Name])+" when "+skill.Trigger)
	}
	return strings.Join(rows, "; ") + "."
}

func skillOwnershipRows(config harness.Config, language locale.Language) string {
	rows := make([]string, 0, len(config.Skills.Domain))
	for _, skill := range config.Skills.Domain {
		rows = append(rows, "| "+skillLink(config, skill.Name, domainSkillLinks[skill.Name])+" | "+render.Cell(domainSkillPurpose[skill.Name].In(language))+" | "+englishLanguage.In(language)+" |")
	}
	return strings.Join(rows, "\n")
}

func acceptance(value harness.Acceptance, language locale.Language) string {
	exit := locale.Text{Ko: "종료 %d", En: "exit %d", Zh: "退出码 %d"}.Format(language, value.Exit)
	parts := []string{exit}
	if value.Line != "" {
		parts = append(parts, locale.Text{Ko: "줄 %s 존재", En: "line %s present", Zh: "存在行 %s"}.Format(language, render.Code(value.Line)))
	}
	if value.ForbidLine != "" {
		parts = append(parts, locale.Text{Ko: "%s와 맞는 줄 없음", En: "no line matching %s", Zh: "无匹配 %s 的行"}.Format(language, render.Code(value.ForbidLine)))
	}
	if len(value.Counts) > 0 {
		parts = append(parts, locale.Text{Ko: "%s 0", En: "%s equal 0", Zh: "%s 为 0"}.Format(language, render.Join(value.Counts, ", ", render.Code)))
	}
	return strings.Join(parts, ", ")
}

func stackActivation(stack harness.Stack) string {
	if len(stack.Markers) == 0 {
		return "always active"
	}
	return "when " + render.Join(stack.Markers, ", ", render.Code) + " exist"
}

func checkTable(config harness.Config) string {
	rows := make([]string, 0, len(config.Stacks))
	for _, stack := range config.Stacks {
		checks := "No check is configured; the C12 stack row defines one before acceptance"
		if len(stack.Checks) > 0 {
			parts := make([]string, len(stack.Checks))
			for index, check := range stack.Checks {
				parts[index] = render.Code(check.Command) + " (" + acceptance(check.Acceptance, locale.English) + ")"
			}
			checks = strings.Join(parts, "; ") + ", with error 0 and warn 0"
		}
		rows = append(rows, "   | "+render.Cell(stack.Name+" files ("+strings.Join(stack.Extensions, ", ")+"), "+stackActivation(stack))+" | "+render.Cell(checks)+" |")
	}
	return strings.Join(rows, "\n")
}

func checkCommands(config harness.Config, language locale.Language) string {
	var commands []string
	for _, stack := range config.Stacks {
		for _, check := range stack.Checks {
			commands = append(commands, render.Code(check.Command)+" ("+stack.Name+")")
		}
	}
	if len(commands) == 0 {
		return noChecks.In(language)
	}
	return strings.Join(commands, ", ") + observationSuffix.In(language)
}

func mcpServers(config harness.Config, language locale.Language) string {
	if len(config.MCP.Servers) == 0 {
		if language == locale.English {
			return "none registered"
		}
		return noneText.In(language)
	}
	names := make([]string, len(config.MCP.Servers))
	for index, server := range config.MCP.Servers {
		names[index] = render.Code(server.Name)
	}
	return strings.Join(names, ", ")
}

func stackRows(config harness.Config, language locale.Language) string {
	rows := make([]string, 0, len(config.Stacks))
	for _, stack := range config.Stacks {
		markers := alwaysActive.In(language)
		if len(stack.Markers) > 0 {
			markers = render.Join(stack.Markers, "·", render.Code)
		}
		checks := noneText.In(language)
		if len(stack.Checks) > 0 {
			parts := make([]string, len(stack.Checks))
			for index, check := range stack.Checks {
				parts[index] = render.Code(check.Command) + " — " + acceptance(check.Acceptance, language)
			}
			checks = strings.Join(parts, "<br>")
		}
		forbidden := noneText.In(language)
		if len(stack.Forbidden) > 0 {
			parts := make([]string, len(stack.Forbidden))
			for index, entry := range stack.Forbidden {
				parts[index] = render.Code(entry.Command) + " — " + entry.Reason
			}
			forbidden = strings.Join(parts, "<br>")
		}
		rows = append(rows, render.Row(stack.Name, markers, render.Join(stack.Extensions, " ", render.Code), checks, forbidden))
	}
	return strings.Join(rows, "\n")
}

func stackSkillTable(config harness.Config) string {
	if len(config.Stacks) == 0 {
		return "No stack profile is configured in `harness.config.json` `stacks`."
	}
	rows := []string{
		"| Stack | Markers (all must exist) | Extensions | Allowed checks and acceptance | Forbidden as verification |",
		"| --- | --- | --- | --- | --- |",
	}
	for _, stack := range config.Stacks {
		markers := "(always active)"
		if len(stack.Markers) > 0 {
			markers = render.Join(stack.Markers, ", ", render.Code)
		}
		checks := "(none)"
		if len(stack.Checks) > 0 {
			parts := make([]string, len(stack.Checks))
			for index, check := range stack.Checks {
				parts[index] = render.Code(check.Command) + " — " + acceptance(check.Acceptance, locale.English)
			}
			checks = strings.Join(parts, "<br>")
		}
		forbidden := "(none)"
		if len(stack.Forbidden) > 0 {
			parts := make([]string, len(stack.Forbidden))
			for index, entry := range stack.Forbidden {
				parts[index] = render.Code(entry.Command) + " — " + entry.Reason
			}
			forbidden = strings.Join(parts, "<br>")
		}
		rows = append(rows, render.Row(stack.Name, markers, render.Join(stack.Extensions, " ", render.Code), checks, forbidden))
	}
	return strings.Join(rows, "\n")
}

func protectedRows(config harness.Config, language locale.Language) string {
	rows := make([]string, len(config.ProtectedPaths))
	for index, entry := range config.ProtectedPaths {
		rows[index] = "| " + render.Cell(entry.Target) + " | " + handEdit.In(language) + " | " + render.Cell(entry.Route) + " |"
	}
	return strings.Join(rows, "\n")
}

func toolName(value *string) string {
	if value == nil {
		return "not registered"
	}
	return render.Code(*value)
}

func toolList(values []string) string {
	if len(values) == 0 {
		return "none"
	}
	return render.Join(values, ", ", render.Code)
}

func mcpToolTable(config harness.Config) string {
	runner := "not registered"
	if config.MCP.ToolRunner != nil {
		capabilities := make([]string, 0, len(config.MCP.ToolRunner.Capabilities))
		for capability := range config.MCP.ToolRunner.Capabilities {
			capabilities = append(capabilities, render.Code(capability))
		}
		runner = render.Code(config.MCP.ToolRunner.Tool) + " (capabilities: " + strings.Join(capabilities, ", ") + ")"
	}
	return strings.Join([]string{
		render.Row("Context tool", toolName(config.MCP.ContextTool)),
		render.Row("Inventory tool", toolName(config.MCP.InventoryTool)),
		render.Row("Batch query tool", toolName(config.MCP.BatchTool)),
		render.Row("Canonical entry runner", runner),
		render.Row("Tool prefixes that execute project code", toolList(config.MCP.ExecutingToolPrefixes)),
		render.Row("Tools that send input to a running session", toolList(config.MCP.HumanInputTools)),
		render.Row("Observation tools", toolList(config.MCP.ObservationTools)),
	}, "\n")
}

func harnessRows(config harness.Config) string {
	var rows []string
	switch config.Agents.Main {
	case "claude":
		rows = append(rows, render.Row("Claude Code (main)", "`CLAUDE.md` imports this file with `@AGENTS.md`", "Skills in `"+config.Skills.Root+"/` load through the Skill tool; `.claude/settings.json` runs "+render.Code(blueprint.HookPath+" claude")+" on SessionStart, UserPromptSubmit, PreToolUse, PostToolUse, PostToolUseFailure, and Stop (C16)"))
	case "codex":
		rows = append(rows, render.Row("Codex (main)", "Codex reads this file as the repository `AGENTS.md`", "Skills in `"+config.Skills.Root+"/` load by reading each `SKILL.md`; `.codex/hooks.json` runs "+render.Code(blueprint.HookPath+" codex")+" on SessionStart, UserPromptSubmit, PreToolUse, PostToolUse, and Stop once the project is trusted and the hooks are reviewed in `/hooks` (C16)"))
	}
	for _, subagent := range config.Subagents {
		entry := "Reads `AGENTS.md` natively"
		if subagent.Runtime == "claude" {
			entry = "Reads `CLAUDE.md`, which imports this file"
		}
		rows = append(rows, render.Row(subagent.ID+" (subagent, runtime "+subagent.Runtime+")", entry, "Started only by the main agent through "+render.Code("node "+blueprint.SubagentLauncherPath+" "+subagent.ID+" <task file>")+" for read-only delegation; the main agent verifies every result before use (C12 subagent row)"))
	}
	return strings.Join(rows, "\n")
}

func subagentRows(config harness.Config, language locale.Language) string {
	if len(config.Subagents) == 0 {
		return render.Row(noneText.In(language), "", "", "", "")
	}
	rows := make([]string, 0, len(config.Subagents))
	for _, subagent := range config.Subagents {
		endpoint := userAccount.In(language)
		switch {
		case subagent.LocalProvider != "":
			endpoint = localOllama.In(language) + " (`--oss --local-provider " + subagent.LocalProvider + "`)"
		case subagent.BaseURL != "":
			endpoint = render.Code(subagent.BaseURL)
		}
		model := subagent.Model
		if model == "" {
			model = userAccount.In(language)
		} else {
			model = render.Code(model)
		}
		key := noneText.In(language)
		switch {
		case subagent.TokenEnv != "":
			key = render.Code(subagent.TokenEnv) + " → " + render.Code(subagent.AuthVariable)
		case subagent.TokenValue != "":
			key = render.Code(subagent.AuthVariable + "=" + subagent.TokenValue)
		}
		rows = append(rows, render.Row(render.Code(subagent.ID), subagent.Runtime, endpoint, model, key))
	}
	return strings.Join(rows, "\n")
}

func ruleRows(config harness.Config, language locale.Language) string {
	if len(config.Rules) == 0 {
		return render.Row(noneText.In(language), "", "", "")
	}
	rows := make([]string, 0, len(config.Rules))
	for _, id := range config.Rules {
		rule := catalog.RuleByID(id)
		enforcement := skillEnforced.In(language)
		if rule.Hook {
			enforcement = hookEnforced.In(language)
		}
		rows = append(rows, render.Row(render.Code(rule.ID), rule.Clause, rule.Definition.In(language), enforcement))
	}
	return strings.Join(rows, "\n")
}

func inlineRuleTable(config harness.Config) string {
	rows := []string{
		"| Rule | Clause | Extensions | Added text the hook counts | Allowed path |",
		"| --- | --- | --- | --- | --- |",
	}
	for _, id := range config.Rules {
		if id == catalog.RuleNoInlineComment {
			rows = append(rows, render.Row(render.Code(id), "C04", "every authored code extension", "line, block, and documentation comments", "—"))
		}
	}
	for _, rule := range config.InlineRules {
		allowed := "—"
		if rule.AllowPath != "" {
			allowed = render.Code(rule.AllowPath)
		}
		rows = append(rows, render.Row(render.Code(rule.ID), rule.Clause, render.Join(rule.Extensions, " ", render.Code), render.Code(rule.Pattern), allowed))
	}
	return strings.Join(rows, "\n")
}
