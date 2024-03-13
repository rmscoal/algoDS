package dp

func longestPalindromicSubstring(s string) string {
	if len(s) <= 1 {
		return s
	}

	result := string(s[0])

	// Make our memoization matrix
	palindrom := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		palindrom[i] = make([]bool, len(s))
		palindrom[i][i] = true
	}

	for i := 0; i < len(s)-1; i++ {
		left, right := i, i+1
		ppRight := false

		for left >= 0 && right <= len(s)-1 {
			if s[left] == s[right] && palindrom[left+1][right-1] {
				palindrom[left][right] = true
			} else if s[left] == s[right] && palindrom[left][right-1] && right-left == 1 {
				palindrom[left][right] = true
			}

			if palindrom[left][right] {
				if right-left+1 > len(result) {
					result = string(s[left : right+1])
				}
			}

			if ppRight {
				right++
			} else {
				left--
			}
			ppRight = !ppRight
		}
	}

	return result
}

/**
Cool Solution
*/

func expand(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return s[i+1 : j]
}

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if str1 := expand(s, i, i); len(str1) > len(res) {
			res = str1
		}

		if str2 := expand(s, i, i+1); len(str2) > len(res) {
			res = str2
		}
	}

	return res
}
