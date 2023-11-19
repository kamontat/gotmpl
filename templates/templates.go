package templates

import (
	"bytes"
	"os"
	"text/template"

	"github.com/kc-workspace/go-lib/mapper"
)

type Template struct {
	template *template.Template
}

func (t *Template) Name() string {
	return t.template.Name()
}

// ParseContent will convert template content to output string
func (t *Template) ParseContent(data mapper.Mapper) (string, error) {
	var buf bytes.Buffer
	var err = t.template.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// WriteFile will parse template and write to output file path
func (t *Template) WriteFile(data mapper.Mapper, output string) error {
	// Create output file
	file, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	return t.template.Execute(file, data)
}
