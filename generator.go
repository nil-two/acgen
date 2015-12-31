package main

type Config struct {
}

type Generator interface {
	Generate(c *Config) (code string, err error)
}
