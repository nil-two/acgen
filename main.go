package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/ogier/pflag"
)

var (
	name    = "acgen"
	version = "0.1.0"

	flag       = pflag.NewFlagSet(name, pflag.ContinueOnError)
	outputType = flag.StringP("type", "t", "", "")
	isHelp     = flag.BoolP("help", "h", false, "")
	isVersion  = flag.BoolP("version", "v", false, "")
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [OPTION]... FILE
Generate auto-completions for specifyed TYPE.

Options:
  -t, --type=TYPE   output auto-completion for specified TYPE
                      TYPE=bash|zsh|fish|yash
  -h, --help        display this help text and exit
  -v, --version     output version information and exit
`[1:], name)
}

func printVersion() {
	fmt.Fprintln(os.Stderr, version)
}

func printErr(err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
}

func guideToHelp() {
	fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n",
		name)
}

func main() {
	flag.SetOutput(ioutil.Discard)
	if err := flag.Parse(os.Args[1:]); err != nil {
		printErr(err)
		guideToHelp()
		os.Exit(2)
	}
	switch {
	case *isHelp:
		printUsage()
		os.Exit(0)
	case *isVersion:
		printVersion()
		os.Exit(0)
	}
	switch {
	case flag.NArg() < 1:
		printErr(fmt.Errorf("no input file"))
		guideToHelp()
		os.Exit(2)
	case *outputType == "":
		printErr(fmt.Errorf("no specify TYPE"))
		guideToHelp()
		os.Exit(2)
	case Generators[*outputType] == nil:
		printErr(fmt.Errorf("%s: is not supported", *outputType))
		guideToHelp()
		os.Exit(2)
	}

	file := flag.Arg(0)
	conf, err := ioutil.ReadFile(file)
	if err != nil {
		printErr(err)
		guideToHelp()
		os.Exit(2)
	}
	command := &Command{}
	if err = yaml.Unmarshal(conf, command); err != nil {
		printErr(err)
		guideToHelp()
		os.Exit(2)
	}

	if err = Generators[*outputType](os.Stdout, command); err != nil {
		printErr(err)
		os.Exit(1)
	}
}
