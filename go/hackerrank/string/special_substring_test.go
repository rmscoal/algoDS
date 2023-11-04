package string

import (
	"fmt"
	"testing"
)

func TestAllSame(t *testing.T) {
	tests := map[string]struct {
		s      string
		answer bool
	}{
		"test_1": {
			s:      "ab",
			answer: false,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test All Same %s", name), func(t *testing.T) {
			if result := allSame(test.s, rune(test.s[0])); result != test.answer {
				t.Fatalf("Expectation is %t but got %t\n", test.answer, result)
			}
		})
	}
}

func TestSuffixAndPrefixSame(t *testing.T) {
	tests := map[string]struct {
		s      string
		answer bool
	}{
		"test_1": {
			s:      "ab",
			answer: false,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test All Same %s", name), func(t *testing.T) {
			if result := suffixAndPrefixSame(test.s); result != test.answer {
				t.Fatalf("Expectation is %t but got %t\n", test.answer, result)
			}
		})
	}
}

func TestSubstring(t *testing.T) {
	tests := map[string]struct {
		n      int32
		s      string
		answer int64
	}{
		"test_1": {
			n:      7,
			s:      "abcbaba",
			answer: 10,
		},
		"test_2": {
			n:      4,
			s:      "aaaa",
			answer: 10,
		},
		"test_3": {
			n:      5,
			s:      "asasd",
			answer: 7,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Special Substring %s", name), func(t *testing.T) {
			if result := substrCount(test.n, test.s); result != test.answer {
				t.Fatalf("Expectation is %d but got %d\n", test.answer, result)
			}
		})
	}
}

func TestSpecialSubstring(t *testing.T) {
	tests := map[string]struct {
		n      int32
		s      string
		answer int64
	}{
		"test_1": {
			n:      7,
			s:      "abcbaba",
			answer: 10,
		},
		"test_2": {
			n:      4,
			s:      "aaaa",
			answer: 10,
		},
		"test_3": {
			n:      5,
			s:      "asasd",
			answer: 7,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Special Substring %s", name), func(t *testing.T) {
			if result := specialSubstrCount(test.n, test.s); result != test.answer {
				t.Fatalf("Expectation for %s is %d but got %d\n", name, test.answer, result)
			}
		})
	}
}
