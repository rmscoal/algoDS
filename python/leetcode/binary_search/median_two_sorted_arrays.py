import unittest
from typing import List


class Solution:
    def find_median_sorted_arrays(self, nums1: List[int], nums2: List[int]) -> float:
        length = len(nums1) + len(nums2)
        left, right, merged_index = -1, -1, -1
        left_moved, right_moved = False, False
        target = (length - 1) // 2
        if length % 2 == 0:
            target_2 = target + 1
            value_tmp = 0

        while left < len(nums1) and right < len(nums2):
            # Checks matched targeted index
            if length % 2 == 1:
                if merged_index == target and left_moved:
                    return float(nums1[left])
                elif merged_index == target and right_moved:
                    return float(nums2[right])
            else:
                if merged_index == target and left_moved:
                    value_tmp = nums1[left]
                elif merged_index == target and right_moved:
                    value_tmp = nums2[right]
                elif merged_index == target_2 and left_moved:
                    return float((value_tmp + nums1[left]) / 2)
                elif merged_index == target_2 and right_moved:
                    return float((value_tmp + nums2[right]) / 2)

            # Checks bound
            if right + 1 == len(nums2):
                left += 1
                left_moved = True
                right_moved = False
                merged_index += 1
                continue
            elif left + 1 == len(nums1):
                right += 1
                left_moved = False
                right_moved = True
                merged_index += 1
                continue

            # Solves which to move next
            if left == -1 and right == -1:
                if nums1[0] < nums2[0]:
                    left += 1
                    left_moved = True
                    right_moved = False
                else:
                    right += 1
                    left_moved = False
                    right_moved = True
            elif left == -1 or right == -1:
                if left == -1:
                    if nums1[0] < nums2[right + 1]:
                        left += 1
                        left_moved = True
                        right_moved = False
                    else:
                        right += 1
                        right_moved = True
                        left_moved = False
                else:
                    if nums1[left+1] < nums2[0]:
                        left += 1
                        left_moved = True
                        right_moved = False
                    else:
                        right += 1
                        right_moved = True
                        left_moved = False
            elif nums1[left + 1] < nums2[right + 1]:
                left += 1
                left_moved = True
                right_moved = False
            else:
                right += 1
                left_moved = False
                right_moved = True

            merged_index += 1

        return 1

    def find_median_sorted_arrays_cleaner(self, nums1: List[int], nums2: List[int]) -> float:
        n = len(nums1)
        m = len(nums2)
        i = 0
        j = 0
        m1 = 0
        m2 = 0

        # Find median.
        for _ in range(0, (n + m) // 2 + 1):
            m2 = m1
            if i < n and j < m:
                if nums1[i] > nums2[j]:
                    m1 = nums2[j]
                    j += 1
                else:
                    m1 = nums1[i]
                    i += 1
            elif i < n:
                m1 = nums1[i]
                i += 1
            else:
                m1 = nums2[j]
                j += 1

        # Check if the sum of n and m is odd.
        if (n + m) % 2 == 1:
            return float(m1)
        else:
            ans = float(m1) + float(m2)
            return ans / 2.0

    def find_median_sorted_arrays_sort(self, nums1: List[int], nums2: List[int]) -> float:
        num = nums1 + nums2
        N = len(num)

        num.sort()
        if N % 2 == 1:
            return float(num[N//2])
        else:
            return (num[N//2 - 1] + num[N//2]) / 2.0

    def find_median_sorted_arrays_binary_search(self, nums1: list[int], nums2: list[int]) -> float:
        A, B = nums1, nums2
        total = len(nums1) + len(nums2)
        half = total // 2

        if len(B) < len(A):
            A, B = B, A

        l, r = 0, len(A) - 1
        while True:
            i = (l+r) // 2  # A
            j = half - i - 2  # B

            Aleft = A[i] if i >= 0 else float('-infinity')
            Aright = A[i+1] if i + 1 < len(A) else float('infinity')
            Bleft = B[j] if j >= 0 else float('-infinity')
            Bright = B[j+1] if j + 1 < len(B) else float('infinity')

            if Aleft <= Bright and Bleft <= Aright:
                if total % 2 == 1:
                    return float(min(Aright, Bright))
                else:
                    return float(max(Aleft, Bleft) + min(Aright, Bright)) / 2
            elif Aleft > Bright:
                r = i - 1
            else:
                l = i + 1


class TestSolution(unittest.TestCase):
    def test_find_median_sorted_arrays(self):
        self.assertEqual(
            2, int(Solution().find_median_sorted_arrays([1, 3], [2])))
        self.assertEqual(
            float((2+3)/2), Solution().find_median_sorted_arrays([1, 2], [3, 4]))
        self.assertEqual(
            float((2+3)/2), Solution().find_median_sorted_arrays([1, 3], [2, 7]))

        self.assertEqual(
            2, int(Solution().find_median_sorted_arrays_binary_search([1, 3], [2])))
        self.assertEqual(
            float((2+3)/2), Solution().find_median_sorted_arrays_binary_search([1, 3], [2, 7]))
        self.assertEqual(
            float(2), Solution().find_median_sorted_arrays_binary_search([1, 2, 3], []))


if __name__ == "__main__":
    unittest.main()
