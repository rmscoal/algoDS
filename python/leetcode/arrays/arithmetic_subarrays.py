"""
A sequence of numbers is called arithmetic if it consists of at least two
elements, and the difference between every two consecutive elements is the same.
More formally, a sequence s is arithmetic if and only if 
    s[i+1] - s[i] == s[1] - s[0]
for all valid i.

For example, these are arithmetic sequences:

1, 3, 5, 7, 9
7, 7, 7, 7
3, -1, -5, -9

The following sequence is not arithmetic:
1, 1, 2, 5, 7

You are given an array of n integers, nums, and two arrays of m integers each,
l and r, representing the m range queries, where the ith query is the range
[l[i], r[i]]. All the arrays are 0-indexed.

Return a list of boolean elements answer, where answer[i] is true if the subarray
nums[l[i]], nums[l[i]+1], ... , nums[r[i]] can be rearranged to form an arithmetic
sequence, and false otherwise.
"""

from typing import List, Tuple
import random


class Solution:
    def split_by_pivot(self, nums: List[int]) -> Tuple[List[int], List[int], List[int]]:
        pivot = nums.pop(random.randint(0, len(nums) - 1))
        left = []
        right = []

        for i in nums:
            if i < pivot:
                left.append(i)
            else:
                right.append(i)

        return left, [pivot], right

    def quicksort(self, nums: List[int]) -> List[int]:
        if len(nums) <= 1:
            return nums

        left, pivot, right = self.split_by_pivot(nums)

        left = self.quicksort(left)
        right = self.quicksort(right)

        return left + pivot + right

    def check_arithmetic_subarrays(self, nums: List[int], l: List[int], r: List[int]) -> List[bool]:
        result = [False for _ in range(len(l))]

        for i in range(len(l)):
            arr = nums[l[i]:r[i] + 1]
            arr = self.quicksort(arr)

            is_arithmetic = True
            for j in range(len(arr) - 1):
                if arr[j+1] - arr[j] != arr[1] - arr[0]:
                    is_arithmetic = False
                    break

            result[i] = is_arithmetic

        return result


if __name__ == "__main__":
    print(Solution().check_arithmetic_subarrays(
        [4, 6, 5, 9, 3, 7], [0, 0, 2], [2, 3, 5]))
