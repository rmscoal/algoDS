package arrays

// 5 ms runtime in leetcode
func removeDuplicates(nums []int) int {
	j := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}

// 103 ms runtime in leetcode
func removeDuplicatesV2(nums []int) int {
	left := 0
	right := 1
	totalDuplicates := 0

	for left < len(nums) && right < len(nums) {
		if nums[left] == 101 || nums[right] == 101 {
			break
		}

		if nums[left] == nums[right] {
			nums = append(nums[:right], nums[right+1:]...)
			nums = append(nums, 101)
			totalDuplicates++
		} else {
			left++
			right++
		}
	}
	return len(nums) - totalDuplicates
}
