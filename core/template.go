package core

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kamontat/gotmpl/fpath"
	"github.com/kc-workspace/go-lib/utils"
)

type Template struct {
	Input  string
	Output string
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
	var abs = fpath.Resolve(base, raw)
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
			panic("template file extension must be '.gotmpl'")
		}
		return append(result, abs)
	}
}

func buildOutputPath(dir, filename string) string {
	return filepath.Join(dir, strings.Replace(filename, TEMPLATE_EXTENSION, "", 1))
}

func buildTemplateOutput(inputs []string, output string, base string) (result []string) {
	result = make([]string, len(inputs))

	var isSingleMode = len(inputs) == 1
	var abs = fpath.Resolve(base, output)
	var isDir = filepath.Ext(abs) == ""

	if !isDir {
		if !isSingleMode {
			panic("template is directory, so output must be directory as well")
		}

		// both template and output are file
		result[0] = buildOutputPath(filepath.Dir(abs), filepath.Base(abs))

		return
	}

	for i, input := range inputs {
		result[i] = buildOutputPath(abs, filepath.Base(input))
	}

	return
}
