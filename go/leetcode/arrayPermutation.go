package leetcode

func (A *ArrayPermuatation) Solver(nums []int) []int {
	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}

	return ans
}

func (A *ArrayPermuatation) OtherSolver(nums []int) []int {
	ans := make([]int, len(nums))

	for i, n := range nums {
		ans[i] = nums[n]
	}

	return ans
}
