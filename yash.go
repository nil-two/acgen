package main

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

func init() {
	Generators["yash"] = GenerateYashCompletion
}

func toYashOpt(f *Flag) string {
	var opts []string
	for _, short := range f.Short {
		opt := short
		if f.Arg != "" {
			opt += ":"
		}
		opts = append(opts, opt)
	}
	for _, long := range f.Long {
		opt := "--" + long
		if f.Arg != "" {
			opt += ":"
		}
		opts = append(opts, opt)
	}

	description := f.Description

	return fmt.Sprintf("'%s; %s'",
		strings.Join(opts, " "), description)
}

type Yash struct {
	Name string
	Opts []string
}

func NewYash(c *Config) (y *Yash, err error) {
	y = new(Yash)
	y.Name = c.Name
	for _, flag := range c.Flags {
		y.Opts = append(y.Opts, toYashOpt(&flag))
	}
	return y, nil
}

var YashCompletionTemplateText = `
function completion/{{.Name}} {
	typeset OPTIONS ARGOPT PREFIX
	OPTIONS=({{range .Opts}}
		{{.}}{{end}}
	)
	command -f completion//parseoptions -es
	case $ARGOPT in
	(-)
		command -f completion//completeoptions
		;;
	(*)
		complete -f
		;;
	esac
}
# vim: set ft=sh ts=8 sts=8 sw=8 noet:
`[1:]

func GenerateYashCompletion(w io.Writer, c *Config) error {
	tmpl, err := template.New("yash").Parse(YashCompletionTemplateText)
	if err != nil {
		return err
	}
	y, err := NewYash(c)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, y)
}
