package swerve

import (
	"fmt"
	"io"

	"github.com/peterh/liner"
)

type TerminalPrompt struct {
	prompt string

	l *liner.State
}

func NewTerminalPrompt(prompt string, l *liner.State) *TerminalPrompt {
	return &TerminalPrompt{prompt, l}
}

func (t *TerminalPrompt) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func (t *TerminalPrompt) Println(a ...interface{}) {
	fmt.Println(a...)
}

func (t *TerminalPrompt) Readline() string {
	line, err := t.l.Prompt(t.prompt)
	panicOnNonEOFErr(err)

	return line
}
func (t *TerminalPrompt) ReadPassword() string {
	pass, err := t.l.PasswordPrompt(t.prompt)
	panicOnNonEOFErr(err)

	return pass
}

func panicOnErr(err error) {
	if err == io.EOF {
		return
	}

	if err != nil {
		panic(err)
	}
}
