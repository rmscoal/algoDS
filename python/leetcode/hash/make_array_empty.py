from typing import List

from collections import Counter, defaultdict


class Solution:
    def minOperationsSlow(self, nums: List[int]) -> int:
        count_map = {}
        ways = 0
        for i in nums:
            if i in count_map:
                count_map[i] += 1
            else:
                count_map[i] = 1

        for key in count_map:
            count = count_map[key]
            found = False
            for i in range(count // 3, -1, -1):
                for j in range(count // 2, -1, -1):
                    if ((3 * i) + (2 * j)) == count:
                        found = True
                        ways += i + j
                        break
                if found:
                    break
            if not found:
                return -1

        return ways

    def minOperationsCounter(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        if cnt.most_common()[-1][1] == 1:
            return -1
        return sum((c-1) // 3 + 1 for c in cnt.values())

    def minOperations(self, nums: List[int]) -> int:
        """
        1   2   3   4   5   6   7   8   9   10  11  12  13
        -1  1   1   2   2   2   3   3   3   4   4   4   5

        Either ((c-1)//3) + 1 OR (c + 2)//3
        """
        count_map = defaultdict(int)
        ways = 0
        for i in nums:
            count_map[i] += 1
        for key in count_map:
            count = count_map[key]
            if count == 1:
                return -1
            ways += (count + 2) // 3
        return ways


if __name__ == "__main__":
    print(Solution().minOperationsSlow(nums=[2, 3, 3, 2, 2, 4, 2, 3, 4]))
    print(Solution().minOperationsCounter(nums=[2, 3, 3, 2, 2, 4, 2, 3, 4]))
    print(Solution().minOperations(nums=[2, 3, 3, 2, 2, 4, 2, 3, 4]))
