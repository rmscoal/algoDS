package leetcode

import "strconv"

func (P *PalindromeNumber) Solver(x int) bool {
	var s string = strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
