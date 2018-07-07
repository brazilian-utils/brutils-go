package helpers

import (
	"regexp"
	"strings"
)

// OnlyNumbers removes all non
// numeric characters from a string
func OnlyNumbers(numbers string) string {
	numericStr := regexp.MustCompilePOSIX("[0-9]+").FindAllString(numbers, -1)
	return strings.Join(numericStr[:], "")
}
