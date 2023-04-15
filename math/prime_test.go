package math

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"One", 1, false},
		{"Two", 2, true},
		{"Three", 3, true},
		{"Four", 4, false},
		{"Five", 5, true},
		{"Six", 6, false},
		{"Seven", 7, true},
		{"Eight", 8, false},
		{"Nine", 9, false},
		{"Ten", 10, false},
		{"Eleven", 11, true},
		{"Twelve", 12, false},
		{"Thirteen", 13, true},
		{"Fourteen", 14, false},
		{"Fifteen", 15, false},
		{"Sixteen", 16, false},
		{"Seventeen", 17, true},
		{"Eighteen", 18, false},
		{"Nineteen", 19, true},
		{"Twenty", 20, false},
		{"Twenty-one", 21, false},
		{"Twenty-two", 22, false},
		{"Twenty-three", 23, true},
		{"Twenty-four", 24, false},
		{"Twenty-five", 25, false},
		{"Twenty-six", 26, false},
		{"Twenty-seven", 27, false},
		{"Twenty-eight", 28, false},
		{"Twenty-nine", 29, true},
		{"Thirty", 30, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsPrime(tt.input) != tt.expected {
				t.Errorf("IsPrime(%d) = %v; want %v", tt.input, !tt.expected, tt.expected)
			}
		})
	}
}
