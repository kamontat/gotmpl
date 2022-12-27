package datatype

// ToBool will try to convert interface{} to boolean
func ToBool(i interface{}) (bool, bool) {
	b, ok := i.(bool)
	if ok {
		return b, ok
	}

	return false, false
}
