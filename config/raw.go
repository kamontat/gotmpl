package config

import (
	"log"
	"strings"

	"github.com/kamontat/gotmpl/maps"
)

func parseRawData(raw []string) (result map[string]interface{}) {
	result = make(map[string]interface{})
	for _, raw := range raw {
		var array = strings.Split(raw, "=")
		if len(array) != 2 {
			log.Panic("cannot convert raw data to map")
		}

		var key = array[0]
		var value = array[1]
		maps.Set(result, key, value)
	}

	return
}
