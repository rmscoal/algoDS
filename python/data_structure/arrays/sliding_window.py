from typing import List

'''
Sliding window is another common approach to solving problems related to arrays.
A sliding window is actually implemented using two pointers! Before we start, we
need to talk about the concept of a subarray.
The idea behind a sliding window is to consider only valid subarrays. Recall that
a subarray can be defined by a left bound (the index of the first element) and a
right bound (the index of the last element). In sliding window, we maintain two
variables left and right, which at any given time represent the current subarray
under consideration.

Amortized Analysis of O(n)
'''

'''
Question 1: Given an array of positive integers nums and an integer k, find the
length of the longest subarray whose sum is less than or equal to k.

Example:
    nums = [3, 1, 2, 7, 4, 2, 1, 1, 5] and k = 8 answer = 4
'''


def sliding_window_q_1(arr: List[int], k: int) -> int:
    # Indicates the left index
    left = 0
    # Keeps track of the current sum of the subarray (window)
    cursor = 0
    # The max length of array
    max = 0

    for right in range(len(arr)):
        cursor += arr[right]
        while cursor > k:
            cursor -= arr[left]
            left += 1

        if max <= right - left + 1:
            max = right - left + 1

    return max


'''
Question 2: You are given a binary string s (a string containing only "0" and "1").
You may choose up to one "0" and flip it to a "1". What is the length of the longest
substring achievable that contains only "1"?

Example: 
    For example, given s = "1101100111", the answer is 5. If you perform the flip at
    index 2, the string becomes 1111100111.

[1, 1, 0, 1, 1, 0, 0, 1, 1, 1]
'''


def sliding_window_q_2_solution_1(s: str) -> int:
    '''
    Here we are literally switching the bits
    '''
    left = 0
    flipped_index = 0
    is_flipped = False
    max = 0
    arr = [char for char in s]

    for right in range(len(arr)):
        while arr[right] != "1":
            if is_flipped:
                arr[flipped_index], left = "0", flipped_index + 1
                flipped_index, is_flipped = 0, False
            else:
                arr[right], is_flipped, flipped_index = "1", True, right

        if max <= right - left + 1:
            max = right - left + 1

    return max


def sliding_window_q_2_solution_2(s: str) -> int:
    '''
    Here we see the problem as, what is the longest
    substring that contains at most one "0"
    '''
    left = 0
    cursor = 0
    max = 0
    arr = [char for char in s]

    for right in range(len(arr)):
        if arr[right] == "0":
            cursor += 1

        while cursor > 1:
            if arr[left] == "0":
                cursor -= 1
            left += 1

        if max <= right - left + 1:
            max = right - left + 1

    return max


'''
Question 3: Given an array of positive integers `nums` and an integer k, return the 
number of subarrays where the product of all the elements in the subarray is strictly
less than k.

Example: Given the input nums = [10, 5, 2, 6], k = 100, the answer is 8. The subarrays
with products less than k are:
    [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
'''


def not_sliding_window_q_3_solution_1(arr: List[int], k: int) -> int:
    total_geq_k = 0
    n = len(arr)

    total_combinations = int(len(arr) * (len(arr) + 1) / 2)

    for left in range(0, n, 1):
        cursor = arr[left]
        for right in range(left, n, 1):
            if right != left:
                cursor *= arr[right]

            if cursor >= k:
                total_geq_k += 1

    return total_combinations - total_geq_k


def sliding_window_q_3_solution_2(arr: List[int], k: int) -> int:
    # We can never have a valid window since the elements of
    # arr is positive integers.
    if k <= 1:
        return 0

    left = total = 0
    cursor = 1
    n = len(arr)

    for right in range(n):
        cursor *= arr[right]

        while cursor >= k:
            # Even if left surpass right that is really fine
            # because the cursor will be reset to 1 due to
            # division by itself
            cursor //= arr[left]
            left += 1

        # Count the total of subarray between arr[left:right+1]
        total += right - left + 1

    return total


'''
Question 4: Given an integer array nums and an integer k, find the sum of the subarray
with the largest sum whose length is k.
'''


def sliding_window_q_4_solution_1(arr: List[int], k: int) -> int:
    left = max = cursor = 0
    n = len(arr)

    for right in range(0, k):
        cursor += arr[right]
        if max < cursor:
            max = cursor

    for right in range(k, n):
        cursor += arr[right] - arr[left]
        left += 1
        if max < cursor:
            max = cursor

    return max


def sliding_window_q_4_solution_2(arr: List[int], k: int) -> int:
    ans = cursor = 0
    n = len(arr)

    for right in range(0, k):
        cursor += arr[right]

    ans = cursor

    for right in range(k, n):
        cursor += arr[right] - arr[right - k]
        ans = max(ans, cursor)

    return ans


if __name__ == "__main__":
    print("Answer Question 1", sliding_window_q_1(
        [3, 1, 2, 7, 4, 2, 1, 1, 5], 8))  # 4

    print()

    ##################
    #   Question 2   #
    ##################
    # Using solution 1
    print("Answer Question 2 Case 1 with Solution 1",
          sliding_window_q_2_solution_1("1101100111"))  # 5
    print("Answer Question 2 Case 2 with Solution 1",
          sliding_window_q_2_solution_1("01111101101"))  # 8
    print("Answer Question 2 Case 3 with Solution 1",
          sliding_window_q_2_solution_1("011111001101"))  # 6

    print()

    # Using solution 2
    print("Answer Question 2 Case 1 with Solution 2",
          sliding_window_q_2_solution_2("1101100111"))  # 5
    print("Answer Question 2 Case 2 with Solution 2",
          sliding_window_q_2_solution_2("01111101101"))  # 8
    print("Answer Question 2 Case 3 with Solution 2",
          sliding_window_q_2_solution_2("011111001101"))  # 6

    print()

    ##################
    #   Question 3   #
    ##################
    # Using solution 1
    print("Answer Question 3 Case 1 with Solution 1",
          not_sliding_window_q_3_solution_1(arr=[10, 5, 2, 6], k=100))  # 8
    print("Answer Question 3 Case 2 with Solution 1",
          not_sliding_window_q_3_solution_1(arr=[1, 1], k=100))  # 3
    print("Answer Question 3 Case 3 with Solution 1",
          not_sliding_window_q_3_solution_1(arr=[1, 3, 100, 500, 7], k=100))  # 4

    print()

    # Using solution 2
    print("Answer Question 3 Case 1 with Solution 2",
          sliding_window_q_3_solution_2(arr=[10, 5, 2, 6], k=100))  # 8
    print("Answer Question 3 Case 2 with Solution 2",
          sliding_window_q_3_solution_2(arr=[1, 1], k=100))  # 3
    print("Answer Question 3 Case 3 with Solution 2",
          sliding_window_q_3_solution_2(arr=[1, 3, 100, 500, 7], k=100))  # 4

    print()

    ##################
    #   Question 4   #
    ##################
    # Using solution 1
    print("Answer Question 4 Case 1 with Solution 1",
          sliding_window_q_4_solution_1(arr=[10, 5, 2, 6], k=2))  # 15

    print()

    # Using solution 2
    print("Answer Question 4 Case 1 with Solution 2",
          sliding_window_q_4_solution_2(arr=[10, 5, 2, 6], k=2))  # 15
