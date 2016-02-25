package acgen

import (
	"fmt"
	"io"
	"sync"
)

// A Command represents a command which has flags.
type Command struct {
	Name  string
	Flags []*Flag
}

// A Flag represents the information of a flag.
type Flag struct {
	Short       []string // short options
	Long        []string // long options
	Arg         string   // argument's name
	Description string   // help message
}

// A Generator writes a completion for command to w.
type Generator func(w io.Writer, c *Command) error

var (
	generatorsMu sync.Mutex
	generators   = make(map[string]Generator)
)

// RegisterGenerator makes a completion generator available
// by the provided name.
func RegisterGenerator(name string, g Generator) {
	generatorsMu.Lock()
	defer generatorsMu.Unlock()
	if _, dup := generators[name]; dup {
		panic("RegisterGenerator called twice for generator " + name)
	}
	generators[name] = g
}

// LookGenerator returns a completion generator
// specified by its completion generator name.
func LookGenerator(name string) (g Generator, err error) {
	generatorsMu.Lock()
	defer generatorsMu.Unlock()
	if _, ok := generators[name]; !ok {
		return nil, fmt.Errorf("%s: is not supported", name)
	}
	return generators[name], nil
}
