package core

import (
	"fmt"
	"strings"
)

func split(input, separator string) (string, string, error) {
	var values = strings.Split(input, separator)
	var length = len(values)
	if length == 1 {
		return values[0], values[0], nil
	} else if length == 2 {
		return values[0], values[1], nil
	} else {
		return "", "", fmt.Errorf("invalid input [sep=%s, input=%s]", separator, input)
	}
}
