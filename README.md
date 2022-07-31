acgen
=====

Generate auto-completions.

Usage
-----

```
$ acgen [OPTION]... SOURCE
Generate auto-completions for specifyed TYPE
by SOURCE written in YAML.

Options:
  -t, --type=TYPE   output auto-completion for specified TYPE
                      TYPE=bash|zsh|fish|yash|tcsh
  -h, --help        display this help text and exit
  -v, --version     output version information and exit
```

Installation
------------

### compiled binary

See [releases](https://github.com/nil2nekoni/acgen/releases)

### go get

```
go get github.com/nil2nekoni/acgen/cmd/acgen
```

Options
-------

### --help

Display a help message.

### --version

Output the version of acgen.

### -t, --type=TYPE

Specify the target shell.

acgen is currently supporting `bash` `zsh` `fish` `yash` `tcsh`.

```
$ cat cmd.yml
name: cmd
flags:
  - long: ['source']
  - long: ['destination']

$ acgen --type=fish cmd.yml
complete -c 'cmd' -l 'source' -d ''
complete -c 'cmd' -l 'destination' -d ''

$ acgen --type=tcsh cmd.yml
complete cmd 'c/--/(source destination)/'
```

Source
------

Describe in YAML.
Rough structures as follows:

```yaml
name: '<program name>'
flags:
  - short: ['<short option>', '<short option>', ...]
    long:  ['<long option>', '<long option>', ...]
    arg: '<argument>'
    description: '<description>'

  - short: ['<short option>', '<short option>', ...]
    long:  ['<long option>', '<long option>', ...]
    arg: '<argument>'
    description: '<description>'

  ...
```

#### name

`name` is a comamnd's name such as `cat` and `sed`.

#### short

`short` are short options
such as `n` and `e`.

Header hyphen must be removed.

#### long

`long` are long options
such as `quiet` and `script`.

Header hyphen must be removed.

#### arg

`arg` is a arguments for flag
such as `script-file` for `file`.

If this is ommitted, the flag interpreted as no argument flag.

#### description

`description` is a description for flag
such as `add the script to the ...` for `script`.

### example

```yaml
# subset of sed
---
name: sed
flags:
  - short: ['n']
    long: ['quiet', 'silent']
    description: 'suppress automatic printing of pattern space'

  - short: ['e']
    long: ['expression']
    arg: 'script'
    description: 'add the script to the commands to be executed'

  - short: ['f']
    long: ['file']
    arg: 'script-file'
    description: 'add the contents of script-file to the commands to be executed'
```

License
-------

MIT License

Author
------

nil2 <nil2@nil2.org>
