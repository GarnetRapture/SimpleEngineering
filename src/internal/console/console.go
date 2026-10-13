package console

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"simpleharness/src/internal/locale"
)

var ErrInputEnded = errors.New("console: input ended before a required answer")

type Option struct {
	Label  string
	Detail string
}

type Console struct {
	reader   *bufio.Reader
	writer   io.Writer
	language locale.Language
}

func New(input io.Reader, output io.Writer) *Console {
	return &Console{reader: bufio.NewReader(input), writer: output, language: locale.English}
}

func (console *Console) SetLanguage(language locale.Language) {
	console.language = language
}

func (console *Console) Language() locale.Language {
	return console.language
}

func (console *Console) Say(text locale.Text, args ...any) {
	console.Line(text.Format(console.language, args...))
}

func (console *Console) Line(text string) {
	fmt.Fprintln(console.writer, text)
}

func (console *Console) Heading(text locale.Text) {
	fmt.Fprintf(console.writer, "\n== %s ==\n", text.In(console.language))
}

func (console *Console) readAnswer() (string, bool, error) {
	fmt.Fprint(console.writer, "> ")
	line, err := console.reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", false, err
	}
	answer := strings.TrimSpace(line)
	ended := errors.Is(err, io.EOF)
	if ended {
		fmt.Fprintln(console.writer)
	}
	return answer, ended, nil
}

func (console *Console) printOptions(options []Option, marks []string) {
	for index, option := range options {
		mark := ""
		if marks != nil {
			mark = marks[index] + " "
		}
		if option.Detail == "" {
			fmt.Fprintf(console.writer, "  %d) %s%s\n", index+1, mark, option.Label)
			continue
		}
		fmt.Fprintf(console.writer, "  %d) %s%s — %s\n", index+1, mark, option.Label, option.Detail)
	}
}

func (console *Console) Choose(question string, options []Option, defaultIndex int) (int, error) {
	console.Line(question)
	console.printOptions(options, nil)
	for {
		if defaultIndex >= 0 {
			console.Say(locale.ChooseHint, defaultIndex+1)
		} else {
			console.Say(locale.ChooseRequiredHint)
		}
		answer, ended, err := console.readAnswer()
		if err != nil {
			return 0, err
		}
		if answer == "" && defaultIndex >= 0 {
			return defaultIndex, nil
		}
		number, parseErr := strconv.Atoi(answer)
		if parseErr == nil && number >= 1 && number <= len(options) {
			return number - 1, nil
		}
		if ended {
			return 0, ErrInputEnded
		}
		console.Say(locale.InvalidChoice, answer, len(options))
	}
}

func (console *Console) ChooseMany(question string, options []Option, selected []bool) ([]bool, error) {
	current := append([]bool(nil), selected...)
	console.Line(question)
	for {
		marks := make([]string, len(options))
		for index := range options {
			marks[index] = "[ ]"
			if current[index] {
				marks[index] = "[x]"
			}
		}
		console.printOptions(options, marks)
		console.Say(locale.ManyHint)
		answer, ended, err := console.readAnswer()
		if err != nil {
			return nil, err
		}
		if answer == "" {
			return current, nil
		}
		invalid := ""
		for _, field := range strings.FieldsFunc(answer, func(r rune) bool { return r == ',' || r == ' ' || r == '\t' }) {
			number, parseErr := strconv.Atoi(field)
			if parseErr != nil || number < 1 || number > len(options) {
				invalid = field
				break
			}
			current[number-1] = !current[number-1]
		}
		if invalid != "" {
			console.Say(locale.InvalidChoice, invalid, len(options))
		}
		if ended {
			return current, nil
		}
	}
}

func (console *Console) Text(question, defaultValue string, validate func(string) bool, invalid locale.Text) (string, error) {
	console.Line(question)
	for {
		if defaultValue != "" {
			fmt.Fprintf(console.writer, "  Enter = %s\n", defaultValue)
		} else {
			console.Line("  " + locale.TextRequiredHint.In(console.language))
		}
		answer, ended, err := console.readAnswer()
		if err != nil {
			return "", err
		}
		if answer == "" {
			answer = defaultValue
		}
		if answer != "" && validate(answer) {
			return answer, nil
		}
		if ended {
			return "", ErrInputEnded
		}
		if answer != "" {
			console.Say(invalid)
		}
	}
}

func (console *Console) OptionalText(question string) (string, error) {
	console.Line(question)
	answer, _, err := console.readAnswer()
	return answer, err
}

func (console *Console) Confirm(question string, defaultYes bool) (bool, error) {
	hint := locale.ConfirmNoHint
	if defaultYes {
		hint = locale.ConfirmYesHint
	}
	for {
		console.Line(question + " " + hint.In(console.language))
		answer, ended, err := console.readAnswer()
		if err != nil {
			return false, err
		}
		switch strings.ToLower(answer) {
		case "":
			return defaultYes, nil
		case "y", "yes", "예", "네", "是":
			return true, nil
		case "n", "no", "아니오", "아니요", "否":
			return false, nil
		}
		if ended {
			return defaultYes, nil
		}
		console.Say(locale.InvalidConfirm)
	}
}
