package main

import (
	"io"
	"strings"
	"text/template"
)

func init() {
	Generators["fish"] = GenerateFishCompletion
}

func toFishString(s string) string {
	return `'` + strings.Replace(s, `'`, `'"'"'`, -1) + `'`
}

type Fish struct {
	Statements []string
}

func NewFish(c *Command) (f *Fish, err error) {
	f = new(Fish)
	for _, flag := range c.Flags {
		options := []string{"complete", "-c", toFishString(c.Name)}
		for _, short := range flag.Short {
			options = append(options, "-s", toFishString(short))
		}
		for _, long := range flag.Long {
			options = append(options, "-l", toFishString(long))
		}
		options = append(options, "-d", toFishString(flag.Description))
		statement := strings.Join(options, " ")
		f.Statements = append(f.Statements, statement)
	}
	return f, nil
}

var FishCompletionTemplateText = `
{{range .Statements}}{{.}}
{{end}}`[1:]

func GenerateFishCompletion(w io.Writer, c *Command) error {
	tmpl, err := template.New("fish").Parse(FishCompletionTemplateText)
	if err != nil {
		return err
	}
	f, err := NewFish(c)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, f)
}
