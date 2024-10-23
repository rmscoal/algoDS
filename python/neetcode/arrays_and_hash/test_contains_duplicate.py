from typing import List
import unittest


class TestContainsDuplicate(unittest.TestCase):
    def contains_duplicate(self, nums: List[int]):
        mapper = {}

        for i in nums:
            if i in mapper:
                return True
            else:
                mapper[i] = 1

        return False

    def test_contains_duplicate(self):
        self.assertTrue(self.contains_duplicate([1, 2, 3, 1]))

    def contains_duplicate_ii(self, nums: List[int], k: int):
        i, j = 0, 1
        mapper = { nums[i]: 1 }

        while j < len(nums):
            if nums[j] in mapper:
                return True
            else:
                mapper[nums[j]] = 1
                j += 1
                if abs(j - i) > k:
                    mapper.pop(nums[i])
                    i += 1
        
        return False

    def test_contains_duplicate_ii(self):
        self.assertTrue(self.contains_duplicate_ii([1,2,3,1], 3))
        self.assertTrue(self.contains_duplicate_ii([1,0,1,1], 1))
        self.assertFalse(self.contains_duplicate_ii([1,2,3,1,2,3], 2))
        


if __name__ == "__main__":
    unittest.main()
