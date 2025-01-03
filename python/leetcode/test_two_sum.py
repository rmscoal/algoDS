import unittest


def two_sum(nums, target):
    N = len(nums)
    look_up_map = {}

    for i in range(N):
        curr = nums[i]
        if curr in look_up_map:
            ans = [i, look_up_map[curr]]
            ans.sort()
            return ans
        key = target - curr
        look_up_map[key] = i

    return []


class TwoSum(unittest.TestCase):
    def test(self):
        self.assertEqual([0,1], two_sum([2,7,11,15], 9))
        self.assertEqual([1,2], two_sum([3,2,4], 6))
        self.assertEqual([0,1], two_sum([3,3], 6))

if __name__ == '__main__':
    unittest.main()
