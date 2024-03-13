package arrays

import "sort"

func nextPermutation(nums []int) {
	j := len(nums) - 1
	for ; j > 0; j-- {
		if nums[j-1] < nums[j] {
			for l := len(nums) - 1; l >= j; l-- {
				if nums[l] > nums[j-1] {
					nums[j-1], nums[l] = nums[l], nums[j-1]
					break
				}
			}
			break
		}
	}

	sort.Ints(nums[j:])
}
