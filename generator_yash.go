package acgen

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

func init() {
	RegisterGenerator("yash", generateYashCompletion)
}

func escapeYashString(s string) string {
	return strings.Replace(s, `'`, `'"'"'`, -1)
}

func toYashOpt(f *Flag) string {
	var opts []string
	for _, short := range f.Short {
		opt := escapeYashString(short)
		if f.Arg != "" {
			opt += ":"
		}
		opts = append(opts, opt)
	}
	for _, long := range f.Long {
		opt := "--" + escapeYashString(long)
		if f.Arg != "" {
			opt += ":"
		}
		opts = append(opts, opt)
	}

	description := escapeYashString(f.Description)

	return fmt.Sprintf("'%s; %s'",
		strings.Join(opts, " "), description)
}

type yash struct {
	Name string
	Opts []string
}

func newYash(c *Command) (y *yash, err error) {
	y = new(yash)
	y.Name = c.Name
	for _, flag := range c.Flags {
		y.Opts = append(y.Opts, toYashOpt(flag))
	}
	return y, nil
}

var yashCompletionTemplateText = `
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
`[1:]

func generateYashCompletion(w io.Writer, c *Command) error {
	tmpl, err := template.New("yash").Parse(yashCompletionTemplateText)
	if err != nil {
		return err
	}
	y, err := newYash(c)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, y)
}
