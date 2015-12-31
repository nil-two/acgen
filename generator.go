package main

type Config struct {
	Name  string
	Flags []struct {
		Short       []string
		Long        []string
		Arg         string
		Description string
	}
}

type Generator interface {
	Generate(c *Config) (code string, err error)
}

var Generators = make(map[string]Generator)
