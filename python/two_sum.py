from typing import List

class Solution:
    @staticmethod
    def two_sum(nums: List[int], target: int) -> List[int]:
        mapper = {}
        res = []

        for i in range(0, len(nums)):
            current = nums[i]
            if current in mapper:
                res.extend([i, mapper[current]])
            else:
                mapper[target-current] = i

        return res

if __name__ == "__main__":
    print('Two Sum', Solution.two_sum([2,7,11,15], 9))
    print('Two Sum', Solution.two_sum([3,2,4], 6))
    print('Two Sum', Solution.two_sum([3,3], 6))