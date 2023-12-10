package arrays

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)

	fmt.Printf("Lower string %s\n", s)

	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
			left++
			continue
		}

		if !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
			right--
			continue
		}

		if rune(s[right]) != rune(s[left]) {
			return false
		}

		left++
		right--
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	if result := isPalindrome("A man, a plan, a canal: Panama"); result != true {
		t.Fatalf("Expected to be true but got %t", result)
	}
	if result := isPalindrome("race a car"); result != false {
		t.Fatalf("Expected to be false but got %t", result)
	}
	if result := isPalindrome(" "); result != true {
		t.Fatalf("Expected to be true but got %t", result)
	}
	if result := isPalindrome("me;me k.e ;m; e m"); result != true {
		t.Fatalf("Expected to be true but got %t", result)
	}
	if result := isPalindrome("0P"); result != false {
		t.Fatalf("Expected to be false but got %t", result)
	}
}
