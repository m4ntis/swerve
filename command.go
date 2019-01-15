package swerve

// Command represents a command of the interactive debugger.
type Command struct {
	Name string

	Run func(args []string)

	Desc  string
	Usage string
	Help  string
}
