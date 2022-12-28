package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/kamontat/gotmpl/fpath"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/utils"
	"gopkg.in/yaml.v3"
)

func New(data []string, raw []string, setting *Setting) (config *Config) {
	var log = logger.Get("config", "new")

	config = &Config{
		Data:    NewData(setting.WorkingDirectory, data, raw),
		Setting: setting,
	}

	log.Debug(`data: %s`, config.Data.String())
	log.Debug(`setting: wd=%s, debug=%t`, config.Setting.WorkingDirectory, config.Setting.DebugMode)

	return
}

func NewData(base string, files, raw []string) *data {
	var log = logger.Get("config", "data")
	var underlay = make(map[string]map[string]interface{})
	for _, file := range files {
		var ext = filepath.Ext(file)
		var abspath = fpath.Resolve(base, file)

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
			log.ErrorString(`cannot parse %s file, skipping`, abspath)
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
