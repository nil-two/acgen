package acgen

import (
	"io"
	"strings"
	"text/template"
)

func init() {
	RegisterGenerator("fish", generateFishCompletion)
}

func toFishString(s string) string {
	return `'` + strings.Replace(s, `'`, `'"'"'`, -1) + `'`
}

type fish struct {
	Statements []string
}

func newFish(c *Command) (f *fish, err error) {
	f = new(fish)
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

var fishTemplate = template.Must(template.New("fish").Parse(`
{{range .Statements}}{{.}}
{{end}}`[1:]))

func generateFishCompletion(w io.Writer, c *Command) error {
	f, err := newFish(c)
	if err != nil {
		return err
	}
	return fishTemplate.Execute(w, f)
}
