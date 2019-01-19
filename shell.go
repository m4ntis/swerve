package swerve

import (
	"fmt"
	"strings"

	"github.com/peterh/liner"
)

// Shell represents an interactive shell, holding a list of given commands.
type Shell struct {
	cmds   []Command
	hashed map[string]Command

	p        Prompt
	lastLine string

	exitc chan struct{}
}

// New returns a Shell with a terminal Prompt, with basic command name
// completion.
func New(prompt string) *Shell {
	s := &Shell{
		cmds:   []Command{},
		hashed: map[string]Command{},
	}

	s.p = s.defaultPrompt(prompt)
	s.Add(builtins(s)...)
	return s
}

// NewWithPrompt returns a Shell with a specified Prompt.
func NewWithPrompt(p Prompt) *Shell {
	s := &Shell{
		cmds:   []Command{},
		hashed: map[string]Command{},

		p: p,
	}

	s.Add(builtins(s)...)
	return s
}

// Run runs the Shell indefinitely, reading a line from the prompt and running
// the appropriate command with it's arguments.
func (s *Shell) Run() {
	for {
		select {
		case <-s.exitc:
			return
		default:
			s.readCommand()
		}
	}
}

// Add adds commands to the Shell's command list.
//
// Add will panic if a command with a name or alias identical to an existing
// command is added.
func (s *Shell) Add(cmds ...Command) {
	for _, cmd := range cmds {
		s.sortedAdd(cmd)
		s.hash(cmd)
	}
}

// sortedAdd adds a single Command to the Shell's command list in alphabetical
// name order.
func (s *Shell) sortedAdd(cmd Command) {
	for i, c := range s.cmds {
		if c.Name > cmd.Name {
			s.cmds = append(s.cmds[:i], append([]Command{cmd}, s.cmds[i:]...)...)
			return
		}
	}
}

// hash adds references to a command by it's name and aliases.
func (s *Shell) hash(cmd Command) {
	names := append(cmd.Aliases, cmd.Name)

	for _, name := range names {
		_, ok := s.hashed[name]
		if ok {
			panic("You cannot have two commands with identical names")
		}

		s.hashed[name] = cmd
	}
}

func (s *Shell) readCommand() {
	line := s.p.Readline()

	// Handle empty lines as a repeat of last command
	if line == "" {
		if s.lastLine == "" {
			return
		}

		line = s.lastLine
	}
	s.lastLine = line

	// Parse line
	args := strings.Fields(line)
	cmd, ok := s.hashed[args[0]]
	if !ok {
		s.p.Printf("%s isn't a valid command, run 'help' for a list\n", args[0])
		return
	}

	// Run command
	if cmd.ValidateArgs != nil && !cmd.ValidateArgs(args[1:]) {
		return
	}
	cmd.Run(s.p, args[1:])
}

// defaultPrompt creates a terminal prompt with basic command name completion.
func (s *Shell) defaultPrompt(prompt string) Prompt {
	l := liner.NewLiner()

	l.SetCompleter(func(line string) (opts []string) {
		for _, cmd := range s.cmds {
			if strings.HasPrefix(cmd.Name, strings.ToLower(line)) {
				opts = append(opts, cmd.Name)
			}
		}
		return opts
	})

	return NewTerminalPrompt(prompt, l)
}

// help generates an alphabetically sorted, multi-line help string for the
// Shell's command list, based on their name, aliases and description.
func (s *Shell) help() string {
	help := "The following commands are available:\n"

	longest := s.longestTitleLength()
	for _, cmd := range s.cmds {
		title := cmd.Title()

		help += fmt.Sprintf("    %s %s %s\n",
			title,
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
