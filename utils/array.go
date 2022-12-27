package utils

// CloneArray will create copy of input array
// https://github.com/kamontat/fthelper/blob/4970bd51fbd41418187dd47c4e5710bd04e4241d/shared/utils/array.go#L4-L11
func CloneArray[T interface{}](a []T, extra ...T) []T {
	var base = make([]T, 0)

	base = append(base, a...)
	base = append(base, extra...)

	return base
}
