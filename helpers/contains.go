package helpers

// Contains verifies if a slice of string contains the given string
func Contains(h []string, n string) bool {
	for _, s := range h {
		if s == n {
			return true
		}
	}

	return false
}
