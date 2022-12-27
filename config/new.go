package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/kamontat/gotmpl/utils"
	"gopkg.in/yaml.v3"
)

func New(data []string, raw []string, setting *Setting) *Config {
	return &Config{
		data:    NewData(setting.WorkingDirectory, data, raw),
		Setting: setting,
	}
}

func NewData(base string, files, raw []string) *data {
	var underlay = make(map[string]map[string]interface{})
	for _, file := range files {
		var ext = filepath.Ext(file)
		var abspath = utils.ResolvePath(base, file)

		var mapping = make(map[string]interface{})

		var content = utils.MustR(os.ReadFile(abspath))
		if ext == ".yaml" || ext == ".yml" {
			utils.Must(yaml.Unmarshal(content, &mapping))
		} else if ext == ".json" || ext == ".json5" {
			var _mapping = make(map[string]interface{})
			utils.Must(json.Unmarshal(content, &_mapping))
			for k, v := range _mapping {
				mapping[k] = v
			}
		} else {
			log.Panicf("invalid file extension %s\n", abspath)
		}

		underlay[abspath] = mapping
	}

	if len(raw) > 0 {
		underlay["raw"] = parseRawData(raw)
	}

	return &data{
		underlay: underlay,
	}
}
