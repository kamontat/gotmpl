package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kamontat/gotmpl/utils"
)

type Template struct {
	Input  string
	Output string
}

func (t *Template) String() string {
	return fmt.Sprintf("%s => %s", t.Input, t.Output)
}

const (
	NAME_SEPARATOR     = ":"
	TEMPLATE_SEPARATOR = "="
	TEMPLATE_EXTENSION = ".gotmpl"
)

func buildTemplateName(name string, input string) string {
	if name == input {
		return name
	}

	return name
}

func buildTemplateInput(raw, base string) (result []string) {
	result = make([]string, 0)
	var abs = utils.ResolvePath(base, raw)
	var stat = utils.MustR(os.Stat(abs))

	if stat.IsDir() {
		var files = utils.MustR(os.ReadDir(abs))
		for _, file := range files {
			var filename = file.Name()
			if filepath.Ext(filename) == TEMPLATE_EXTENSION {
				result = append(result, filepath.Join(abs, filename))
			}
		}

		return result
	} else {
		// Input single file
		if filepath.Ext(abs) != TEMPLATE_EXTENSION {
			log.Panic("template file extension must be '.gotmpl'")
		}
		return append(result, abs)
	}
}

func buildTemplateOutput(inputs []string, output string, base string) (result []string) {
	result = make([]string, len(inputs))

	var isSingleMode = len(inputs) == 1
	var abs = utils.ResolvePath(base, output)
	var isDir = filepath.Ext(abs) == ""

	if !isDir {
		if !isSingleMode {
			log.Panic("template is directory, so output must be directory as well")
		}

		// both template and output are file
		result[0] = abs
		return
	}

	for i, input := range inputs {
		var base = strings.Replace(filepath.Base(input), TEMPLATE_EXTENSION, "", 1)
		var output = filepath.Join(abs, base)
		result[i] = output
	}

	return
}
