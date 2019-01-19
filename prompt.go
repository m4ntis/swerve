package swerve

// Prompt represents a prompt with which the user interacts.
//
// The Prompt is responsible both for reading user input as well as outputing
// back to the user. Each Shell holds a reference to a single Prompt, used to
// read commands and arguments from the user. Additionally, the same Prompt is
// passed to each command for user interaction capabilities.
type Prompt interface {
	Printf(format string, a ...interface{})
	Println(a ...interface{})

	Readline() string
	ReadPassword() string
}
