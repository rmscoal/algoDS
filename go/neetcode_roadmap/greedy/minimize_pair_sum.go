package greedy

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minPairSum(nums []int) int {
	/*
		Input: nums = [3,5,4,2,4,6]

		- Sort nums
		[2, 3, 4, 4, 5, 6]

		maxSum := 0

		I can always do
		for i, j := 0, n; i < j; i, j = i + 1, j - 1 {
			maxSum = max(maxSum, nums[i] - nums[j])
		}


		Input: nums = [3,5,2,3]

		- Sort
		[2, 3, 3, 5]

		maxSum := 0

		for i, j := 0, n; i < j; i, j = i + 1, j - 1 {
			maxSum = max(maxSum, nums[i] - nums[j])
		}

		return maxSum
	*/

	sort.Ints(nums)
	maxSum := 0

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		maxSum = max(maxSum, nums[i]+nums[j])
	}

	return maxSum
}
