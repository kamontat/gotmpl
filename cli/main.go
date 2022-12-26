package main

import (
	"flag"
	"log"

	"github.com/kamontat/gotmpl/config"
	"github.com/kamontat/gotmpl/template"
)

var configType string
var configPath string
var templatePath string
var outputPath string

func main() {
	flag.StringVar(&configPath, "config", "", "config file path")
	flag.StringVar(&configType, "config-type", "yaml", "config type")
	flag.StringVar(&templatePath, "template", "", "template file path")
	flag.StringVar(&outputPath, "output", "", "output file path")

	flag.Parse()

	log.Printf(`Input: 
1. config-path:   %s
2. template-path: %s
3. output-path:   %s
`, configPath, templatePath, outputPath)

	var conf = config.New(configType, configPath)
	var tmpl = template.New(templatePath)

	tmpl.Parse(outputPath, conf)
}
