package datatype

// ToString will try to convert interface{} to string
func ToString(i interface{}) (string, bool) {
	s, ok := i.(string)
	return s, ok
}
