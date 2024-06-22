package ini

import (
	"strings"
)

var (
	eq     func(string, string) bool          = strings.EqualFold
	trimS  func(string) string                = strings.TrimSpace
	trim   func(string, string) string        = strings.Trim
	splitn func(string, string, int) []string = strings.SplitN
	hasPfx func(string, string) bool          = strings.HasPrefix
	hasSfx func(string, string) bool          = strings.HasSuffix
)
