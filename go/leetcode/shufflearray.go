package leetcode

func (S *ShuffleArray) Solver(nums []int, n int) []int {
	ans := make([]int, len(nums))

	for i, j := 0, 0; i < len(nums); i++ {
		if i%2 == 0 {
			ans[i] = nums[j]
		} else {
			ans[i] = nums[n+j]
			j++
		}
	}

	return ans
}
