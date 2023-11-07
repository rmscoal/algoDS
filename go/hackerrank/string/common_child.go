package string

func findMax(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

// https://www.hackerrank.com/challenges/common-child/problem
func commonChild(s1, s2 string) int32 {
	// We are going to use bottom up approach here of Longest Common Subsequence

	// Create our memoization dynamic programming
	n := len(s1) + 1
	memo := make([][]int32, n)
	for i := range memo {
		memo[i] = make([]int32, n)
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if s1[i-1] == s2[j-1] {
				memo[i][j] = 1 + memo[i-1][j-1]
			} else {
				memo[i][j] = findMax(memo[i-1][j], memo[i][j-1])
			}
		}
	}

	return memo[n-1][n-1]
}
