package acgen

import (
	"io"
)

type Flag struct {
	Short       []string
	Long        []string
	Arg         string
	Description string
}

type Command struct {
	Name  string
	Flags []*Flag
}

type Generator func(w io.Writer, c *Command) error

var Generators = make(map[string]Generator)
