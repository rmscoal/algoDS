from typing import List, Union


class Solution:
    def remove_element(self, nums: List[int], val: int) -> int:
        j = 0
        i = 0

        while i < len(nums) and j < len(nums):
            if nums[i] != val:
                nums[j], nums[i] = nums[i], nums[j]
                j += 1

            i += 1

        print("Nums", nums)
        return j

    def remove_element_v2(self, nums: List[Union[int, str]], val: int) -> int:
        i = 0
        j = len(nums) - 1

        while i < len(nums):
            if nums[i] == val:
                nums[i], nums[j] = nums[j], "_"
                j -= 1
            else:
                i += 1

        print("Nums", nums)
        return j + 1


if __name__ == "__main__":
    print(Solution().remove_element([3, 2, 3, 2, 3], 3))
    print(Solution().remove_element([0, 1, 2, 2, 3, 0, 4, 2], 2))
    print(Solution().remove_element_v2([3, 2, 3, 2, 3], 3))
    print(Solution().remove_element_v2([0, 1, 2, 2, 3, 0, 4, 2], 2))
