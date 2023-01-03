package leetcode

import (
	"sort"
)

func (RD *RemoveDuplicate) Solver(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	var targetTemp int = nums[0]
	var elementCount int = len(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == targetTemp {
			nums[i] = 101
			elementCount--
		} else {
			targetTemp = nums[i]
		}
	}

	sort.Ints(nums)
	return elementCount
}

func (RD *RemoveDuplicate) FastSolution(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
