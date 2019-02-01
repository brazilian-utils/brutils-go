package cep

import "testing"

var testCases = []struct {
	cep      string
	expected bool
	message  string
}{
	{
		cep:      "00000000",
		expected: true,
		message:  "valid CEP without special characters",
	},
	{
		cep:      "00000-000",
		expected: true,
		message:  "valid CEP with special characters",
	},
	{
		cep:      "f000x00e",
		expected: false,
		message:  "invalid CEP without special characters",
	},
	{
		cep:      "f000x-00e",
		expected: false,
		message:  "invalid CEP with special characters",
	},
}

func TestValidCEP(t *testing.T) {
	for _, tc := range testCases {
		if result := IsValid(tc.cep); result != tc.expected {
			t.Errorf("Failing for %s \t Expected: %t | Received: %t", tc.message, tc.expected, result)
		}
	}
}
