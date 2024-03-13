package slidingwindow

import "testing"

// You are given a string s and an integer k. You can choose any character
// of the string and change it to any other uppercase English character.
// You can perform this operation at most k times.
func characterReplacement(s string, k int) int {
	// Here we are going for sliding window approach

	// Stores the frequency of each character in the string
	frequency := make(map[byte]int)
	// Stores the most frequent character in the current window
	mostFrequent := 0
	// Result variable
	result := 0
	// Start of the window
	start := 0

	for end := 0; end < len(s); end++ {
		// Increment the frequency of the character at the end of the window
		frequency[s[end]]++
		// Recalculate the most frequent character in the current window
		mostFrequent = max(mostFrequent, frequency[s[end]])

		// If the difference between the most frequent character and the current
		// window length is greater than k, we need to shrink the window by moving
		// start of the window to the left.
		if end-start+1-mostFrequent > k {
			frequency[s[start]]--
			start++
		}

		result = max(result, end-start+1)
	}

	return result
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"ABABBC", 2, 5},  // Case 1
		{"AAAAA", 2, 5},   // Case 2
		{"ABBBBC", 10, 6}, // Case 3
		{"AABABBA", 2, 5}, // Case 4
	}

	for _, tt := range tests {
		got := characterReplacement(tt.s, tt.k)
		if got != tt.want {
			t.Errorf("characterReplacement(%s, %d) = %d; want %d", tt.s, tt.k, got, tt.want)
		}
	}
}
