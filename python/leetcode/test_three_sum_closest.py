import unittest

def three_sum_closest(nums, target) -> int:
    nums.sort()
    totals = set()

    for i in range(len(nums) - 2):
        left, right = i + 1, len(nums) - 1
        while left < right:
            total = nums[i] + nums[left] + nums[right]
            if total == target:
                return target
            elif total < target:
                left += 1
            else:
                right -= 1
            totals.add(total)

    result = totals.pop()
    min_distance = abs(target - result)
    for tt in totals:
        dst = abs(target - tt)
        if dst < min_distance:
            result = tt
            min_distance = dst

    return result

class ThreeSumClosest(unittest.TestCase):
    def test_three_sum_closest(self):
        self.assertEqual(2, three_sum_closest([-1,2,1,-4], 1))
        self.assertEqual(0, three_sum_closest([0,0,0], 1))
        self.assertEqual(3, three_sum_closest([0,1,2], 0))
        self.assertEqual(0, three_sum_closest([-4,2,2,3,3,3], 0))


if __name__ == '__main__':
    unittest.main()
