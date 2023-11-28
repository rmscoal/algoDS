package dp

import "testing"

func maxProduct(nums []int) int {
	/*
		[2,3,-2,4]
		Our dp array
		[1,1, 1,1, 1,1, 1,1]
		[2,2, 6,3, -2,-12, 4,-48]

		[-2,0,-1]
		Our dp array
		[1,1, 1,1, 1,1]
		[-2,-2, 0,0, 0,-1]
	*/
	if len(nums) <= 1 {
		return nums[0]
	}

	// This slice will contain the max-min pair value
	// at each i  and i + 1 respectively which looks
	// at the previous max-min pair hence dynamic.
	dp := make([]int, len(nums)*2)
	// Since the first pair doesn't have a pre-product
	// we make the first element of nums as the max-min
	// pair.
	dp[0], dp[1] = nums[0], nums[0]
	// Our result variable
	biggest := dp[0]

	for i := 1; i < len(nums); i++ {
		j := i * 2
		current := nums[i]
		max, min := current, current

		if current*dp[j-2] < current*dp[j-1] {
			max = findMax(max, current*dp[j-1])
			min = findMin(min, current*dp[j-2])
		} else {
			max = findMax(max, current*dp[j-2])
			min = findMin(min, current*dp[j-1])
		}

		if max > biggest {
			biggest = max
		}

		dp[j] = max
		dp[j+1] = min
	}

	return biggest
}

func TestMaxProductSubarray(t *testing.T) {
	if result := maxProduct([]int{2, 3, -2, 4}); result != 6 {
		t.Fatalf("Expected 6 but got %d", result)
	}

	if result := maxProduct([]int{-2, 0, -1}); result != 0 {
		t.Fatalf("Expected 0 but got %d", result)
	}

	if result := maxProduct([]int{-1, -2, -3}); result != 6 {
		t.Fatalf("Expected 6 but got %d", result)
	}
}
