package utils

type Optional[T interface{}] struct {
	content T
	errors  []error
}

func (o *Optional[T]) IsExist() bool {
	return len(o.errors) == 0
}

func (o *Optional[T]) Errors() []error {
	return o.errors
}

func NewOptional[T interface{}](input T, err error) *Optional[T] {
	var errors = make([]error, 0)
	if err != nil {
		errors = append(errors, err)
	}

	return &Optional[T]{
		content: input,
		errors:  errors,
	}
}

func NilOptional[K interface{}, T interface{}](o *Optional[K], err error) *Optional[T] {
	return &Optional[T]{
		errors: append(o.errors, err),
	}
}

func ConvOptional[K interface{}, T interface{}](o *Optional[K]) *Optional[T] {
	return &Optional[T]{
		errors: o.errors,
	}
}

func MapOptional[T interface{}, K interface{}](o *Optional[T], fn func(in T) (K, error)) *Optional[K] {
	if o.IsExist() {
		return NewOptional(fn(o.content))
	}

	return ConvOptional[T, K](o)
}

func FlatMapOptional[T interface{}, K interface{}](o *Optional[T], fn func(in T) *Optional[K]) *Optional[K] {
	if o.IsExist() {
		return fn(o.content)
	}

	return ConvOptional[T, K](o)
}
