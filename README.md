acgen
=====

[![Build Status](https://travis-ci.org/kusabashira/acgen.svg?branch=master)](https://travis-ci.org/kusabashira/acgen)

Generate auto-completions.

Usage
-----

```
$ acgen [OPTION]... SOURCE
Generate auto-completions for specifyed TYPE
by SOURCE written in YAML.

Options:
  -t, --type=TYPE   output auto-completion for specified TYPE
                      TYPE=bash|zsh|fish|yash
  -h, --help        display this help text and exit
  -v, --version     output version information and exit
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

kusabashira <kusabashira227@gmail.com>
