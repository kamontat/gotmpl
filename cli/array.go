package main

type ArrayFlag []string

func (f *ArrayFlag) String() string {
	return "array flag"
}

func (f *ArrayFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}
