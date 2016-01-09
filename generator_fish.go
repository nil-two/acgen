package acgen

import (
	"io"
	"strings"
	"text/template"
)

func init() {
	RegisterGenerator("fish", GenerateFishCompletion)
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
		opts := []string{"complete", "-c", toFishString(c.Name)}
		for _, short := range flag.Short {
			opts = append(opts, "-s", toFishString(short))
		}
		for _, long := range flag.Long {
			opts = append(opts, "-l", toFishString(long))
		}
		opts = append(opts, "-d", toFishString(flag.Description))
		statement := strings.Join(opts, " ")
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
