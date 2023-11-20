# Template engine

Parse template using input data file (or options)

```
# Kamontat's Template parser

$ kc-tpr <options>...

Options:
	-cwd string
		base directory relative path resolve to (default "$PWD").
	-template-path string
		a require template path (either absolute or relative is accepted)
		if directory is passed, parse all template on that folder and
		assume output to be directory with same input name
		without template extension.
	-data-paths string
		a data path (either yaml or json). You can pass more than 1 time.
	-data string
		a data string in format <key>=<value>. You can pass more than 1 time.
	-output-path string
		a optional output path (either absolute or relative is accepted).
		if not provided, output to stdout.
	-debug
		enabled debug information.
```

## Development

```bash
go run github.com/kamontat/tmpl/cli
```

## asdf plugin

A easy way to install using asdf: https://github.com/kc-workspace/asdf-kc-tpr
