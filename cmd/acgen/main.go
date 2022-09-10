package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/nil-two/acgen"
	"github.com/ogier/pflag"
)

var (
	cmdName    = "acgen"
	cmdVersion = "0.1.0"

	flagset    = pflag.NewFlagSet(cmdName, pflag.ContinueOnError)
	outputType = flagset.StringP("type", "t", "", "")
	isHelp     = flagset.BoolP("help", "h", false, "")
	isVersion  = flagset.BoolP("version", "v", false, "")
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `
Usage: %s [OPTION]... SOURCE
Generate auto-completions for specifyed TYPE
by SOURCE written in YAML.

Options:
  -t, --type=TYPE   output auto-completion for specified TYPE
                      TYPE=bash|zsh|fish|yash|tcsh
  -h, --help        display this help text and exit
  -v, --version     output version information and exit
`[1:], cmdName)
}

func printVersion() {
	fmt.Fprintln(os.Stderr, cmdVersion)
}

func printErr(err interface{}) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", cmdName, err)
}

func guideToHelp() {
	fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", cmdName)
}

func main() {
	flagset.SetOutput(ioutil.Discard)
	if err := flagset.Parse(os.Args[1:]); err != nil {
		printErr(err)
		guideToHelp()
		os.Exit(2)
	}
	if *isHelp {
		printUsage()
		os.Exit(0)
	}
	if *isVersion {
		printVersion()
		os.Exit(0)
	}

	if flagset.NArg() < 1 {
		printErr("no input file")
		guideToHelp()
		os.Exit(2)
	}
	if *outputType == "" {
		printErr("no specify TYPE")
		guideToHelp()
		os.Exit(2)
	}

	generate, err := acgen.LookGenerator(*outputType)
	if err != nil {
		printErr(err)
		guideToHelp()
		os.Exit(2)
	}

	source, err := ioutil.ReadFile(flagset.Arg(0))
	if err != nil {
		printErr(err)
		guideToHelp()
		os.Exit(2)
	}
	command := &acgen.Command{}
	if err = yaml.Unmarshal(source, command); err != nil {
		printErr(err)
		guideToHelp()
		os.Exit(2)
	}

	if err = generate(os.Stdout, command); err != nil {
		printErr(err)
		os.Exit(1)
	}
}
