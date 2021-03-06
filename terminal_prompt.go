package swerve

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/peterh/liner"
)

// TerminalPrompt is a simple terminal Prompt implementation.
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

	if line != "" {
		t.l.AppendHistory(line)
	}

	return line
}
func (t *TerminalPrompt) ReadPassword() string {
	pass, err := t.l.PasswordPrompt(t.prompt)
	panicOnNonEOFErr(err)

	return pass
}

func (t *TerminalPrompt) Clear() {
	// TODO: support windows :(
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func panicOnNonEOFErr(err error) {
	if err == io.EOF {
		return
	}

	if err != nil {
		panic(err)
	}
}
