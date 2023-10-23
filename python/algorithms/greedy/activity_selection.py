"""
You are given n activities with their start and finish times. Select the
maximum number of activities that can be performed by a single person, assuming
that a person can only work on a single activity at a time.

Example 1:
    Input: start[]  =  {10, 12, 20}, finish[] =  {20, 25, 30}
    Output: 0 2
    Explanation: A person can perform at most two activities. The
    maximum set of activities that can be executed
    is {0, 2} [ These are indexes in start[] and finish[] ]
"""

# User function Template for python3

from typing import List


class Solution:
    # Function to find the maximum number of activities that can
    # be performed by a single person.
    def activity_selection(self, n: int, start: List[int], end: List[int]) -> int:
        # Combine start and end to 1
        combined = []
        for i in range(n):
            combined.append([start[i], end[i]])

        # Sort by end date
        combined.sort(key=lambda x: x[1])

        count = 0
        end_date = -1

        for i in range(n):
            if combined[i][0] > end_date:
                count += 1
                end_date = combined[i][1]

        return count


if __name__ == "__main__":
    print("Result", Solution().activity_selection(
        4, [1, 3, 2, 5], [2, 4, 3, 6]))
    print()
    print("Result", Solution().activity_selection(2, [2, 1], [2, 2]))
    print()
    print("Result", Solution().activity_selection(
        5, [7, 6, 2, 7, 3], [10, 6, 10, 9, 5]))
