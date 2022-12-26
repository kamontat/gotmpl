package template

import (
	"log"
	"os"
	"text/template"

	"github.com/kamontat/gotmpl/config"
)

type Template struct {
	underlay *template.Template
}

func (t *Template) Parse(filepath string, conf *config.Config) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.underlay.Execute(file, conf.Variable())
	if err != nil {
		log.Fatalln(err)
	}
}

func New(filepath string) *Template {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl := template.Must(template.New("default").Parse(string(content)))
	return &Template{
		underlay: tmpl,
	}
}
