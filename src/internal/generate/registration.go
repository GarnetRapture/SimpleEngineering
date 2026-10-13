package generate

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	"simpleharness/src/internal/blueprint"
	"simpleharness/src/internal/harness"
)

const hookTimeoutSeconds = 30

const claudeRootExpression = `process.env.CLAUDE_PROJECT_DIR || '.'`
const codexRootExpression = `require('node:child_process').execFileSync('git', ['rev-parse', '--show-toplevel'], { encoding: 'utf8' }).trim()`

type hookHandler struct {
	Type    string `json:"type"`
	Command string `json:"command"`
	Timeout int    `json:"timeout"`
}

type hookGroup struct {
	Matcher string        `json:"matcher,omitempty"`
	Hooks   []hookHandler `json:"hooks"`
}

type claudeHookEvents struct {
	SessionStart       []hookGroup `json:"SessionStart"`
	UserPromptSubmit   []hookGroup `json:"UserPromptSubmit"`
	PreToolUse         []hookGroup `json:"PreToolUse"`
	PostToolUse        []hookGroup `json:"PostToolUse"`
	PostToolUseFailure []hookGroup `json:"PostToolUseFailure"`
	Stop               []hookGroup `json:"Stop"`
}

type claudePermissions struct {
	Allow []string `json:"allow"`
}

type claudeSettingsFile struct {
	Permissions           claudePermissions `json:"permissions"`
	EnabledMcpjsonServers []string          `json:"enabledMcpjsonServers"`
	Hooks                 claudeHookEvents  `json:"hooks"`
}

type codexHookEvents struct {
	SessionStart     []hookGroup `json:"SessionStart"`
	UserPromptSubmit []hookGroup `json:"UserPromptSubmit"`
	PreToolUse       []hookGroup `json:"PreToolUse"`
	PostToolUse      []hookGroup `json:"PostToolUse"`
	Stop             []hookGroup `json:"Stop"`
}

type codexHookFile struct {
	Hooks codexHookEvents `json:"hooks"`
}

func HookCommand(harnessName string) string {
	root := claudeRootExpression
	if harnessName == "codex" {
		root = codexRootExpression
	}
	return `node -e "import(require('node:url').pathToFileURL(require('node:path').resolve(` + root + `, '` + blueprint.HookPath + `')).href).catch(error => { process.exitCode = 1; process.stderr.write(String(error.stack ?? error)); })" ` + harnessName
}

func generatedTargets(config harness.Config) map[string]string {
	targets := map[string]string{"config": "harness.config.json"}
	switch config.Agents.Main {
	case "claude":
		targets["claudeSettings"] = ".claude/settings.json"
		if len(serversFor(config, "claude")) > 0 {
			targets["mcpJson"] = ".mcp.json"
		}
	case "codex":
		targets["codexHooks"] = ".codex/hooks.json"
		if len(serversFor(config, "codex")) > 0 {
			targets["codexConfig"] = ".codex/config.toml"
		}
	}
	return targets
}

func serversFor(config harness.Config, harnessName string) []harness.MCPServer {
	var servers []harness.MCPServer
	for _, server := range config.MCP.Servers {
		for _, name := range server.Harnesses {
			if name == harnessName {
				servers = append(servers, server)
				break
			}
		}
	}
	return servers
}

func encodeJSON(value any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func registrationFiles(config harness.Config) ([]output, error) {
	targets := generatedTargets(config)
	configBytes, err := encodeJSON(config)
	if err != nil {
		return nil, err
	}
	outputs := []output{{Target: targets["config"], Content: configBytes}}
	handler := []hookHandler{{Type: "command", Command: HookCommand(config.Agents.Main), Timeout: hookTimeoutSeconds}}
	every := []hookGroup{{Hooks: handler}}
	matchAll := []hookGroup{{Matcher: ".*", Hooks: handler}}
	switch config.Agents.Main {
	case "claude":
		servers := serversFor(config, "claude")
		names := make([]string, 0, len(servers))
		for _, server := range servers {
			names = append(names, server.Name)
		}
		settings := claudeSettingsFile{
			Permissions:           claudePermissions{Allow: config.Claude.PermissionsAllow},
			EnabledMcpjsonServers: names,
			Hooks: claudeHookEvents{
				SessionStart:       every,
				UserPromptSubmit:   every,
				PreToolUse:         matchAll,
				PostToolUse:        matchAll,
				PostToolUseFailure: matchAll,
				Stop:               every,
			},
		}
		settingsBytes, err := encodeJSON(settings)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output{Target: targets["claudeSettings"], Content: settingsBytes})
		if target, found := targets["mcpJson"]; found {
			mcpBytes, err := encodeJSON(map[string]any{"mcpServers": mcpServerMap(servers)})
			if err != nil {
				return nil, err
			}
			outputs = append(outputs, output{Target: target, Content: mcpBytes})
		}
	case "codex":
		hooks := codexHookFile{Hooks: codexHookEvents{
			SessionStart:     every,
			UserPromptSubmit: every,
			PreToolUse:       matchAll,
			PostToolUse:      matchAll,
			Stop:             every,
		}}
		hooksBytes, err := encodeJSON(hooks)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output{Target: targets["codexHooks"], Content: hooksBytes})
		if target, found := targets["codexConfig"]; found {
			outputs = append(outputs, output{Target: target, Content: []byte(codexServerToml(serversFor(config, "codex")))})
		}
	}
	return outputs, nil
}

func mcpServerMap(servers []harness.MCPServer) map[string]any {
	entries := map[string]any{}
	for _, server := range servers {
		if server.URL != "" {
			entries[server.Name] = map[string]any{"type": "http", "url": server.URL}
			continue
		}
		entries[server.Name] = map[string]any{"command": server.Command, "args": server.Args}
	}
	return entries
}

func codexServerToml(servers []harness.MCPServer) string {
	var lines []string
	for _, server := range servers {
		if len(lines) > 0 {
			lines = append(lines, "")
		}
		lines = append(lines, "[mcp_servers."+server.Name+"]")
		if server.URL != "" {
			lines = append(lines, "url = "+strconv.Quote(server.URL))
			continue
		}
		quoted := make([]string, len(server.Args))
		for index, arg := range server.Args {
			quoted[index] = strconv.Quote(arg)
		}
		lines = append(lines, "command = "+strconv.Quote(server.Command), "args = ["+strings.Join(quoted, ", ")+"]")
	}
	return strings.Join(lines, "\n") + "\n"
}
