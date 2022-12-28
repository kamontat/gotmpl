package config

import (
	"strings"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

func parseRawData(input []string) (result map[string]interface{}) {
	var log = logger.Get("config", "raw-data")
	result = make(map[string]interface{})
	for _, raw := range input {
		var array = strings.Split(raw, "=")
		if len(array) != 2 {
			log.ErrorString("cannot convert %s to valid config, skipping", raw)
		}

		var key = array[0]
		var value = array[1]
		mapper.Set(result, key, value)
	}

	return
}
