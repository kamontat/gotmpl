package datatype

// ToFloat will try to convert interface{} to float
func ToFloat(i interface{}) (float64, bool) {
	f32, ok := i.(float32)
	if ok {
		return float64(f32), ok
	}

	f64, ok := i.(float64)
	if ok {
		return f64, ok
	}

	return -1, false
}
