package string

import (
	"fmt"
	"testing"
)

func TestAlternatingCharacters(t *testing.T) {
	tests := map[string]struct {
		s      string
		answer int32
	}{
		"test_1": {
			s:      "AAAA",
			answer: 3,
		},
		"test_2": {
			s:      "BBBBB",
			answer: 4,
		},
		"test_3": {
			s:      "ABABABAB",
			answer: 0,
		},
		"test_4": {
			s:      "AAABBB",
			answer: 4,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Alternating Characters %s", name), func(t *testing.T) {
			if result := alternatingCharacters(test.s); result != test.answer {
				t.Fatalf("Expectation is %d but got %d\n", test.answer, result)
			}
		})
	}
}
