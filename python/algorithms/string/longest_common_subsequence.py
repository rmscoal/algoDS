"""
Given two strings, S1 and S2, the task is to find the length of the Longest
Common Subsequence, i.e. longest subsequence present in both of the strings.
Examples:

Input: S1 = “AGGTAB”, S2 = “GXTXAYB”
Output: 4
Explanation: The longest subsequence which is present in both strings is “GTAB”.

Input: S1 = “BD”, S2 = “ABCD”
Output: 2
Explanation: The longest subsequence which is present in both strings is “BD”.
"""


def lcs(X: str, Y: str, i: int, j: int) -> int:
    """
    Using RECURSION

    Generate all the possible subsequence and find the longest among them that is
    present in both string.

    - Create a recursive function [say lcs()].
    - Check the relation between the First characters of the strings that are not yet processed.
        Depending on the relation call the next recursive function as mentioned above.
    - Return the length of the LCS received as the answer.
    """
    if i == 0 or j == 0:
        return 0
    elif X[i-1] == Y[j-1]:
        return 1 + lcs(X, Y, i-1, j-1)
    else:
        return max(lcs(X, Y, i, j-1), lcs(X, Y, i-1, j))


def lcs_dp(X, Y, m, n, dp):
    """
    Using MEMOIZATION

    - Create a recursive function. Also create a 2D array to store the result of
    a unique state. 
    - During the recursion call, if the same state is called more than once, then
    we can directly return the answer stored for that state instead of calculating
    again.
    """
    if m == 0 or n == 0:
        return 0

    if dp[m][n] != -1:
        return dp[m][n]

    if X[m-1] == Y[n-1]:
        dp[m][n] = 1 + lcs_dp(X, Y, m - 1, n - 1, dp)
        return dp[m][n]

    dp[m][n] = max(lcs_dp(X, Y, m-1, n, dp), lcs_dp(X, Y, m, n-1, dp))
    return dp[m][n]


if __name__ == "__main__":
    s1 = "AGGTAB"
    s2 = "GXTXAYB"
    dp = [[-1 for _ in range(len(s2)+1)] for _ in range(len(s1) + 1)]
    print("Length of LCS is", lcs_dp(s1, s2, len(s1), len(s2), dp))
