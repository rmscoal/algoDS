package slidingwindow

import (
	"fmt"
	"testing"
)

func maxProfit(prices []int) int {
	maxProfit := 0

	// Sliding Window Approach

	left, right := 0, 1

	var diff int
	for left <= right && right < len(prices) {
		diff = prices[right] - prices[left]
		if diff < 0 {
			left++
		} else {
			maxProfit = max(maxProfit, diff)
			right++
		}
	}

	return maxProfit
}

func TestMaxProfit(t *testing.T) {
	testcases := []struct {
		prices []int
		result int
	}{
		{
			prices: []int{7, 6, 4, 3, 1},
			result: 0,
		},
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			result: 5,
		},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Running testcase %d", idx+1), func(t *testing.T) {
			result := maxProfit(tc.prices)
			if result != tc.result {
				t.Fatalf("Expected %d but got %d", tc.result, result)
			}
		})
	}
}
