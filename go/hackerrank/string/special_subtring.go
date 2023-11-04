package string

func allSame(s string, chr rune) bool {
	for _, v := range s {
		if v != chr {
			return false
		}
	}

	return true
}

func suffixAndPrefixSame(s string) bool {
	if len(s)%2 == 0 {
		return false
	}

	mid := (len(s) - 1) / 2

	chr := rune(s[0])

	return allSame(string(s[:mid]), chr) && allSame(string(s[mid+1:]), chr)
}

// O(N^2)
func substrCount(n int32, s string) int64 {
	var count int64 = 0

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if allSame(string(s[i:j+1]), rune(s[i])) {
				count++
			} else if suffixAndPrefixSame(string(s[i : j+1])) {
				count++
			}
		}
	}

	return count
}

// https://www.hackerrank.com/challenges/special-palindrome-again/problem
func specialSubstrCount(n int32, s string) int64 {
	// Here we are going to find all same first
	// Then find the palindrome
	var count int64 = 0
	index := 0

	for index < len(s) {
		chr := s[index]
		similarCount := 1
		index++
		for index < len(s) && s[index] == chr {
			similarCount++
			index++
		}
		count += int64(similarCount * (similarCount + 1) / 2)
	}

	for i := 1; i < len(s)-1; i++ {
		if s[i] == s[i-1] || s[i] == s[i+1] {
			continue
		}

		left := i - 1
		right := i + 1
		chr := s[i-1]

		for left >= 0 && right < len(s) && s[left] == s[right] && s[left] == chr {
			count++
			left--
			right++
		}
	}

	return count
}
