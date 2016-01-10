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

func RegisterGenerator(generatorName string, g Generator) {
	generators[generatorName] = g
}

func LookGenerator(generatorName string) (g Generator, err error) {
	if _, ok := generators[generatorName]; !ok {
		return nil, fmt.Errorf("%s: is not supported", generatorName)
	}
	return generators[generatorName], nil
}
