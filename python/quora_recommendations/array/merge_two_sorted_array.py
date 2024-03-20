import unittest
from typing import List


def merge(nums1: List[int], nums2: List[int]) -> None:
    if len(nums1) == 0 or len(nums2) == 0:
        return

    ptr1 = 0
    while ptr1 < len(nums1):
        if nums1[ptr1] > nums2[0]:
            nums1[ptr1], nums2[0] = nums2[0], nums1[ptr1]

            # Swap only if it is not sorted and not last element
            x = 0
            while x < len(nums2) - 1 and nums2[x] > nums2[x + 1]:
                nums2[x], nums2[x + 1] = nums2[x + 1], nums2[x]
                x += 1
        else:
            ptr1 += 1


class Test(unittest.TestCase):
    def test_merge(self):
        num1 = [1, 4, 7, 8, 10]
        num2 = [2, 3, 9]

        merge(num1, num2)

        self.assertEqual([1, 2, 3, 4, 7], num1)
        self.assertEqual([8, 9, 10], num2)


if __name__ == "__main__":
    unittest.main()
