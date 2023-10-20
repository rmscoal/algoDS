"""
https://leetcode.com/problems/rearrange-array-elements-by-sign/

You are given a 0-indexed integer array nums of even length consisting of an
equal number of positive and negative integers.

You should rearrange the elements of nums such that the modified array follows
the given conditions:

Every consecutive pair of integers have opposite signs.
For all integers with the same sign, the order in which they were present in
nums is preserved.
The rearranged array begins with a positive integer.

Return the modified array after rearranging the elements to satisfy the 
aforementioned conditions.

Constraints:
- 2 <= nums.length <= 2 * 105
- nums.length is even
- 1 <= |nums[i]| <= 105
- nums consists of equal number of positive and negative integers.
"""

from typing import List


class Solution:
    def rearrange_array(self, nums: List[int]) -> List[int]:
        positive_arr = []
        negative_arr = []

        for i in nums:
            if i > 0:
                positive_arr.append(i)
            else:
                negative_arr.append(i)

        i = 0
        while i < len(nums):
            if i % 2 == 0:
                nums[i] = positive_arr[i//2]
            else:
                nums[i] = negative_arr[(i-1)//2]
            i += 1

        return nums

    def rearrange_array_v2(self, nums: List[int]) -> List[int]:
        result = [0 for _ in range(len(nums))]
        left = 0
        right = 1

        for i in nums:
            if i > 0:
                result[left] = i
                left += 2
            else:
                result[right] = i
                right += 2

        return result


if __name__ == "__main__":
    print(Solution().rearrange_array([3, 1, -2, -5, 2, -4]))
    print(Solution().rearrange_array_v2([3, 1, -2, -5, 2, -4]))
