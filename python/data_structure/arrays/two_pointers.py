"""
Question 1: Given a string s, return true if it is a palindrome, false otherwise.

A string is a palindrome if it reads the same forward as backward. That means,
after reversing it, it is still the same string.

Example: "abcdcba" or "racecar".
"""


from typing import List


def two_pointer_q1_sol1(s: str) -> bool:
    left = 0
    right = len(s) - 1

    while left <= right:
        if s[left] != s[right]:
            return False
        left += 1
        right -= 1

    return True


"""
Question 2: Given a sorted array of unique integers and a target integer, return
true if there exists a pair of numbers that sum to target, false otherwise. This
problem is similar to Two Sum. (In Two Sum, the input is not sorted).

Example, given nums = [1, 2, 4, 6, 8, 9, 14, 15] and target = 13, return true
because 4 + 9 = 13.
"""


def two_pointer_q2_sol1(arr: List[int], target: int) -> bool:
    left = 0
    right = len(arr) - 1

    while left <= right:
        sum = arr[left] + arr[right]
        if sum == target:
            return True
        elif sum > target:
            right -= 1
        else:
            left += 1

    return False


"""
Question 3: Given two sorted integer arrays arr1 and arr2, return a new array
that combines both of them and is also sorted.
"""


def two_pointer_q3_sol1(arr1: List[int], arr2: List[int]):
    result = []
    left = 0
    right = 0

    while left < len(arr1) and right < len(arr2):
        if arr1[left] < arr2[right]:
            result.append(arr1[left])
            left += 1

        else:
            result.append(arr2[right])
            right += 1

    if left < len(arr1):
        result.extend(arr1[left:])

    if right < len(arr2):
        result.extend(arr2[right:])

    return result


"""
Question 4: Given two strings s and t, return true if s is a subsequence of t,
or false otherwise.

A subsequence of a string is a sequence of characters that can be obtained by
deleting some (or none) of the characters from the original string, while
maintaining the relative order of the remaining characters. For example, "ace"
is a subsequence of "abcde" while "aec" is not.
"""


def two_pointer_q4_sol1(sub: str, s: str) -> bool:
    left = 0
    right = 0

    while left < len(sub) and right < len(s):
        if sub[left] == s[right]:
            left += 1
        right += 1

    return left == len(sub)


"""
Question 5: Reverse a string.
"""


def two_pointer_q5_sol1(s: str) -> str:
    arr = [char for char in s]
    left = 0
    right = len(arr) - 1

    while left < right:
        arr[left], arr[right] = arr[right], arr[left]
        left, right = left + 1, right - 1

    return "".join(arr)


if __name__ == "__main__":
    ##################
    #   Question 1   #
    ##################
    print("Answer Question 1 Solution 1 Case 1",
          two_pointer_q1_sol1("racecar"))  # True
    print("Answer Question 1 Solution 1 Case 2",
          two_pointer_q1_sol1("abcdcba"))  # True
    print("Answer Question 1 Solution 1 Case 3",
          two_pointer_q1_sol1("abcdefg"))  # False

    print()

    ##################
    #   Question 2   #
    ##################
    print("Answer Question 2 Solution 1 Case 1",
          two_pointer_q2_sol1([1, 2, 4, 6, 8, 9, 14, 15], target=13))  # True
    print("Answer Question 2 Solution 1 Case 2",
          two_pointer_q2_sol1([1, 2, 4, 6, 8, 9, 14, 15], target=3))  # True
    print("Answer Question 2 Solution 1 Case 3",
          two_pointer_q2_sol1([1, 2, 4, 6, 8, 9, 14, 15], target=1))  # False

    print()

    ##################
    #   Question 3   #
    ##################
    print("Answer Question 3 Solution 1 Case 1",
          two_pointer_q3_sol1([1, 2, 3, 10, 11], [4, 5, 6, 7, 8, 9]))  # [1,2,3,4,5,6,7,8,9,10,11]

    print()

    ##################
    #   Question 4   #
    ##################
    print("Answer Question 4 Solution 1 Case 1",
          two_pointer_q4_sol1(sub="ace", s="abcde"))  # True
    print("Answer Question 4 Solution 1 Case 2",
          two_pointer_q4_sol1(sub="aec", s="abcde"))  # True

    print()

    ##################
    #   Question 5   #
    ##################
    print("Answer Question 5 Solution 1 Case 1",
          two_pointer_q5_sol1("abcde"))  # edcba

    print()
