package utils

import "path/filepath"

// ResolvePath will return absolute path if input path is not absolute yet.
func ResolvePath(cwd, path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Clean(filepath.Join(cwd, path))
}
