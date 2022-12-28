package core

import (
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/kamontat/gotmpl/config"
	"github.com/kc-workspace/go-lib/logger"
)

type Core struct {
	Templates map[string][]*Template
	Config    *config.Config
}

func (c *Core) Parse() error {
	var log = logger.Get("core", "parser")

	// Reuse data for all templates
	var data = c.Config.GetData()
	for name, templates := range c.Templates {
		log.Debug("parsing %s...", name)
		for _, _tmpl := range templates {
			log.Debug("  input: '%s'", _tmpl.Input)
			log.Debug("  output: '%s'", _tmpl.Output)

			// Read template file
			inputContent, err := os.ReadFile(_tmpl.Input)
			if err != nil {
				log.Error(err)
				break
			}

			// Parse template file to Template object
			tmpl, err := template.New(name).Parse(string(inputContent))
			if err != nil {
				log.Error(err)
				break
			}

			// Create output file
			outputContent, err := os.OpenFile(_tmpl.Output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				log.Error(err)
				break
			}

			// execute template
			err = tmpl.Funcs(sprig.FuncMap()).Execute(outputContent, data)
			if err != nil {
				log.Error(err)
				break
			}
		}

	}
	return nil
}
