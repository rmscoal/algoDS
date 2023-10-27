package sorting

import (
	"sort"
)

/*
* Complete the 'maximumToys' function below.
*
* The function is expected to return an INTEGER.
* The function accepts following parameters:
*  1. INTEGER_ARRAY prices
*  2. INTEGER k
 */

func maximumToys(prices []int32, budget int32) int32 {
	// We are going to sort our array
	sort.SliceStable(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	// And then we can use the sliding window
	var max int32 = 0
	var cursor int32 = 0
	left := 0
	right := 0

	for right < len(prices) {
		cursor += prices[right]

		for cursor > budget {
			cursor -= prices[left]
			left++
		}

		if max <= int32(right-left+1) {
			max = int32(right - left + 1)
		}
		right++
	}

	return max
}
