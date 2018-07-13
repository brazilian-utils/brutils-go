package cpf

import (
	"strconv"
	"strings"

	"github.com/brazilian-utils/brutils-go/helpers"
)

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

var verifierIndexes = []int{9, 10}

// IsValid validates if a given CPF is valid
func IsValid(cpf string) bool {
	cpfNumbers := helpers.OnlyNumbers(cpf)

	return isValidChecksum(cpfNumbers) && !isBlacklisted(cpf)
}

func isValidChecksum(cpf string) bool {
	validity := true

	for _, verifier := range verifierIndexes {
		mod := computeMod(strings.Split(cpf[:verifier], ""))

		valid, _ := strconv.Atoi(string(cpf[verifier]))
		validity = validity && (valid == mod)
	}

	return validity
}

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

func isBlacklisted(cpf string) bool {
	return helpers.Contains(blacklist, cpf)
}
