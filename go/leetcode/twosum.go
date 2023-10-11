package leetcode

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
