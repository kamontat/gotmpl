package core

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/kamontat/gotmpl/config"
	"github.com/kamontat/gotmpl/utils"
)

func New(templates []string, conf *config.Config) *Core {
	var _templates = make(map[string]*Template)
	for _, tmpl := range templates {
		var _template = NewTemplate(tmpl, conf.Setting)
		_templates[_template.Name] = _template
	}

	if len(_templates) < 1 {
		log.Panic("cannot be empty template")
	}

	return &Core{
		Templates: _templates,
		Config:    conf,
	}
}

func NewTemplate(input string, setting *config.Setting) *Template {
	next, _output, err := getValues(input, TEMPLATE_SEPARATOR, func(in string) string {
		return strings.Replace(in, TEMPLATE_EXTENSION, "", 1)
	})
	if err != nil {
		log.Panic(err)
	}

	name, _template, err := getValues(next, NAME_SEPARATOR, func(in string) string {
		return utils.ResolvePath(setting.WorkingDirectory, in)
	})
	if err != nil {
		log.Panic(err)
	}

	var template = utils.ResolvePath(setting.WorkingDirectory, _template)
	var output = utils.ResolvePath(setting.WorkingDirectory, _output)

	// Use template file name if user specify only output directory
	if filepath.Ext(output) == "" {
		output = filepath.Join(output, filepath.Base(strings.Replace(template, TEMPLATE_EXTENSION, "", 1)))
	}

	// [Validate] template must be .gotmpl extension
	if filepath.Ext(template) != TEMPLATE_EXTENSION {
		log.Panic("template file extension must be '.gotmpl'")
	}

	return &Template{
		Name:   name,
		Input:  template,
		Output: output,
	}
}
