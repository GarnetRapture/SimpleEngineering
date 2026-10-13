package generate

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"simpleharness/src/internal/harness"
	"simpleharness/src/internal/locale"
	"simpleharness/src/internal/render"
)

type manifestFile struct {
	ID         string `json:"id"`
	Source     string `json:"source"`
	Target     string `json:"target"`
	Render     bool   `json:"render"`
	Executable bool   `json:"executable"`
	Harness    string `json:"harness,omitempty"`
	Requires   string `json:"requires,omitempty"`
}

type manifest struct {
	Files []manifestFile `json:"files"`
}

type output struct {
	Target     string
	Content    []byte
	Executable bool
}

func Write(template fs.FS, config harness.Config, destination string) ([]string, error) {
	outputs, err := Plan(template, config)
	if err != nil {
		return nil, err
	}
	written := make([]string, 0, len(outputs))
	for _, item := range outputs {
		absolute := filepath.Join(destination, filepath.FromSlash(item.Target))
		if err := os.MkdirAll(filepath.Dir(absolute), 0o755); err != nil {
			return written, err
		}
		mode := os.FileMode(0o644)
		if item.Executable {
			mode = 0o755
		}
		file, err := os.OpenFile(absolute, os.O_WRONLY|os.O_CREATE|os.O_EXCL, mode)
		if err != nil {
			return written, err
		}
		_, writeErr := file.Write(item.Content)
		closeErr := file.Close()
		if writeErr != nil {
			return written, writeErr
		}
		if closeErr != nil {
			return written, closeErr
		}
		written = append(written, item.Target)
	}
	return written, nil
}

func Plan(template fs.FS, config harness.Config) ([]output, error) {
	language := locale.Language(config.Output.Language)
	manifestBytes, err := fs.ReadFile(template, "manifest.json")
	if err != nil {
		return nil, err
	}
	var parsed manifest
	if err := json.Unmarshal(manifestBytes, &parsed); err != nil {
		return nil, fmt.Errorf("manifest.json: %w", err)
	}
	context, err := render.NewContext(config)
	if err != nil {
		return nil, err
	}
	selected := make([]manifestFile, 0, len(parsed.Files))
	for _, file := range parsed.Files {
		if !included(file, config) {
			continue
		}
		file.Source = context.Target(file.Source, "manifest "+file.ID)
		file.Target = context.Target(file.Target, "manifest "+file.ID)
		context.Links[file.ID] = file.Target
		selected = append(selected, file)
	}
	for id, target := range generatedTargets(config) {
		context.Links[id] = target
	}
	fragments, err := buildFragments(template, config, language)
	if err != nil {
		return nil, err
	}
	context.Fragments = fragments
	outputs := make([]output, 0, len(selected)+5)
	for _, file := range selected {
		content, err := fs.ReadFile(template, file.Source)
		if err != nil {
			return nil, fmt.Errorf("manifest %s: %w", file.ID, err)
		}
		if file.Render {
			content = []byte(context.Text(string(content), file.Target, file.Source))
		}
		outputs = append(outputs, output{Target: file.Target, Content: content, Executable: file.Executable})
	}
	if err := render.Errorf(context.Errors); err != nil {
		return nil, err
	}
	registrations, err := registrationFiles(config)
	if err != nil {
		return nil, err
	}
	outputs = append(outputs, registrations...)
	seen := map[string]bool{}
	for _, item := range outputs {
		clean := path.Clean(item.Target)
		if path.IsAbs(clean) || strings.HasPrefix(clean, "../") || clean == ".." {
			return nil, fmt.Errorf("output outside the project root: %s", item.Target)
		}
		if seen[clean] {
			return nil, fmt.Errorf("two outputs share one target: %s", item.Target)
		}
		seen[clean] = true
	}
	return outputs, nil
}

func included(file manifestFile, config harness.Config) bool {
	if file.Harness != "" && !usesHarness(config, file.Harness) {
		return false
	}
	switch file.Requires {
	case "":
		return true
	case "subagents":
		return len(config.Subagents) > 0
	case "naming", "inline":
		return hasDomainSkill(config, file.Requires)
	}
	panic("generate: unknown manifest requirement " + file.Requires)
}

func usesHarness(config harness.Config, name string) bool {
	if config.Agents.Main == name {
		return true
	}
	for _, subagent := range config.Subagents {
		if subagent.Runtime == name {
			return true
		}
	}
	return false
}

func hasDomainSkill(config harness.Config, name string) bool {
	for _, skill := range config.Skills.Domain {
		if skill.Name == name {
			return true
		}
	}
	return false
}
