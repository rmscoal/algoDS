package string

import (
	"fmt"
	"testing"
)

func TestMakingAnagram(t *testing.T) {
	tests := map[string]struct {
		a, b   string
		answer int32
	}{
		"test_1": {
			a:      "cde",
			b:      "dcf",
			answer: 2,
		},
		"test_2": {
			a:      "cde",
			b:      "abc",
			answer: 4,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Making Anagram %s", name), func(t *testing.T) {
			result := makeAnagrams(test.a, test.b)
			if result != test.answer {
				t.Fatalf("Expectation is %d but got %d\n", test.answer, result)
			}
		})
	}
}
