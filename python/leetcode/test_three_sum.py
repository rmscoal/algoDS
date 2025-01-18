import unittest
from typing import List


def three_sum(nums: List[int]) -> List[List[int]]:
    nums.sort()
    result = []
    for i in range(len(nums) - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue

        target, left, right = -nums[i], i + 1, len(nums) - 1
        while left < right:
            total = nums[left] + nums[right]
            if total == target:
                result.append([nums[i], nums[left], nums[right]])
                left += 1
                right -= 1
                while left < right and nums[left] == nums[left - 1]:
                    left += 1
                while left < right and nums[right] == nums[right + 1]:
                    right -= 1
            elif total < target:
                left += 1
            else:
                right -= 1
    return result

class ThreeSum(unittest.TestCase):
    def test_three_sum(self):
        self.assertListEqual([[-1,-1,2],[-1,0,1]], three_sum([-1,0,1,2,-1,-4]))
        self.assertListEqual([], three_sum([0,0,1]))
        self.assertListEqual([], three_sum([0, 1, 1]))
        self.assertListEqual([], three_sum([1, 1, 1]))


if __name__ == '__main__':
    unittest.main()
