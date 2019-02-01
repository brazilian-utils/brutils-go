package cep

import (
	"github.com/rafa-acioly/brutils-go/helpers"
)

// Size represents the length of a valid cep
// without special characters
const Size = 8

// IsValid retrieves if a CEP is valid
// removing the special characters.
func IsValid(cep string) bool {
	return len(helpers.OnlyNumbers(cep)) == Size
}
