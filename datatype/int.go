package datatype

// ToInt will try to convert interface{} to integer
func ToInt(i interface{}) (int64, bool) {
	in, ok := i.(int)
	if ok {
		return int64(in), ok
	}

	i8, ok := i.(int8)
	if ok {
		return int64(i8), ok
	}

	i16, ok := i.(int16)
	if ok {
		return int64(i16), ok
	}

	i32, ok := i.(int32)
	if ok {
		return int64(i32), ok
	}

	i64, ok := i.(int64)
	if ok {
		return i64, ok
	}

	return -1, false
}
