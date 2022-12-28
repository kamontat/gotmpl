package core

import (
	"github.com/kamontat/gotmpl/config"
	"github.com/kc-workspace/go-lib/logger"
)

func New(templates []string, conf *config.Config) *Core {
	var _templates = make(map[string][]*Template)
	for _, tmpl := range templates {
		var name, _template = NewTemplate(tmpl, conf.Setting)
		_templates[name] = _template
	}

	if len(_templates) < 1 {
		panic("values argument is required")
	}

	return &Core{
		Templates: _templates,
		Config:    conf,
	}
}

func NewTemplate(raw string, setting *config.Setting) (name string, templates []*Template) {
	templates = make([]*Template, 0)

	var log = logger.Get("core", "template")

	log.Debug("parsing raw template = '%s'", raw)
	__next, _output, err := split(raw, TEMPLATE_SEPARATOR)
	if err != nil {
		log.Error(err)
		return
	}

	_name, _input, err := split(__next, NAME_SEPARATOR)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug("parsed raw template = %s, %s, %s", _name, _input, _output)

	// Build name
	name = buildTemplateName(_name, _input)
	log.Debug("parsed name = %s", name)

	// Build input
	var input = buildTemplateInput(_input, setting.WorkingDirectory)
	log.Debug("parsed inputs = %v", input)

	// Use input to build output
	var output = buildTemplateOutput(input, _output, setting.WorkingDirectory)
	log.Debug("parsed outputs = %v", output)

	for i, in := range input {
		var out = output[i]

		templates = append(templates, &Template{
			Input:  in,
			Output: out,
		})
	}

	return
}
