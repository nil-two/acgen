package main

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

func init() {
	Generators["zsh"] = GenerateZshCompletion
}

func escapeZshString(s string) string {
	return strings.Replace(s, `'`, `'"'"'`, -1)
}

func toZshPropaty(f *Flag) string {
	opts := make([]string, 0, len(f.Short)+len(f.Long))
	for _, short := range f.Short {
		opts = append(opts, "-"+escapeZshString(short))
	}
	for _, long := range f.Long {
		opts = append(opts, "--"+escapeZshString(long))
	}

	exclusive := strings.Join(opts, " ")
	candidate := "'" + opts[0] + "'"
	if len(opts) > 1 {
		candidate = "{'" + strings.Join(opts, "','") + "'}"
	}
	description := escapeZshString(f.Description)
	argument := ""
	if f.Arg != "" {
		argument = ":" + escapeZshString(f.Arg)
	}
	return fmt.Sprintf("'(%s)'%s'[%s]%s'",
		exclusive, candidate, description, argument)
}

type Zsh struct {
	Name      string
	Propaties []string
}

func NewZsh(c *Command) (z *Zsh, err error) {
	z = new(Zsh)
	z.Name = c.Name
	z.Propaties = make([]string, 0, len(c.Flags))
	for _, flag := range c.Flags {
		z.Propaties = append(z.Propaties, toZshPropaty(flag))
	}
	return z, nil
}

var ZshCompletionTemplateText = `
#compdef {{.Name}}
_arguments \{{range .Propaties}}
    {{.}} \{{end}}
    '*:input files:_files'
`[1:]

func GenerateZshCompletion(w io.Writer, c *Command) error {
	tmpl, err := template.New("zsh").Parse(ZshCompletionTemplateText)
	if err != nil {
		return err
	}
	z, err := NewZsh(c)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, z)
}
