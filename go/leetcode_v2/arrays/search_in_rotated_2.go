package arrays

/*
There is an integer array nums sorted in non-decreasing order (not necessarily
with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index
k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1],
..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example,
[0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if
target is in nums, or false if it is not in nums.

You must decrease the overall operation steps as much as possible.
*/

func search(nums []int, target int) bool {
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i] == target {
			return true
		}

		if nums[j] == target {
			return true
		}
	}

	return false
}
