package render

import (
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var placeholder = regexp.MustCompile(`\{\{([A-Za-z][A-Za-z0-9_.:/-]*)\}\}`)
var lineFragment = regexp.MustCompile(`^\s*\{\{generated\.([A-Za-z]+)\}\}\s*$`)

type Context struct {
	values    map[string]any
	Links     map[string]string
	Fragments map[string]string
	Errors    []string
}

func NewContext(config any) (*Context, error) {
	encoded, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	var values map[string]any
	if err := json.Unmarshal(encoded, &values); err != nil {
		return nil, err
	}
	return &Context{values: values, Links: map[string]string{}, Fragments: map[string]string{}}, nil
}

func (context *Context) Lookup(key string) (string, bool) {
	var value any = context.values
	for _, part := range strings.Split(key, ".") {
		object, isObject := value.(map[string]any)
		if !isObject {
			return "", false
		}
		next, found := object[part]
		if !found {
			return "", false
		}
		value = next
	}
	switch typed := value.(type) {
	case string:
		return typed, true
	case float64:
		return strconv.FormatFloat(typed, 'f', -1, 64), true
	}
	return "", false
}

func (context *Context) fail(origin, message string) {
	context.Errors = append(context.Errors, origin+": "+message)
}

func (context *Context) Target(target, origin string) string {
	return placeholder.ReplaceAllStringFunc(target, func(whole string) string {
		key := whole[2 : len(whole)-2]
		value, found := context.Lookup(key)
		if !found {
			context.fail(origin, "unknown target placeholder "+whole)
			return whole
		}
		return value
	})
}

func (context *Context) Text(text, target, origin string) string {
	lines := strings.Split(text, "\n")
	kept := make([]string, 0, len(lines))
	for _, line := range lines {
		match := lineFragment.FindStringSubmatch(line)
		if match == nil {
			kept = append(kept, line)
			continue
		}
		fragment, found := context.Fragments[match[1]]
		if !found {
			context.fail(origin, "unknown placeholder {{generated."+match[1]+"}}")
			kept = append(kept, line)
			continue
		}
		if fragment != "" {
			kept = append(kept, fragment)
		}
	}
	withFragments := placeholder.ReplaceAllStringFunc(strings.Join(kept, "\n"), func(whole string) string {
		key := whole[2 : len(whole)-2]
		if !strings.HasPrefix(key, "generated.") {
			return whole
		}
		fragment, found := context.Fragments[strings.TrimPrefix(key, "generated.")]
		if !found {
			context.fail(origin, "unknown placeholder "+whole)
			return whole
		}
		return fragment
	})
	return placeholder.ReplaceAllStringFunc(withFragments, func(whole string) string {
		key := whole[2 : len(whole)-2]
		switch {
		case strings.HasPrefix(key, "link."):
			linkTarget, found := context.Links[strings.TrimPrefix(key, "link.")]
			if !found {
				context.fail(origin, "unknown link "+whole)
				return whole
			}
			return RelativeLink(target, linkTarget)
		case strings.HasPrefix(key, "path:"):
			return RelativeLink(target, strings.TrimPrefix(key, "path:"))
		}
		value, found := context.Lookup(key)
		if !found {
			context.fail(origin, "unknown placeholder "+whole)
			return whole
		}
		return value
	})
}

func RelativeLink(fromTarget, toTarget string) string {
	fromParts := splitPath(path.Dir(fromTarget))
	toParts := splitPath(toTarget)
	common := 0
	for common < len(fromParts) && common < len(toParts)-1 && fromParts[common] == toParts[common] {
		common++
	}
	parts := make([]string, 0, len(fromParts)-common+len(toParts)-common)
	for range fromParts[common:] {
		parts = append(parts, "..")
	}
	parts = append(parts, toParts[common:]...)
	return strings.Join(parts, "/")
}

func splitPath(value string) []string {
	if value == "." || value == "" {
		return nil
	}
	return strings.Split(value, "/")
}

func Code(text string) string {
	return "`" + text + "`"
}

func Cell(text string) string {
	return strings.ReplaceAll(text, "|", `\|`)
}

func Row(cells ...string) string {
	escaped := make([]string, len(cells))
	for index, cell := range cells {
		escaped[index] = Cell(cell)
	}
	return "| " + strings.Join(escaped, " | ") + " |"
}

func Join(values []string, separator string, format func(string) string) string {
	formatted := make([]string, len(values))
	for index, value := range values {
		formatted[index] = format(value)
	}
	return strings.Join(formatted, separator)
}

func Errorf(errors []string) error {
	if len(errors) == 0 {
		return nil
	}
	return fmt.Errorf("%s", strings.Join(errors, "\n"))
}
