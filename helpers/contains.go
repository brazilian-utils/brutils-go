package helpers

import "strings"

// Contains verifies if a slice of string contains the given string
func Contains(h []string, n string) bool {
	line := strings.Join(h, "")
	return strings.Contains(line, n)
}
