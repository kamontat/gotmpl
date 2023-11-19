package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/utils"
	"gopkg.in/yaml.v2"
)

func New(files []string, data []string) (mapper.Mapper, error) {
	fileData, err := LoadFile(files...)
	if err != nil {
		return mapper.New(), err
	}
	rawData, err := LoadRaw(data...)
	if err != nil {
		return mapper.New(), err
	}

	return mapper.Merge(fileData, rawData, mapper.New()), nil
}

func LoadFile(paths ...string) (mapping mapper.Mapper, err error) {
	mapping = mapper.New()
	for _, path := range paths {
		var ext = filepath.Ext(path)
		var content = utils.MustR(os.ReadFile(path))
		if ext == ".yaml" || ext == ".yml" {
			utils.Must(yaml.Unmarshal(content, &mapping))
		} else if ext == ".json" || ext == ".json5" {
			var _mapping = make(map[string]interface{})
			utils.Must(json.Unmarshal(content, &_mapping))
			for k, v := range _mapping {
				mapping[k] = v
			}
		} else {
			err = fmt.Errorf("invalid data file extension (%s)", ext)
		}
	}
	return
}

func LoadRaw(data ...string) (mapping mapper.Mapper, err error) {
	var log = logger.DefaultManager.New("data", "raw")
	mapping = mapper.New()
	for _, raw := range data {
		var array = strings.Split(raw, "=")
		if len(array) != 2 {
			log.Errorf("cannot convert %s to valid config, skipping", raw)
		}

		var key = array[0]
		var value = array[1]
		mapper.Set(mapping, key, value)
	}

	return
}
