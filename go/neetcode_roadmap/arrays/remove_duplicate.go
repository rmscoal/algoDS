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

/*
Given an integer array nums sorted in non-decreasing order, remove some duplicates
in-place such that each unique element appears at most twice. The relative order
of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you
must instead have the result be placed in the first part of the array nums. More
formally, if there are k elements after removing the duplicates, then the first
k elements of nums should hold the final result. It does not matter what you leave
beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the
input array in-place with O(1) extra memory.

Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
*/

func removeDuplicates_2(nums []int) int {
	left := 0
	right := 0
	length := len(nums)
	count := 0

	for right < length && left <= right {
		if nums[right] == nums[left] {
			count++
			for count > 2 {
				nums = append(nums[:left], nums[left+1:]...)
				count--
				length--
				right--
			}
			right++
		} else {
			left = right
			count = 0
		}
	}
	return length
}
