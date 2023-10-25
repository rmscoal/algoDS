package string

import (
	"fmt"
	"testing"
)

func TestReverseWord(t *testing.T) {
	tests := map[string]struct {
		s        string
		expected string
	}{
		"test_1": {
			s:        "the sky is blue",
			expected: "blue is sky the",
		},
		"test_2": {
			s:        "  hello world  ",
			expected: "world hello",
		},
		"test_3": {
			s:        "a good   example",
			expected: "example good a",
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Reverse Word in String %s", name), func(t *testing.T) {
			result := reverseWords(test.s)
			if result != test.expected {
				t.Fatalf("Expected %s but got %s\n", test.expected, result)
			}
		})
	}
}
