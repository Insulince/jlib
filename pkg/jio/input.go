package jio

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh/terminal"
)

func Input() (string, error) {
	var input string
	if _, err := fmt.Scanln(&input); err != nil && !strings.Contains(err.Error(), "unexpected newline") {
		return "", errors.Wrap(err, "scanning line")
	}
	return input, nil
}

func MustInput() string {
	input, err := Input()
	if err != nil {
		panic(errors.Wrap(err, "must input"))
	}
	return input
}

func InputWithPrompt(prompt string) (string, error) {
	Output(prompt)
	input, err := Input()
	if err != nil {
		return "", errors.Wrap(err, "input")
	}
	return input, nil
}

func MustInputWithPrompt(prompt string) string {
	input, err := InputWithPrompt(prompt)
	if err != nil {
		panic(errors.Wrap(err, "must input with prompt"))
	}
	return input
}

func InputWithPromptf(prompt string, args ...interface{}) (string, error) {
	Outputf(prompt, args...)
	input, err := Input()
	if err != nil {
		return "", errors.Wrap(err, "input")
	}
	return input, nil
}

func MustInputWithPromptf(prompt string, args ...interface{}) string {
	input, err := InputWithPromptf(prompt, args...)
	if err != nil {
		panic(errors.Wrap(err, "must input with prompt format"))
	}
	return input
}

func InputWithPromptln(prompt string) (string, error) {
	Outputln(prompt)
	input, err := Input()
	if err != nil {
		return "", errors.Wrap(err, "input")
	}
	return input, nil
}

func MustInputWithPromptln(prompt string) string {
	input, err := InputWithPromptln(prompt)
	if err != nil {
		panic(errors.Wrap(err, "must input with prompt line"))
	}
	return input
}

// PRIVATE INPUT

func PrivateInput() (string, error) {
	SilentOutput("⚷")
	raw, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", errors.Wrap(err, "privately scanning line")
	}
	return string(raw), nil
}

func MustPrivateInput() string {
	input, err := PrivateInput()
	if err != nil {
		panic(errors.Wrap(err, "must private input"))
	}
	return input
}

func PrivateInputWithPrompt(prompt string) (string, error) {
	Output(prompt)
	input, err := PrivateInput()
	if err != nil {
		return "", errors.Wrap(err, "input")
	}
	return input, nil
}

func MustPrivateInputWithPrompt(prompt string) string {
	input, err := PrivateInputWithPrompt(prompt)
	if err != nil {
		panic(errors.Wrap(err, "must private input with prompt"))
	}
	return input
}

func PrivateInputWithPromptf(prompt string, args ...interface{}) (string, error) {
	Outputf(prompt, args...)
	input, err := PrivateInput()
	if err != nil {
		return "", errors.Wrap(err, "private input")
	}
	return input, nil
}

func MustPrivateInputWithPromptf(prompt string, args ...interface{}) string {
	input, err := PrivateInputWithPromptf(prompt, args...)
	if err != nil {
		panic(errors.Wrap(err, "must private input with prompt format"))
	}
	return input
}

func PrivateInputWithPromptln(prompt string) (string, error) {
	Outputln(prompt)
	input, err := PrivateInput()
	if err != nil {
		return "", errors.Wrap(err, "private input")
	}
	return input, nil
}

func MustPrivateInputWithPromptln(prompt string) string {
	input, err := PrivateInputWithPromptln(prompt)
	if err != nil {
		panic(errors.Wrap(err, "must private input with prompt line"))
	}
	return input
}
