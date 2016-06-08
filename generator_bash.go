package acgen

import (
	"io"
	"strings"
	"text/template"
)

func init() {
	RegisterGenerator("bash", generateBashCompletion)
}

type bash struct {
	Name string
	Opts []string
}

func escapeBashString(s string) string {
	return strings.Replace(s, `'`, `'"'"'`, -1)
}

func newBash(c *Command) (b *bash, err error) {
	b = new(bash)
	b.Name = c.Name
	for _, flag := range c.Flags {
		for _, long := range flag.Long {
			opt := "--" + escapeBashString(long)
			if flag.Arg != "" {
				opt += "="
			}
			b.Opts = append(b.Opts, opt)
		}
	}
	return b, nil
}

var bashTemplate = template.Must(template.New("bash").Parse(`
_{{.Name}}()
{
  local cur=${COMP_WORDS[COMP_CWORD]}
  local opts='{{range .Opts}}
    {{.}}{{end}}
  '
  case $cur in
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
`[1:]))

func generateBashCompletion(w io.Writer, c *Command) error {
	b, err := newBash(c)
	if err != nil {
		return err
	}
	return bashTemplate.Execute(w, b)
}
