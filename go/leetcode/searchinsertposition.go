package leetcode

func (S *SearchInsertPosition) Solver(nums []int, target int) int {
	var i, j int
	var len int = len(nums) - 1
	for i, j = 0, len; i < j; i, j = i+1, j-1 {
		if nums[i] == target {
			return i
		} else if nums[j] == target {
			return j
		} else {
			if nums[i] > target {
				return i
			} else if nums[j] < target {
				return j + 1
			}
		}
	}

	if nums[i] < target {
		return i + 1
	}
	return i
}
