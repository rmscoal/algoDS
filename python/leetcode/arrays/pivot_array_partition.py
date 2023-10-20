"""
You are given a 0-indexed integer array nums and an integer pivot. Rearrange
nums such that the following conditions are satisfied:

1. Every element less than pivot appears before every element greater than pivot.
2. Every element equal to pivot appears in between the elements less than and
greater than pivot.
3. The relative order of the elements less than pivot and the elements greater
than pivot is maintained.
More formally, consider every pi, pj where pi is the new position of the ith
element and pj is the new position of the jth element. For elements less than
pivot, if i < j and nums[i] < pivot and nums[j] < pivot, then pi < pj. Similarly
for elements greater than pivot, if i < j and nums[i] > pivot and nums[j] > pivot,
then pi < pj.

Return nums after the rearrangement.
"""

from typing import List


class Solution:
    def pivot_array(self, nums: List[int], pivot: int) -> List[int]:
        """
        O(n): Linear
        """
        pivots: List[int] = []
        lessers: List[int] = []
        greaters: List[int] = []

        for i in nums:
            if i == pivot:
                pivots.append(i)
            elif i < pivot:
                lessers.append(i)
            else:
                greaters.append(i)

        return lessers + pivots + greaters

    def should_shift(self, left_value, right_value, pivot) -> bool:
        if left_value > pivot and right_value <= pivot:
            return True
        elif left_value == pivot and right_value < left_value:
            return True
        else:
            return False

    def pivot_array_v2(self, nums: List[int], pivot: int) -> List[int]:
        """
        Time Limit Exceeded
        """
        if len(nums) <= 1:
            return nums

        shifted = True
        while shifted:
            left = 0
            right = 1
            shifted = False
            while right < len(nums):
                shouldShift = self.should_shift(nums[left], nums[right], pivot)
                if shouldShift:
                    nums[left], nums[right] = nums[right], nums[left]
                    shifted = True
                left += 1
                right += 1

        return nums

    def pivot_array_v3(self, nums: List[int], pivot: int) -> List[int]:
        """
        O(n) Time: Two Pointer
        O(1) Space
        """
        res = [0 for _ in range(len(nums))]
        left = 0
        right = len(nums) - 1

        for i in range(len(nums)):
            if nums[i] < pivot:
                res[left] = nums[i]
                left += 1

            if nums[len(nums) - i - 1] > pivot:
                res[right] = nums[len(nums) - i - 1]
                right -= 1

        while left <= right:
            res[left] = pivot
            left += 1

        return res


if __name__ == "__main__":
    print(Solution().pivot_array([9, 12, 5, 10, 14, 3, 10], 10))
    print(Solution().pivot_array_v2([9, 12, 5, 10, 14, 3, 10], 10))
    print(Solution().pivot_array_v3([9, 12, 5, 10, 14, 3, 10], 10))
