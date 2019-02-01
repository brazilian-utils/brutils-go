package cpf

import (
	"strconv"
	"strings"

	"github.com/brazilian-utils/brutils-go/helpers"
)

// blacklist is a list of invalid (blacklisted) values for CPF
var blacklist = []string{
	"00000000000",
	"11111111111",
	"22222222222",
	"33333333333",
	"44444444444",
	"55555555555",
	"66666666666",
	"77777777777",
	"88888888888",
	"99999999999",
}

// verifierIndexes are the position of verification digits
var verifierIndexes = []int{9, 10}

// CPFSize represents the valid size
// for a CPF ignoring the special characters
const CPFSize = 11

// IsValid validates if a given CPF is valid
func IsValid(cpf string) bool {
	cpfNumbers := helpers.OnlyNumbers(cpf)

	return hasValidLength(cpfNumbers) && !isBlacklisted(cpf) && isValidChecksum(cpfNumbers)
}

// isValidChecksum performs checksum validation
func isValidChecksum(cpf string) bool {
	validity := true

	for _, verifier := range verifierIndexes {
		mod := computeMod(strings.Split(cpf[:verifier], ""))

		valid, _ := strconv.Atoi(string(cpf[verifier]))
		validity = validity && (valid == mod)
	}

	return validity
}

// Compute the mod for the current slice of strings
func computeMod(digits []string) (res int) {
	weight := len(digits) + 1

	var mod int
	for _, digitStr := range digits {
		digit, _ := strconv.Atoi(digitStr)
		mod += digit * weight
		weight--
	}

	if mod = mod % 11; mod >= 2 {
		res = 11 - mod
	}

	return
}

// hasValidLength verify if the given CPF has valid length
func hasValidLength(cpf string) bool {
	return len(cpf) == CPFSize
}

func isBlacklisted(cpf string) bool {
	return helpers.Contains(blacklist, cpf)
}
