package arrays

import "testing"

func productExceptSelfExtraMem(nums []int) []int {
	// Given:
	// nums = [1, 2, 3, 4]
	// We create a prefix product not containing self
	// which contains all product before the current
	// num value.
	// pref = [1, 1, 2, 6]
	// And also a suffix product that contains all the
	// product after the current num value.
	// suff = [24, 12, 4, 1]
	// Hence, the answer would be pref[i] * suff[i].
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	prefix[0] = 1
	suffix[len(nums)-1] = 1

	// Left to right
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	// Right to left
	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	// Result product
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}

func productExceptSelf(nums []int) []int {
	// Combine the above method to be in one
	// result array
	result := make([]int, len(nums))
	result[0] = 1

	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	curr := 1
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] *= nums[i+1] * curr
		curr *= nums[i+1]
	}

	return result
}

func TestProductExceptSelf(t *testing.T) {
	if result := productExceptSelfExtraMem([]int{1, 2, 3, 4}); !assertArray[int]([]int{24, 12, 8, 6}, result) {
		t.Fatalf("Wrong bro")
	}
	if result := productExceptSelf([]int{1, 2, 3, 4}); !assertArray[int]([]int{24, 12, 8, 6}, result) {
		t.Fatalf("Wrong bro")
	}
}
