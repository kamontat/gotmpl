package fpath

import "path/filepath"

// Resolve will return absolute path if input path is not absolute yet.
func Resolve(cwd, path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Clean(filepath.Join(cwd, path))
}
