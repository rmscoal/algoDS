package string

/*
Given two strings, S1 and S2, the task is to find the length of the Longest
Common Subsequence, i.e. longest subsequence present in both of the strings.
Examples:

Input: S1 = “AGGTAB”, S2 = “GXTXAYB”
Output: 4
Explanation: The longest subsequence which is present in both strings is “GTAB”.

Input: S1 = “BD”, S2 = “ABCD”
Output: 2
Explanation: The longest subsequence which is present in both strings is “BD”.
*/

/*
Using MEMOIZATION BOTTOM-UP APPROACH

Generate all the possible subsequence and find the longest among them that is
present in both string.

- Create a recursive function [say lcs()].
- Check the relation between the First characters of the strings that are not yet processed.
	Depending on the relation call the next recursive function as mentioned above.
- Return the length of the LCS received as the answer.
*/

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(s1, s2 string) int {
	memo := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		memo[i] = make([]int, len(s2)+1)
	}

	// A B C D
	// A B D C

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				memo[i][j] = 1 + memo[i-1][j-1]
			} else {
				// i = 4, j = 1
				memo[i][j] = findMax(memo[i][j-1], memo[i-1][j])
			}
		}
	}

	return memo[len(s1)][len(s2)]
}
