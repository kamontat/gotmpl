# Usage

```
$ gotmpl [-config=<path>] [-debug] value [values...]

value syntax: [<name>:]<template_path>[=<output_path>]
template syntax: <file_name>.<file_extension>.gotmpl
	- /tmp/template.yaml.gotmpl
	- /tmp/template.yaml.gotmpl=/tmp/output.yaml
	- ./template.yaml.gotmpl=./output.yaml
	- template:./template.yaml.gotmpl
	- template:./template.yaml.gotmpl=./output.yaml

By default if no output specify, it will output to template directory.

  -cwd string
    	current directory for relative path resolve to (default "/Users/natcha/Desktop/gotmpl")
  -data value
    	data files, either yaml or json (you can pass more than 1 times)
  -debug
    	enable debug information
  -raw value
    	raw data in format <key>=<value> (you can pass more than 1 times)
```

## Development

```bash
go run github.com/kamontat/gotmpl/cli
```
