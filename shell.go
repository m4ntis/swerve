package swerve

import (
	"fmt"
	"strings"
)

type Shell struct {
	cmds   []Command
	hashed map[string]Command

	p Prompt
}

func New() *Shell {
	return &Shell{}
}

// Run runs the Shell indefinitely, reading a line from the prompt and running
// the appropriate command with it's arguments.
func (s *Shell) Run() {
	for {
		line := s.Readline()
		args := strings.Fields(line)

		cmd, ok := s.hashed[args[0]]
		if !ok {
			s.p.Printf("%s isn't a valid command, run 'help' for a list\n", line)
			continue
		}

		cmd.Run(p, args[1:])
	}
}

// Add adds commands to the Shell's command list.
//
// Add will panic if a command with a name or alias identical to an existing
// command is added.
func (s *Shell) Add(cmds ...Command) {
	for _, cmd := range cmds {
		s.sortedAdd(s.cmds, cmd)
		s.hash(cmd)
	}
}

// SetPrompt sets the Shell's Prompt.
func (s *Shell) SetPrompt(p Prompt) {
	s.p = p
}

// sortedAdd adds a single Command to the Shell's command list in alphabetical
// name order.
func (s *Shell) sortedAdd(cmd Command) {
	for i, c := range s.cmds {
		if c.Name > cmd.Name {
			s.cmds = append(s.cmds[:i], append([]Command{cmd}, s.cmds[i:]...))
			return
		}
	}
}

// hash adds references to a command by it's name and aliases.
func (s *Shell) hash(cmd Command) {
	names := append(cmd.Aliases, name)

	for _, name := range names {
		_, ok := s.hashed[name]
		if ok {
			panic("You cannot have two commands with identical names")
		}

		s.hashed[name] = cmd
	}
}

// help generates an alphabetically sorted, multi-line help string for the
// Shell's command list, based on their name, aliases and description.
func (s *Shell) help() string {
	help := "The following commands are available:\n"

	longest := s.longestTitleLength()
	for _, cmd := range s.cmds {
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
func (s *Shell) longestTitleLength() int {
	longest := 0
	for _, cmd := range s.cmds {
		titleLen := len(cmd.Title())
		if titleLen > longest {
			longest = titleLen
		}
	}
	return longest
}
