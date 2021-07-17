package io

import (
	"fmt"
)

func Output(msg string) {
	fmt.Print("> ", msg)
}

func Outputf(msg string, args ...interface{}) {
	Output(fmt.Sprintf(msg, args...))
}

func Outputln(msg string) {
	Output(fmt.Sprintln(msg))
}

// SILENT OUTPUT

func SilentOutput(msg string) {
	fmt.Print(msg)
}

func SilentOutputf(msg string, args ...interface{}) {
	SilentOutput(fmt.Sprintf(msg, args...))
}

func SilentOutputln(msg string) {
	SilentOutput(fmt.Sprintln(msg))
}
