package templates

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func NewContent(path string, content string) (*Template, error) {
	tmpl, err := template.New(path).Funcs(sprig.FuncMap()).Parse(content)
	if err != nil {
		return nil, err
	}

	return &Template{
		template: tmpl,
	}, nil
}

func NewFile(filepaths ...string) (*Template, error) {
	tmpl, err := template.ParseFiles(filepaths...)
	if err != nil {
		return nil, err
	}

	return &Template{
		template: tmpl,
	}, nil
}
