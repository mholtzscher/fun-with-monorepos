package math

import "testing"

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedOutput uint64
	}{
		{"Factorial of 0", 0, 1},
		{"Factorial of 1", 1, 1},
		{"Factorial of 2", 2, 2},
		{"Factorial of 5", 5, 120},
		{"Factorial of 10", 10, 3628800},
		{"Factorial of -1", -1, 0},
		{"Factorial of -5", -5, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := Factorial(tc.input)
			if actualOutput != tc.expectedOutput {
				t.Errorf("Factorial(%d) = %d; expected %d", tc.input, actualOutput, tc.expectedOutput)
			}
		})
	}
}
