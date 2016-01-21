package acgen

import (
	"fmt"
	"io"
	"sync"
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

var (
	generatorsMu sync.Mutex
	generators   = make(map[string]Generator)
)

func RegisterGenerator(name string, g Generator) {
	generatorsMu.Lock()
	defer generatorsMu.Unlock()
	generators[name] = g
}

func LookGenerator(name string) (g Generator, err error) {
	generatorsMu.Lock()
	defer generatorsMu.Unlock()
	if _, ok := generators[name]; !ok {
		return nil, fmt.Errorf("%s: is not supported", name)
	}
	return generators[name], nil
}
