package locale

import "fmt"

type Language string

const (
	Korean  Language = "ko"
	English Language = "en"
	Chinese Language = "zh"
)

var Languages = []Language{Korean, English, Chinese}

var LanguageNames = map[Language]string{
	Korean:  "한국어",
	English: "English",
	Chinese: "简体中文",
}

type Text struct {
	Ko string
	En string
	Zh string
}

func (text Text) In(language Language) string {
	switch language {
	case Korean:
		return text.Ko
	case English:
		return text.En
	case Chinese:
		return text.Zh
	}
	panic("locale: unsupported output language " + string(language))
}

func (text Text) Format(language Language, args ...any) string {
	return fmt.Sprintf(text.In(language), args...)
}
