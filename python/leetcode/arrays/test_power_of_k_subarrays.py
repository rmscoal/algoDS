from typing import List
import unittest

class TestPowerOfKSubarrays(unittest.TestCase):
    def results_array(self, nums: List[int], k: int) -> List[int]:
        result = []
        queue = []

        left, right = 0, 0
        while left <= right and right < len(nums):
            # Checks whether the current window is consecutive and in ascending order.
            if k > 1 and right > 0 and nums[right] != nums[right-1]+1:
                queue.append(right-1)

            # Checks if the current window lenght is k. This could happen
            # during the inilization of the sliding window
            if right-left+1 != k:
                right += 1
                continue

            # Checks if the queue is empty
            if len(queue) == 0:
                result.append(nums[right])
            else:
                result.append(-1)

            if len(queue) > 0 and queue[0] == left:
                queue.pop(0)

            left += 1
            right += 1

        return result
    
    def test_results_array(self):
        self.assertEqual(self.results_array([1, 3, 4], 2), [-1, 4])
        self.assertEqual(self.results_array([1, 1, 1, 2], 2), [-1, -1, 2])
        self.assertEqual(self.results_array([1, 2, 3, 4, 3, 2, 5], 3), [3, 4, -1, -1, -1])

if __name__ == "__main__":
    unittest.main()