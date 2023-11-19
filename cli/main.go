package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kamontat/tmpl/data"

	"github.com/kamontat/tmpl/paths"
	"github.com/kamontat/tmpl/templates"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/utils"
)

var cwd string
var debug bool
var dataPaths ArrayFlag
var dataStrings ArrayFlag
var templatePath string
var outputPath string

func main() {
	var tmpl, err = templates.NewFile(paths.Resolve(cwd, templatePath))
	if err != nil {
		panic(err)
	}

	d, err := data.New(paths.Resolves(cwd, dataPaths), dataStrings)
	if err != nil {
		panic(err)
	}

	if outputPath == "" {
		content, err := tmpl.ParseContent(d)
		if err != nil {
			panic(err)
		}
		fmt.Print(content)
	} else {
		err = tmpl.WriteFile(d, paths.Resolve(cwd, outputPath))
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	flag.StringVar(&cwd, "cwd", utils.MustR(os.Getwd()), "")
	flag.BoolVar(&debug, "debug", false, "")

	flag.Var(&dataPaths, "data-paths", "")
	flag.Var(&dataStrings, "data", "")

	flag.StringVar(&templatePath, "template-path", "", "")
	flag.StringVar(&outputPath, "output-path", "", "")

	flag.Usage = usage

	flag.Parse()

	if debug {
		logger.DefaultManager.SetLevel(logger.DEBUG)
	}
}
