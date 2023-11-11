package dp

import (
	"fmt"
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := map[string]struct {
		s, answer string
	}{
		"test_1": {
			s:      "babad",
			answer: "bab",
		},
		"test_2": {
			s:      "cbbd",
			answer: "bb",
		},
		"test_3": {
			s:      "abcdedcfg",
			answer: "cdedc",
		},
		"test_4": {
			s:      "a",
			answer: "a",
		},
		"test_5": {
			s:      "ba",
			answer: "b",
		},
		"test_6": {
			s:      "bac",
			answer: "b",
		},
		"test_7": {
			s:      "bacc",
			answer: "cc",
		},
		"test_8": {
			s:      "aacabdkacaa",
			answer: "aca",
		},
		"test_9": {
			s:      "abcddddddddddddddddddddef",
			answer: "dddddddddddddddddddd",
		},
		"test_10": {
			s:      "abbbc",
			answer: "bbb",
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Longest Palindromic Substring %s", name), func(t *testing.T) {
			if result := longestPalindromicSubstring(test.s); result != test.answer {
				t.Fatalf("My Solution - Expected %s is %s but got %s\n", name, test.answer, result)
			}
			if result := longestPalindrome(test.s); result != test.answer {
				t.Fatalf("Cool Solution - Expected %s is %s but got %s\n", name, test.answer, result)
			}
		})
	}
}
