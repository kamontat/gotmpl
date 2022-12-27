package utils

import (
	"errors"
	"log"
	"strings"
)

func DeepMergeMap(a, b map[interface{}]interface{}) map[interface{}]interface{} {
	out := make(map[interface{}]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		// If you use map[string]interface{}, ok is always false here.
		// Because yaml.Unmarshal will give you map[interface{}]interface{}.
		if v, ok := v.(map[interface{}]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[interface{}]interface{}); ok {
					out[k] = DeepMergeMap(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

// str is <key>=<value> format string
func ToMap(str []string) map[interface{}]interface{} {
	var result = make(map[interface{}]interface{})
	for _, raw := range str {
		var array = strings.Split(raw, "=")
		if len(array) != 2 {
			log.Panic(errors.New("cannot parse input string to map"))
		}

		var key = array[0]
		var value = array[1]

		// TODO: support nested key using dotnotation
		result[key] = value
	}

	return result
}
