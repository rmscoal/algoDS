package leetcode

func (T *TwoSum) SlowSolver(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	panic("Error")
}

func (T *TwoSum) Solver(nums []int, target int) []int {
	mapper := make(map[int]int, len(nums))

	for index, value := range nums {
		if index_mapper, exist := mapper[target-value]; exist {
			return []int{index, index_mapper}
		} else {
			mapper[value] = index
		}
	}

	return []int{}
}
