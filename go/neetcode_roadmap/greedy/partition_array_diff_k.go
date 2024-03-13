package greedy

import "sort"

func partitionArray(nums []int, k int) int {
	/*
		Input: nums = [3,6,1,2,5], k = 2

		- Sort the nums array
		[1,2,3,5,6]

		Then we can have left and right

		left = 0
		right = 1
		cursor = nums[right] - nums[left]
		if cursor > k {
			we would cut this off
			make left = right + 1
			and right = left + 1

			until we exhaust
		}
	*/

	// Result variable
	count := 1

	if len(nums) <= 1 {
		return count
	}

	// Sort the nums
	sort.Ints(nums)

	left := 0
	right := left + 1

	for right < len(nums) {
		if nums[right]-nums[left] > k {
			count++
			left = right
			right = left + 1
		} else {
			right++
		}
	}

	return count
}
