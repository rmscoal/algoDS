from typing import List
import unittest


class Solution:
    def naive(self, nums: List[int]) -> None:
        zero_counter = 0
        one_counter = 0
        two_counter = 0

        for i in range(len(nums)):
            if nums[i] == 0:
                zero_counter += 1
            elif nums[i] == 1:
                one_counter += 1
            else:
                two_counter += 1

        for i in range(len(nums)):
            if zero_counter > 0:
                nums[i] = 0
                zero_counter -= 1
            elif one_counter > 0:
                nums[i] = 1
                one_counter -= 1
            else:
                nums[i] = 2
                two_counter -= 1

    def swap(self, nums: List[int], i: int, j: int) -> None:
        '''
        Swap i and j in nums
        '''
        nums[i], nums[j] = nums[j], nums[i]

    def three_way_partition(self, nums: List[int]) -> None:
        start = current = 0
        pivot = 1
        end = len(nums) - 1

        while current <= end:
            if nums[current] < pivot:       # We found 0
                self.swap(nums, start, current)
                current += 1
                start += 1
            elif nums[current] > pivot:     # We found 2
                self.swap(nums, current, end)
                end -= 1
            else:                           # We found 1
                current += 1


class TestCase(unittest.TestCase):
    def test_counter(self):
        num = [0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0]
        Solution().naive(num)
        self.assertEqual([0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2], num)

    def test_three_way(self):
        num = [0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0]
        Solution().three_way_partition(num)
        self.assertEqual([0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2], num)


if __name__ == "__main__":
    unittest.main()
