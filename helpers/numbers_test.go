package helpers

import "testing"

func TestOnlyNumbers(t *testing.T) {
	tables := []struct {
		input  string
		output string
	}{
		{"test", ""},
		{"123test", "123"},
		{"test123", "123"},
		{"123test123", "123123"},
	}

	for _, table := range tables {
		formated := OnlyNumbers(table.input)
		if formated != table.output {
			t.Errorf("Ouput invalid, given %v, expected %v", formated, table.output)
		}
	}
}
