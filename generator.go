package acgen

import (
	"fmt"
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

var generators = make(map[string]Generator)

func RegisterGenerator(name string, g Generator) {
	generators[name] = g
}

func LookGenerator(name string) (g Generator, err error) {
	if _, ok := generators[name]; !ok {
		return nil, fmt.Errorf("%s: is not supported", name)
	}
	return generators[name], nil
}
