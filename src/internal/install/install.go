package install

import (
	"io"
	"os/exec"

	"simpleharness/src/internal/catalog"
	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/system"
)

type Item struct {
	Tool      catalog.Tool
	Installed bool
	Command   []string
	Manual    string
}

var managerLabels = map[string]string{
	catalog.ManagerWinget: "winget",
	catalog.ManagerBrew:   "Homebrew",
	catalog.ManagerApt:    "apt",
}

func ManagerLabel(manager string) string {
	return managerLabels[manager]
}

func Supported(manager string) bool {
	_, found := managerLabels[manager]
	return found
}

func recommendedIDs(info system.Info, config harness.Config) []string {
	ids := []string{"git", "node"}
	if info.OS == "windows" {
		ids = append(ids, "pwsh")
		if info.IsWindows10() {
			ids = append(ids, "windows-terminal")
		}
	}
	ids = append(ids, "jq", "ripgrep", "sqlite3", "duckdb")
	if config.Project.Kind == "program" {
		ids = append(ids, "llvm", "gcc")
	}
	switch config.Project.Language {
	case "python":
		ids = append(ids, "python", "uv")
	case "cpp":
		if info.OS == "windows" {
			ids = append(ids, "msvc")
		}
		ids = append(ids, "llvm", "xmake")
	case "rust":
		if info.OS == "windows" {
			ids = append(ids, "msvc")
		}
		ids = append(ids, "rustup")
	case "go":
		ids = append(ids, "go")
	case "php":
		ids = append(ids, "php", "composer")
	}
	runtimes := map[string]bool{config.Agents.Main: true}
	needsOllama := false
	for _, subagent := range config.Subagents {
		runtimes[subagent.Runtime] = true
		if subagent.Provider == "ollama-local" {
			needsOllama = true
		}
	}
	if runtimes["claude"] {
		ids = append(ids, "claude-code")
	}
	if runtimes["codex"] {
		ids = append(ids, "codex")
	}
	if needsOllama {
		ids = append(ids, "ollama")
	}
	seen := map[string]bool{}
	unique := make([]string, 0, len(ids))
	for _, id := range ids {
		if !seen[id] {
			seen[id] = true
			unique = append(unique, id)
		}
	}
	return unique
}

func Plan(info system.Info, config harness.Config) []Item {
	ids := recommendedIDs(info, config)
	items := make([]Item, 0, len(ids))
	for _, id := range ids {
		tool, found := catalog.Tools[id]
		if !found {
			panic("install: unknown tool " + id)
		}
		item := Item{Tool: tool, Installed: present(tool, info)}
		if arguments, found := tool.Packages[info.PackageManager]; found && Supported(info.PackageManager) {
			item.Command = installCommand(info.PackageManager, arguments)
		} else if manual, found := tool.Manual[info.OS]; found {
			item.Manual = manual
		} else {
			item.Manual = tool.Homepage
		}
		items = append(items, item)
	}
	return items
}

func present(tool catalog.Tool, info system.Info) bool {
	if tool.Detector == "msvc" {
		return system.HasMSVC()
	}
	for _, command := range tool.PresenceCommands(info.OS) {
		if system.Has(command) {
			return true
		}
	}
	return false
}

func installCommand(manager string, arguments []string) []string {
	switch manager {
	case catalog.ManagerWinget:
		return append(append([]string{"winget", "install"}, arguments...), "--accept-package-agreements", "--accept-source-agreements")
	case catalog.ManagerBrew:
		return append([]string{"brew", "install"}, arguments...)
	case catalog.ManagerApt:
		return append([]string{"sudo", "apt-get", "install", "-y"}, arguments...)
	}
	panic("install: unsupported package manager " + manager)
}

func Prepare(manager string, input io.Reader, output, errors io.Writer) error {
	if manager != catalog.ManagerApt {
		return nil
	}
	return run([]string{"sudo", "apt-get", "update"}, input, output, errors)
}

func Run(item Item, input io.Reader, output, errors io.Writer) error {
	return run(item.Command, input, output, errors)
}

func run(command []string, input io.Reader, output, errors io.Writer) error {
	process := exec.Command(command[0], command[1:]...)
	process.Stdin = input
	process.Stdout = output
	process.Stderr = errors
	return process.Run()
}
