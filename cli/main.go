package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kamontat/gotmpl/config"
	"github.com/kamontat/gotmpl/core"
	"github.com/kamontat/gotmpl/logger"
	"github.com/kamontat/gotmpl/utils"
)

var defaultCwd = utils.MustR(os.Getwd())
var defaultDebug = false

var data ArrayFlag
var rawData ArrayFlag
var cwd string
var debug bool

func usage() string {
	return fmt.Sprintf(`# Go Template
$ gotmpl [options...] value [values...]

## Options
  --cwd string [default=%s]
		base path for resolve relative path
  --data path [multiple]
        data files to fill information on template.
		either on yaml or json format.
  --raw <key>=<value> [multiple]
	    raw data to fill information on template.
		dot-notation are accepted.
  --debug [default=%t]
		enable debug information and logging

## Value
syntax: 
	[<name>:]<template_path>[=<output_path>]
name: 
	any string are accepted.
	default to template_path if not provided.
template: 
	relative or absolute are accepted.
	relative will resolve by using --cwd option.
	must have %s as extension; otherwise will throw error.
output:
	relative or absolute are accepted.
	directory or filename are accepted as well.
	default to template file without template extension,
	if only directory provided, use template file as output file.

### Example
  '/tmp/dir/gh.txt.gotmpl'
	name: /tmp/dir/gh.txt.gotmpl
	template: /tmp/dir/gh.txt.gotmpl
	output: /tmp/dir/gh.txt
  '/tmp/content.md.gotmpl=/tmp/readme.txt'
	name: /tmp/content.md.gotmpl
	template: /tmp/content.md.gotmpl
	output: /tmp/readme.txt
  './file.json.gotmpl=./untitled.json'
	name: ./file.json.gotmpl
	template: /$CWD/file.json.gotmpl
	output: /$CWD/untitled.json
  'default:config.yaml.gotmpl'
	name: default
	template: /$CWD/config.yaml.gotmpl
	output: /$CWD/config.yaml
  'custom:values.yaml.gotmpl=./output/values.yaml'
	name: custom
	template: /$CWD/values.yaml.gotmpl
	output: /$CWD/output/values.yaml
  'custom:values.yaml.gotmpl=output'
	name: custom
	template: /$CWD/values.yaml.gotmpl
	output: /$CWD/output/values.yaml

`, defaultCwd, defaultDebug, core.TEMPLATE_EXTENSION)
}

func main() {
	flag.Var(&data, "data", "")
	flag.Var(&rawData, "raw", "")
	flag.StringVar(&cwd, "cwd", defaultCwd, "")
	flag.BoolVar(&debug, "debug", defaultDebug, "enable debug information")
	flag.Usage = func() {
		fmt.Print(usage())
	}

	flag.Parse()

	logger.Setup(debug)
	var conf = config.New(data, rawData, &config.Setting{
		WorkingDirectory: cwd,
		DebugMode:        debug,
	})

	var c = core.New(flag.Args(), conf)
	utils.Must(c.Parse())
}
