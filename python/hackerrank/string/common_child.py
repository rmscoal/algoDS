def common_child_slow(s1: str, s2: str) -> int:
    """
    https://www.hackerrank.com/challenges/common-child/
    """
    if len(s1) == 0 or len(s2) == 0:
        return 0
    elif s1[len(s1) - 1] == s2[len(s2) - 1]:
        return 1 + common_child_slow(s1[:len(s1)-1], s2[:len(s2)-1])
    else:
        return max(common_child_slow(s1[:len(s1)-1], s2), common_child_slow(s1, s2[:len(s2)-1]))


def common_child_memoization(s1: str, s2: str) -> int:
    """
    https://www.hackerrank.com/challenges/common-child/

    Bottom up Approach
    """
    dp = [[0 for _ in range(len(s2) + 1)] for _ in range(len(s1) + 1)]

    for i in range(1, len(s1) + 1):
        for j in range(1, len(s2) + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = dp[i-1][j-1] + 1
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])

    return dp[len(s1)][len(s2)]


if __name__ == "__main__":
    print("Common Child Result", common_child_memoization("AGGTAB", "GXTXAYB"))
    print("Common Child Result", common_child_memoization(
        "ABCD", "ABDC"))
    print("Common Child Result", common_child_memoization(
        "WEWOUCUIDGCGTRMEZEPXZFEJWISRSBBSYXAYDFEJJDLEBVHHKS", "FDAGCXGKCTKWNECHMRXZWMLRYUCOCZHJRRJBOAJOQJZZVUYXIC"))
