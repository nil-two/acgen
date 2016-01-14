package acgen

import (
	"bytes"
	"fmt"
	"testing"
)

var exampleCommand = &Command{
	Name: "sed",
	Flags: []*Flag{
		&Flag{
			Short:       []string{"n"},
			Long:        []string{"quiet", "silent"},
			Description: "suppress automatic printing of pattern space",
		},
		&Flag{
			Short:       []string{"e"},
			Long:        []string{"expression"},
			Arg:         "script",
			Description: "add the script to the commands to be executed",
		},
		&Flag{
			Short:       []string{"f"},
			Long:        []string{"file"},
			Arg:         "script-file",
			Description: "add the contents of script-file to the commands to be executed",
		},
	},
}

func dumpCommand(c *Command) string {
	w := bytes.NewBuffer(make([]byte, 0))
	fmt.Fprintf(w, "  name = %q\n", c.Name)
	for i, flag := range c.Flags {
		fmt.Fprintf(w, "  flags[%d] = {%q, %q, %q, %q}\n",
			i, flag.Short, flag.Long, flag.Arg, flag.Description)
	}
	return w.String()
}

func TestGenerateBashCompletion(t *testing.T) {
	w := bytes.NewBuffer(make([]byte, 0))
	if err := generateBashCompletion(w, exampleCommand); err != nil {
		t.Errorf("generateBashCompletion returns %s, want nil\nsource:\n%s\n",
			err, dumpCommand(exampleCommand))
	}
	expect := `
_sed()
{
  local cur="${COMP_WORDS[COMP_CWORD]}"
  local opts='
    --quiet
    --silent
    --expression=
    --file=
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
complete -F _sed sed
`[1:]
	actual := w.String()
	if actual != expect {
		t.Errorf("generateBashCompletion returns:\n%s\nwant:\n%s\n",
			actual, expect)
	}
}

func TestGenerateZshCompletion(t *testing.T) {
	w := bytes.NewBuffer(make([]byte, 0))
	if err := generateZshCompletion(w, exampleCommand); err != nil {
		t.Errorf("generateZshCompletion returns %s, want nil\nsource:\n%s\n",
			err, dumpCommand(exampleCommand))
	}
	expect := `
#compdef sed
_arguments \
    '(-n --quiet --silent)'{'-n','--quiet','--silent'}'[suppress automatic printing of pattern space]' \
    '(-e --expression)'{'-e','--expression'}'[add the script to the commands to be executed]:script' \
    '(-f --file)'{'-f','--file'}'[add the contents of script-file to the commands to be executed]:script-file' \
    '*:input files:_files'
`[1:]
	actual := w.String()
	if actual != expect {
		t.Errorf("generateZshCompletion returns:\n%s\nwant:\n%s\n",
			actual, expect)
	}
}
