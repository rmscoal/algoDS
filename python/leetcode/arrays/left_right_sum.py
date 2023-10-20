"""
Given a 0-indexed integer array nums, find a 0-indexed integer array answer 
where:

- answer.length == nums.length.
- answer[i] = |leftSum[i] - rightSum[i]|.
Where:
- leftSum[i] is the sum of elements to the left of the index i in the array nums.
If there is no such element, leftSum[i] = 0.
- rightSum[i] is the sum of elements to the right of the index i in the array 
nums. If there is no such element, rightSum[i] = 0.

Return the array answer.
"""

from typing import List


class Solution:
    def left_right_difference(self, nums: List[int]) -> List[int]:
        leftSum = []
        rightSum = []
        answer = [0 for _ in range(len(nums))]

        for i in range(len(nums)):
            if i != 0:
                total = 0
                for j in range(0, i):
                    total += nums[j]
                leftSum.append(total)
            else:
                leftSum.append(0)

            if i != len(nums) - 1:
                total = 0
                for j in range(i+1, len(nums)):
                    total += nums[j]
                rightSum.append(total)
            else:
                rightSum.append(0)

        for i in range(len(answer)):
            answer[i] = abs(leftSum[i] - rightSum[i])

        return answer

    def left_right_difference_v2(self, nums: List[int]) -> List[int]:
        left_cursor = 0
        right_cursor = 0
        answer = [0 for _ in range(len(nums))]

        for i in range(1, len(nums)):
            right_cursor += nums[i]

        for i in range(len(nums)):
            if i != 0:
                left_cursor += nums[i-1]
                right_cursor -= nums[i]

            answer[i] = abs(left_cursor - right_cursor)

        return answer


if __name__ == "__main__":
    print(Solution().left_right_difference([10, 4, 8, 3]))
    print(Solution().left_right_difference([1]))
    print(Solution().left_right_difference_v2([10, 4, 8, 3]))
    print(Solution().left_right_difference_v2([1]))
