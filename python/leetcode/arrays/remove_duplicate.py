"""
Given an integer array nums sorted in non-decreasing order, remove the duplicates
in-place such that each unique element appears only once. The relative order of
the elements should be kept the same. Then return the number of unique elements
in nums.
"""

from typing import List, Union


class Solution:
    def remove_duplicates(self, nums: List[Union[int, str]]) -> int:
        left = 0
        right = 1
        total_duplicate = 0

        while (left < len(nums) and right < len(nums)):
            if nums[left] == "_" or nums[right] == "_":
                break

            if nums[left] == nums[right]:
                nums.pop(right)
                nums.append("_")
                total_duplicate += 1
            else:
                left += 1
                right += 1

        return len(nums) - total_duplicate


if __name__ == "__main__":
    print(Solution().remove_duplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]))
    print(Solution().remove_duplicates([1, 1, 2]))
