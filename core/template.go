package core

import (
	"fmt"
	"strings"
)

type Template struct {
	Name   string
	Input  string
	Output string
}

const (
	NAME_SEPARATOR     = ":"
	TEMPLATE_SEPARATOR = "="
	TEMPLATE_EXTENSION = ".gotmpl"
)

func getValues(input, separator string, fn func(in string) string) (string, string, error) {
	var values = strings.Split(input, separator)
	var length = len(values)
	if length == 1 {
		return values[0], fn(values[0]), nil
	} else if length == 2 {
		return values[0], values[1], nil
	} else {
		return "", "", fmt.Errorf("invalid input [sep=%s, input=%s]", separator, input)
	}
}
