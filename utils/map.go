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
	var movement = make(map[interface{}]interface{})
	var result = movement
	for _, raw := range str {
		var array = strings.Split(raw, "=")
		if len(array) != 2 {
			log.Panic(errors.New("cannot parse input string to map"))
		}

		var key = array[0]
		var value = array[1]

		// Source: https://github.com/kamontat/fthelper/blob/4970bd51fbd41418187dd47c4e5710bd04e4241d/shared/maps/utils.go#L27-L49
		var keys = strings.Split(key, ".")
		var length = len(keys)
		for i, k := range keys {
			if i == length-1 {
				if value == "" {
					delete(movement, k)
				} else {
					movement[k] = value
				}
			} else {
				v, ok := movement[k].(map[interface{}]interface{})
				if !ok {
					movement[k] = make(map[interface{}]interface{})
					movement = movement[k].(map[interface{}]interface{})
				} else {
					movement = v
				}
			}
		}
	}

	return result
}
