from typing import List, Dict
import unittest


class TestLongestConsecutiveSequence(unittest.TestCase):
    def longest_consecutive_sequence_sort(self, nums: List[int]) -> int:
        """
        Just sort le
        """
        # Sort the nums
        nums.sort()

        N = len(nums)
        result = 1 if N >= 1 else 0
        count = 1
        for i in range(1, N):
            if nums[i] == nums[i - 1] + 1:
                count += 1
            elif nums[i] == nums[i - 1]:
                continue
            else:
                count = 1

            result = max(result, count)

        return result

    def longest_consecutive_sequence(self, nums: List[int]) -> int:
        N = len(nums)
        if N <= 1:
            return N

        hash: Dict[int, bool] = {}
        for i in range(N):
            hash[nums[i]] = True

        result = 1
        for i in hash:
            # We only want to start at the beginning of the sequence
            if i - 1 in hash:
                continue
            else:
                count = 1
                j = i + 1
                while True:
                    if j in hash:
                        count += 1
                    else:
                        result = max(result, count)
                        break
                    j += 1

        return result

    def test_longest_consecutive_sequence(self):
        self.assertEqual(4, self.longest_consecutive_sequence(
            [100, 4, 200, 1, 3, 2]))
        self.assertEqual(9, self.longest_consecutive_sequence(
            [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]))
        self.assertEqual(1, self.longest_consecutive_sequence(
            [0]))
        self.assertEqual(0, self.longest_consecutive_sequence(
            []))
        self.assertEqual(3, self.longest_consecutive_sequence(
            [1, 2, 0, 1]))

        self.assertEqual(4, self.longest_consecutive_sequence_sort(
            [100, 4, 200, 1, 3, 2]))
        self.assertEqual(9, self.longest_consecutive_sequence_sort(
            [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]))
        self.assertEqual(1, self.longest_consecutive_sequence_sort(
            [0]))
        self.assertEqual(0, self.longest_consecutive_sequence_sort(
            []))
        self.assertEqual(3, self.longest_consecutive_sequence_sort(
            [1, 2, 0, 1]))


if __name__ == "__main__":
    unittest.main()
