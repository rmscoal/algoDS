import unittest
import heapq
from typing import List


class Solution:
    def max_sliding_window_heapify(self, nums: List[int], k: int) -> List[int]:
        # Timeout exceeded
        results = []
        left = 0
        right = k-1

        while right < len(nums):
            cp = nums[left:right+1]
            heapq._heapify_max(cp)
            results.append(cp[0])
            left += 1
            right += 1

        return results


class TestMaxSlidingWindow(unittest.TestCase):
    def test_case_k_1(self):
        obj = Solution()
        nums = [1, 3, -1, -3, 5, 3, 6, 7]
        k = 1
        self.assertEqual(obj.max_sliding_window(
            nums, k), [1, 3, -1, -3, 5, 3, 6, 7])

    def test_case_k_2(self):
        obj = Solution()
        nums = [1, 3, -1, -3, 5, 3, 6, 7]
        k = 2
        self.assertEqual(obj.max_sliding_window(
            nums, k), [3, 3, -1, 5, 5, 6, 7])

    def test_case_k_3(self):
        obj = Solution()
        nums = [1, 3, -1, -3, 5, 3, 6, 7]
        k = 3
        self.assertEqual(obj.max_sliding_window(
            nums, k), [3, 3, 5, 5, 6, 7])


if __name__ == '__main__':
    unittest.main()
