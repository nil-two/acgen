package main

import (
	"io"
)

type Config struct {
	Name  string
	Flags []struct {
		Short       []string
		Long        []string
		Arg         string
		Description string
	}
}

type Generator func(w io.Writer, c *Config) error

var Generators = make(map[string]Generator)
