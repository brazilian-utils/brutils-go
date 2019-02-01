package cnpj

import (
	"strconv"
	"strings"

	"github.com/brazilian-utils/brutils-go/helpers"
)

var blacklist = []string{
	"00000000000000",
	"11111111111111",
	"22222222222222",
	"33333333333333",
	"44444444444444",
	"55555555555555",
	"66666666666666",
	"77777777777777",
	"88888888888888",
	"99999999999999",
}

// verifierIndexes is the position of verification digits
var verifierIndexes = []int{12, 13}

// CNPJSize represents a valid size for CNPJ
// ignoring special characters
const CNPJSize = 14

// IsValid validates if a given CNPJ is valid
func IsValid(cnpj string) bool {
	digits := helpers.OnlyNumbers(cnpj)

	return hasValidLength(digits) && !isBlacklisted(cnpj) && isValidChecksum(digits)
}

// isValidChecksum perform checksum validation
func isValidChecksum(cnpj string) bool {
	validity := true

	for _, verifier := range verifierIndexes {
		mod := computeMod(strings.Split(cnpj[:verifier], ""), verifier)

		valid, _ := strconv.Atoi(string(cnpj[verifier]))
		validity = validity && (valid == mod)
	}

	return validity
}

// Compute the mod for the current slice of strings
func computeMod(digits []string, verifier int) (res int) {
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	if verifier == verifierIndexes[len(verifierIndexes)-1] {
		weights = append([]int{6}, weights...)
	}

	var mod int
	for index, digitStr := range digits {
		digit, _ := strconv.Atoi(digitStr)
		mod += digit * weights[index]
	}

	if mod = mod % 11; mod >= 2 {
		res = 11 - mod
	}

	return
}

// hasValidLength verify if the given CNPJ
// has valid length
func hasValidLength(cnpj string) bool {
	return len(cnpj) == CNPJSize
}

func isBlacklisted(cnpj string) bool {
	return helpers.Contains(blacklist, cnpj)
}
