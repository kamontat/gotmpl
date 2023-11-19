package paths

import (
	"path/filepath"
)

// Resolve will return absolute path if input paths is not absolute
// or join input paths if it already absolute.
func Resolve(cwd string, paths ...string) string {
	var path = filepath.Join(paths...)
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(cwd, path)
}

func Resolves(cwd string, paths []string) []string {
	var output = make([]string, len(paths))
	for i, path := range paths {
		output[i] = Resolve(cwd, path)
	}
	return output
}
