package helpers

import (
	"regexp"
)

// OnlyNumbers removes all non
// numeric characters from a string
func OnlyNumbers(numbers string) string {
	return regexp.MustCompile("[^0-9]").ReplaceAllString(numbers, "")
}
