package core

import (
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/kamontat/gotmpl/config"
	"github.com/kamontat/gotmpl/utils"
)

type Core struct {
	Templates map[string][]*Template
	Config    *config.Config
}

func (c *Core) Parse() error {
	log.Printf(`
%s`, c.Config)

	for name, templates := range c.Templates {
		log.Printf("%s:\n", name)
		for _, _tmpl := range templates {
			log.Printf("  %v\n", _tmpl)

			// Read template file
			var inputContent = utils.NewOptional(os.ReadFile(_tmpl.Input))
			// Parse template file to Template object
			var tmpl = utils.MapOptional(inputContent, func(ctn []byte) (*template.Template, error) {
				return template.New(name).Parse(string(ctn))
			})
			// Create output file and execute template
			var result = utils.FlatMapOptional(tmpl, func(tmpl *template.Template) *utils.Optional[string] {
				var outputContent = utils.NewOptional(os.OpenFile(_tmpl.Output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666))
				return utils.MapOptional(outputContent, func(out *os.File) (string, error) {
					return "", tmpl.Funcs(sprig.FuncMap()).Execute(out, c.Config.GetData())
				})
			})

			if !result.IsExist() {
				log.Println("Errors: ")
				for i, err := range result.Errors() {
					log.Printf("  %d) %v", i, err)
				}
			}
		}

	}
	return nil
}
