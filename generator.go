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

var Generators = make(map[string]Generator)

func RegisterGenerator(generatorName string, g Generator) {
	Generators[generatorName] = g
}

func LookGenerator(generatorName string) (g Generator, err error) {
	if _, ok := Generators[generatorName]; !ok {
		return nil, fmt.Errorf("%s: is not supported", generatorName)
	}
	return Generators[generatorName], nil
}
