package leetcode

func (RE *RemoveElement) Solver(nums []int, val int) (int, []int) {
	var elementCount int = len(nums)

	for i := 0; i < elementCount; i++ {

		if nums[i] == val {
			nums[i] = nums[elementCount-1]
			nums[elementCount-1] = '_'
			elementCount--
			i--
		}
	}

	return elementCount, nums
}
