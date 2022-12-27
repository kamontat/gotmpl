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

var data ArrayFlag
var rawData ArrayFlag
var cwd string
var debug bool

func usage() string {
	return `$ gotmpl [-config=<path>] [-debug] value [values...]

value syntax: [<name>:]<template_path>[=<output_path>]
template syntax: <file_name>.<file_extension>.gotmpl
	- /tmp/template.yaml.gotmpl
	- /tmp/template.yaml.gotmpl=/tmp/output.yaml
	- ./template.yaml.gotmpl=./output.yaml
	- template:./template.yaml.gotmpl
	- template:./template.yaml.gotmpl=./output.yaml

By default if no output specify, it will output to template directory.

`
}

func main() {
	flag.Var(&data, "data", "data files, either yaml or json (you can pass more than 1 times)")
	flag.Var(&rawData, "raw", "raw data in format <key>=<value> (you can pass more than 1 times)")
	flag.StringVar(&cwd, "cwd", utils.MustR(os.Getwd()), "current directory for relative path resolve to")
	flag.BoolVar(&debug, "debug", false, "enable debug information")
	flag.Usage = func() {
		fmt.Print(usage())
		flag.PrintDefaults()
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
