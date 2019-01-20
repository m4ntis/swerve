package swerve

import (
	"fmt"
	"strings"
)

// Command represents a command of the interactive debugger.
type Command struct {
	// Name is the command's unique name, used to identify the command. Spaces
	// aren't supported.
	Name string
	// Aliases are optional additional names for the commands. Spaces aren't
	// supported.
	Aliases []string

	// Run is the main command's method, executed whenever the command is
	// executed with valid arguments.
	Run func(p Prompt, args []string)

	// ValidateArgs is an optional function to validate a Commands arguments
	// before being run.
	ValidateArgs func(p Prompt, args []string) (ok bool)

	// Desc is the command's shortest description. Desc is displayed when using
	// the builtin 'help' command.
	Desc string
	// Usage is the command's usage string. If provided, it will be displayed
	// when using 'help <cmd>'
	Usage string
	// Help is the command's long string. If provided, it will be displayed when
	// using 'help <cmd>'
	Help string
}

// Title returnes the command's title, a formatted string consisting of it's
// name and aliases (if any).
func (c *Command) Title() string {
	if len(c.Aliases) == 0 {
		return c.Name
	}

	return fmt.Sprintf("%s (alias: %s)", c.Name, strings.Join(c.Aliases, " | "))
}
