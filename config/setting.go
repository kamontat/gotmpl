package config

import "fmt"

type Setting struct {
	WorkingDirectory string
	DebugMode        bool
}

func (s *Setting) String() string {
	return fmt.Sprintf(`setting:
  wd: %s
  debug: %t
`, s.WorkingDirectory, s.DebugMode)
}
