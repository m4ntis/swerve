package swerve

func builtins(s *Shell) []Command {
	return []Command{
		Command{
			Name: "clear",

			Run: func(p Prompt, args []string) {
				p.Clear()
			},

			Desc: "Clear the screen",
		},
		Command{
			Name:    "exit",
			Aliases: []string{"quit", "q"},

			Run: func(p Prompt, args []string) {
				close(s.exitc)
			},

			Desc: "Exit the CLI",
		},
		Command{
			Name:    "help",
			Aliases: []string{"h"},

			Run: func(p Prompt, args []string) {
				if len(args) == 0 {
					p.Println(s.help())
					return
				}

				cmd, ok := s.hashed[args[0]]
				if !ok {
					p.Printf("%s isn't a valid command, run 'help' for a list\n",
						args[0])
					return
				}

				if cmd.Desc != "" {
					p.Println(cmd.Desc)
				}
				if cmd.Usage != "" {
					if cmd.Desc != "" {
						p.Println()
					}
					p.Println("    " + cmd.Usage)
				}
				if cmd.Help != "" {
					if cmd.Desc != "" || cmd.Usage != "" {
						p.Println()
					}
					p.Println(cmd.Help)
				}
			},

			Desc:  "Get a list of commands or help on each",
			Usage: "help [command]",
			Help:  "Run 'help' to get a list of commands, or help about a specific command by appending it's name",
		},
	}
}
