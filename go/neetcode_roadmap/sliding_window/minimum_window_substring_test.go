package slidingwindow

import "testing"

// Given two strings s and t of lengths m and n respectively, return the minimum
// window substring of s such that every character in t (including duplicates) is
// included in the window. If there is no such substring, return the empty string "".
//
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
func minWindow(s string, t string) string {
	if len(t) > len(s) || len(t) == 0 {
		return ""
	}

	left := 0
	hashmapTarget := make(map[byte]int)
	hashmapCurrentWindow := make(map[byte]int)
	result := ""

	for i := 0; i < len(t); i++ {
		hashmapTarget[t[i]]++
		hashmapCurrentWindow[s[i]]++
	}

	right := len(t) - 1
	for right < len(s) && left <= right-len(t)+1 {
		// Check matches in current window
		match := true
		for runTarget, countTarget := range hashmapTarget {
			currentWindowCount := hashmapCurrentWindow[runTarget]
			if currentWindowCount < countTarget {
				match = false
				break
			}
		}
		if match {
			if result == "" {
				result = s[left : right+1]
			} else if right-left+1 < len(result) {
				result = s[left : right+1]
			}
			// If it matches we can minimize our window by moving left to the right
			hashmapCurrentWindow[s[left]]--
			left++
		}

		// If we can still move right
		if right < len(s)-1 && !match {
			right++
			hashmapCurrentWindow[s[right]]++
		} else if !match { // We instead move left only if it previously doesn't match
			hashmapCurrentWindow[s[left]]--
			left++
		}
	}

	return result
}

func minWindow_Faster(s string, t string) string {
	if len(t) > len(s) || t == "" {
		return ""
	}

	hashmapTarget := make(map[byte]int)
	hashmapCurrentWindow := make(map[byte]int)
	result := ""
	matches := 0

	for i := 0; i < len(t); i++ {
		hashmapTarget[t[i]]++
		hashmapCurrentWindow[s[i]]++
	}

	for runTarget, countTarget := range hashmapTarget {
		if hashmapCurrentWindow[runTarget] >= countTarget {
			matches++
		}
	}

	left := 0
	right := len(t) - 1
	for right < len(s) && left <= right-len(t)+1 {
		switch {
		case matches >= len(hashmapTarget):
			if right-left+1 < len(result) || result == "" {
				result = string(s[left : right+1])
			}
			fallthrough
		case right == len(s)-1:
			hashmapCurrentWindow[s[left]]--
			if hashmapCurrentWindow[s[left]] < hashmapTarget[s[left]] {
				matches--
			}
			left++
		default:
			right++
			hashmapCurrentWindow[s[right]]++
			if hashmapCurrentWindow[s[right]] == hashmapTarget[s[right]] {
				matches++
			}
		}
	}

	return result
}

func minWindow_Codeium(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	required := len(tMap)
	left, right := 0, 0
	formed := 0

	windowCounts := make(map[byte]int)
	ans := []int{-1, 0, 0}

	for right < len(s) {
		char := s[right]
		windowCounts[char]++

		if count, ok := tMap[char]; ok && windowCounts[char] == count {
			formed++
		}

		for formed == required {
			if ans[0] == -1 || right-left+1 < ans[0] {
				ans[0] = right - left + 1
				ans[1] = left
				ans[2] = right
			}

			windowCounts[s[left]]--
			if count, ok := tMap[s[left]]; ok && windowCounts[s[left]] < count {
				formed--
			}
			left++
		}

		right++
	}

	if ans[0] == -1 {
		return ""
	}

	return s[ans[1] : ans[2]+1]
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected string
	}{
		{"", "", ""},
		{"ADOBECODEBANC", "XYZ", ""},
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"aa", "aa", "aa"},
		{"aacab", "aa", "aa"},
		{"aacab", "aadcba", ""},
		{"aacabd", "aadcba", "aacabd"},
		{"abc", "b", "b"},
	}

	for _, test := range tests {
		result := minWindow(test.s, test.t)
		if result != test.expected {
			t.Errorf("minWindow - For %s and %s, expected %s, but got %s", test.s, test.t, test.expected, result)
		}

		result = minWindow_Faster(test.s, test.t)
		if result != test.expected {
			t.Errorf("minWindow_Cleaner - For %s and %s, expected %s, but got %s", test.s, test.t, test.expected, result)
		}

		result = minWindow_Codeium(test.s, test.t)
		if result != test.expected {
			t.Errorf("minWindow_Codeium - For %s and %s, expected %s, but got %s", test.s, test.t, test.expected, result)
		}
	}
}
