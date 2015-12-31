package main

import (
	"io"
)

type Flag struct {
	Short       []string
	Long        []string
	Arg         string
	Description string
}

type Config struct {
	Name  string
	Flags []Flag
}

type Generator func(w io.Writer, c *Config) error

var Generators = make(map[string]Generator)
