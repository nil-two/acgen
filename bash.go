package main

import (
	"io"
	"strings"
	"text/template"
)

func init() {
	Generators["bash"] = GenerateBashCompletion
}

type Bash struct {
	Name string
	Opts []string
}

func NewBash(c *Config) (b *Bash, err error) {
	b = new(Bash)
	b.Name = c.Name
	for _, flag := range c.Flags {
		for _, long := range flag.Long {
			opt := "--" + strings.Replace(long, `'`, `'"'"'`, -1)
			if flag.Arg != "" {
				opt += "="
			}
			b.Opts = append(b.Opts, opt)
		}
	}
	return b, nil
}

var BashCompletionTemplateText = `
_{{.Name}}()
{
  local cur="${COMP_WORDS[COMP_CWORD]}"
  local opts='{{range .Opts}}
    {{.}}{{end}}
  '
  case "$cur" in
    -*)
      COMPREPLY=( $(compgen -W "$opts" -- "$cur") )
      ;;
    *)
      _filedir
      ;;
  esac
  [[ ${COMPREPLY[0]} == *= ]] && compopt -o nospace
}
complete -F _{{.Name}} {{.Name}}
`[1:]

func GenerateBashCompletion(w io.Writer, c *Config) error {
	tmpl, err := template.New("bash").Parse(BashCompletionTemplateText)
	if err != nil {
		return err
	}
	b, err := NewBash(c)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, b)
}
