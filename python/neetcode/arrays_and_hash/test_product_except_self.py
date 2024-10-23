from typing import List
import unittest


class TestProductExceptSelf(unittest.TestCase):
    def product_except_self_other(self, nums: List[int]) -> List[int]:
        """
        :type nums: List[int]
        :rtype: List[int]

        Description:
            Create a prefix and suffix array which contains the
            product of its prefix and suffix. Then multiply both
            of them to get the resulting array.
        """
        N = len(nums)
        prefix = [1 for _ in range(N)]
        suffix = [1 for _ in range(N)]
        answer = [1 for _ in range(N)]

        for i in range(1, N):
            prefix[i] = prefix[i - 1] * nums[i - 1]
        for i in range(N - 2, - 1, -1):
            suffix[i] = suffix[i + 1] * nums[i + 1]
        for i in range(N):
            answer[i] = prefix[i] * suffix[i]

        return answer

    def product_except_self(self, nums: List[int]) -> List[int]:
        """
        :type nums: List[int]
        :rtype: List[int]

        Description:
            Given: 
                [1, 2, 3, 4]
            Create a new empty array filled with ones:
                [1, 1, 1, 1]
            Initialize carry = 1

            Traverse the array from left to right. For each loop,
            multiply answer[i] with carry then multiply carry with
            nums[i].

            Traverse the array from right to left and do the same like
            the above.
        """

        N = len(nums)
        answer = [1 for _ in range(N)]
        carry = 1

        # Traverse from left to right
        for i in range(N):
            answer[i] *= carry
            carry *= nums[i]

        # Reset carry and traverse from right to left
        carry = 1
        for i in range(N - 1, -1, -1):
            answer[i] *= carry
            carry *= nums[i]

        return answer

    def test_product_except_self(self):
        self.assertEqual(self.product_except_self_other(
            [1, 2, 3, 4]), [24, 12, 8, 6])
        self.assertEqual(self.product_except_self_other(
            [-1, 1, 0, -3, 3]), [0, 0, 9, 0, 0])
        self.assertEqual(self.product_except_self(
            [1, 2, 3, 4]), [24, 12, 8, 6])
        self.assertEqual(self.product_except_self(
            [-1, 1, 0, -3, 3]), [0, 0, 9, 0, 0])


if __name__ == "__main__":
    unittest.main()
