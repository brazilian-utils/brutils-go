package cep

import (
	"github.com/brazilian-utils/brutils-go/helpers"
)

// Every CEP has exactly 8 characters
const cepSize = 8

// IsValid validates if a given CEP is valid
func IsValid(cep string) bool {
	digits := helpers.OnlyNumbers(cep)

	return hasValidLength(digits)
}

// Validates the string length
func hasValidLength(cep string) bool {
	return len(cep) == cepSize
}
