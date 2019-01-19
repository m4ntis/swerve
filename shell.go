package swerve

import (
	"fmt"
	"strings"
)

type Shell struct {
	Cmds []Command
}

// Add adds commands to the Shell's command list.
//
// Add will panic if a command with an identical name to an existing command is
// added.
func (s Shell) Add(cmds ...Command) {
	for _, cmd := range cmds {
		s.sortedAdd(s.Cmds, cmd)
	}
}

// sortedAdd adds a single Command to the Shell's command list in alphabetical
// name order.
func (s Shell) sortedAdd(cmd Command) {
	for i, c := range s.cmds {
		if c.Name == cmd.Name {
			panic("You cannot have two commands with identical names")
		}

		if c.Name > cmd.Name {
			s.Cmds = append(s.Cmds[:i], append([]Command{cmd}, s.Cmds[i:]...))
			return
		}
	}
}

// help generates an alphabetically sorted, multi-line help string for the
// Shell's command list, based on their name, aliases and description.
func (s Shell) help() string {
	help := "The following commands are available:\n"

	longest := s.longestTitleLength()
	for _, cmd := range s.Cmds {
		help += fmt.Sprintf("    %s %s %s\n",
			cmd.Title(),
			strings.Repeat("-", longest-len(title)+1),
			cmd.Desc)
	}

	help += "Type 'help' followed by a command's name or alias for full documentation"
	return help
}

// longestTitleLength returns the length of the longest title from the commands
// in the shell.
//
// This can more efficiently be stored in the shell's context and calculated
// once on Shell.Add, but that would add the the Shell's context and
// documentation.
func (s Shell) longestTitleLength() int {
	longest := 0
	for _, cmd := range s.cmds {
		titleLen := len(cmd.Title())
		if titleLen > longest {
			longest = titleLen
		}
	}
	return longest
}
