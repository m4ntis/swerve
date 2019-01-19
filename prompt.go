package swerve

type Prompt interface {
	Printf(format string, a ...interface{})
	Println(a ...interface{})

	Readline() string
	ReadPassword() string
}
