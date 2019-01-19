package swerve

import (
	"fmt"
	"strings"
)

// Command represents a command of the interactive debugger.
type Command struct {
	Name    string
	Aliases []string

	Run func(args []string)

	Desc  string
	Usage string
	Help  string
}

func (c *Command) Title() string {
	if len(c.Aliases) == 0 {
		return c.Name
	}

	return fmt.Sprintf("%s (alias: %s)", c.Name, strings.Join(c.Aliases, " | "))
}
