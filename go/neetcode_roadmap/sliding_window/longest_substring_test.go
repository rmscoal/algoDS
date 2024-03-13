package slidingwindow

import (
	"fmt"
	"testing"
)

// Longest Substring Without Repeating Characters
func lengthOfLongestSubstring_Codeium(s string) int {
	// Input: "abcabcbb"
	// Output: 3
	//
	// Input: "bbbbb"
	// Output: 1
	m := make(map[rune]int)
	start, maxLen := 0, 0

	for i, r := range s {
		if idx, ok := m[r]; ok {
			if idx >= start {
				start = idx + 1
			}
		}
		m[r] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	left, right := 0, 1
	hashmap := make(map[byte]bool)
	hashmap[s[left]] = true
	maxLen := 1

	for left <= right && right < len(s) {
		if exists, ok := hashmap[s[right]]; exists && ok {
			for {
				delete(hashmap, s[left])
				left++
				if exists, ok := hashmap[s[right]]; !exists && !ok {
					break
				}
			}
		}
		hashmap[s[right]] = true
		maxLen = max(maxLen, right-left+1)
		right++
	}

	return maxLen
}

func lengthOfLongestSubstring_Fast(s string) int {
	hm := make(map[rune]int, len(s))
	l, mx := 0, 0
	for r, v := range s {
		if value, ok := hm[v]; ok {
			l = max(value+1, l)
		}
		hm[v] = r
		mx = max(mx, r-l+1)
	}
	return mx
}

func TestLengthOfLongestSubstring_Codeium(t *testing.T) {
	testcases := []struct {
		input  string
		output int
	}{
		{input: "abcabcbb", output: 3},
		{input: "bbbbb", output: 1},
		{input: "pwwkew", output: 3},
		{input: "abcdefghijklmnopqrstuvwxyz", output: 26},
		// Add more test cases as needed
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Running testcase %d for lengthOfLongestSubstring_Codeium", idx+1), func(t *testing.T) {
			result := lengthOfLongestSubstring_Codeium(tc.input)
			if result != tc.output {
				t.Fatalf("Expected %d but got %d", tc.output, result)
			}
		})
		t.Run(fmt.Sprintf("Running testcase %d for lengthOfLongestSubstring", idx+1), func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			if result != tc.output {
				t.Fatalf("Expected %d but got %d", tc.output, result)
			}
		})
		t.Run(fmt.Sprintf("Running testcase %d for lengthOfLongestSubstring_Fast", idx+1), func(t *testing.T) {
			result := lengthOfLongestSubstring_Fast(tc.input)
			if result != tc.output {
				t.Fatalf("Expected %d but got %d", tc.output, result)
			}
		})
	}
}
