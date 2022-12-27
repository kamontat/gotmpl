package core

import (
	"log"

	"github.com/kamontat/gotmpl/config"
)

func New(templates []string, conf *config.Config) *Core {
	var _templates = make(map[string][]*Template)
	for _, tmpl := range templates {
		var name, _template = NewTemplate(tmpl, conf.Setting)
		_templates[name] = _template
	}

	if len(_templates) < 1 {
		log.Panic("cannot be empty template")
	}

	return &Core{
		Templates: _templates,
		Config:    conf,
	}
}

func NewTemplate(raw string, setting *config.Setting) (string, []*Template) {
	__next, _output, err := split(raw, TEMPLATE_SEPARATOR)
	if err != nil {
		log.Panic(err)
	}

	_name, _input, err := split(__next, NAME_SEPARATOR)
	if err != nil {
		log.Panic(err)
	}

	// Build name
	var name = buildTemplateName(_name, _input)
	// Build input
	var input = buildTemplateInput(_input, setting.WorkingDirectory)
	// Use input to build output
	var output = buildTemplateOutput(input, _output, setting.WorkingDirectory)

	var templates = make([]*Template, 0)
	for i, in := range input {
		var out = output[i]

		templates = append(templates, &Template{
			Input:  in,
			Output: out,
		})
	}

	return name, templates
}
