package acgen

import (
	"io"
	"strings"
	"text/template"
)

func init() {
	RegisterGenerator("tcsh", generateTcshCompletion)
}

type tcsh struct {
	Name string
	Opt  string
}

func newTcsh(c *Command) (t *tcsh, err error) {
	opts := make([]string, 0)
	for _, flag := range c.Flags {
		for _, opt := range flag.Long {
			opts = append(opts, opt)
		}
	}
	return &tcsh{
		Name: c.Name,
		Opt:  strings.Join(opts, " "),
	}, nil
}

var tcshTemplate = template.Must(template.New("tcsh").Parse(`
complete {{.Name}} 'c/--/({{.Opt}})/'
`[1:]))

func generateTcshCompletion(w io.Writer, c *Command) error {
	t, err := newTcsh(c)
	if err != nil {
		return err
	}
	return tcshTemplate.Execute(w, t)
}
