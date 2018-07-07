package cpf

import (
	"github.com/brazilian-utils/golang/helpers"
)

func isValidCpf(cpf string) bool {
	return false
}

// Validate if a given string
// is a valid CPF
func Validate(cpf string) bool {
	cpf = helpers.OnlyNumbers(cpf)
	return isValidCpf(cpf)
}
