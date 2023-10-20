"""
Given n points on a 2D plane where points[i] = [xi, yi], Return the widest 
vertical area between two points such that no points are inside the area.

A vertical area is an area of fixed-width extending infinitely along the y-axis
(i.e., infinite height). The widest vertical area is the one with the maximum
width.

Note that points on the edge of a vertical area are not considered included in
the area.
"""

from typing import List, Tuple
import random


class Solution:
    def split_with_pivot(self, nums: List[int]) -> Tuple[List[int], int, List[int]]:
        pivot = nums.pop(random.randint(0, len(nums) - 1))
        left = []
        right = []

        for i in nums:
            if i < pivot:
                left.append(i)
            else:
                right.append(i)

        return left, pivot, right

    def quicksort(self, nums: List[int]) -> List[int]:
        if len(nums) <= 1:
            return nums

        left, pivot, right = self.split_with_pivot(nums)

        left = self.quicksort(left)
        right = self.quicksort(right)

        return left + [pivot] + right

    def max_width_of_vertical_area(self, points: List[List[int]]) -> int:
        # Input -> [[8,7],[9,9],[7,4],[9,7]]

        y_points = []

        for i in points:
            y_points.append(i[0])

        # y_points -> [8, 9, 7, 9]

        y_points = self.quicksort(y_points)

        # y_points -> [7, 8, 9, 9]

        max_width = 0
        for i in range(1, len(y_points)):
            max_width = max(y_points[i] - y_points[i-1], max_width)

        return max_width


if __name__ == "__main__":
    print(Solution().max_width_of_vertical_area(
        [[8, 7], [9, 9], [7, 4], [9, 7]]))

    print(Solution().max_width_of_vertical_area(
        [[3, 1], [9, 0], [1, 0], [1, 4], [5, 3], [8, 8]]
    ))
