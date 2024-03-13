package binarysearch

import (
	"fmt"
	"math"
	"testing"
)

// Time Limit Exceed
func minEatingSpeed_Brute_Force(piles []int, h int) int {
	// Naive approach
	// We can find our lowest and greatest pile from the piles.
	// Then we brute force them starting k = lowest till highest.

	var highest int = math.MinInt

	for i := 0; i < len(piles); i++ {
		highest = max(highest, piles[i])
	}

	for k := 1; k <= highest; k++ {
		counter := 0

		for i := 0; i < len(piles); i++ {
			if piles[i] <= k {
				counter++
			} else {
				counter += piles[i] / k
				if piles[i]%k != 0 {
					counter++
				}
			}
		}

		if counter <= h {
			return k
		}
	}

	return 1
}

func minEatingSpeed_BinarySearch(piles []int, h int) int {
	// Find the maximum number of pile from the given piles
	var lowest, highest int = 1, 0
	for _, pile := range piles {
		highest = max(highest, pile)
	}
	var ans int = highest

	for lowest <= highest {
		speed := (highest + lowest) / 2
		time := 0

		for _, pile := range piles {
			if pile <= speed {
				time++
			} else {
				time += pile / speed
				if pile%speed != 0 {
					time++
				}
			}
		}

		if time <= h {
			ans = min(ans, speed)
			highest = speed - 1
		} else {
			lowest = speed + 1
		}
	}

	return ans
}

func TestMinEatingSpeed(t *testing.T) {
	testcases := []struct {
		piles  []int
		h      int
		result int
	}{
		{
			piles:  []int{3, 6, 7, 11},
			h:      8,
			result: 4,
		},
		{
			piles:  []int{30, 11, 23, 4, 20},
			h:      5,
			result: 30,
		},
		{
			piles:  []int{30, 11, 23, 4, 20},
			h:      6,
			result: 23,
		},
		{
			piles:  []int{312884470},
			h:      312884469,
			result: 2,
		},
		{
			piles:  []int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184},
			h:      823855818,
			result: 14,
		},
		{
			piles:  []int{1, 1, 1, 999999999},
			h:      10,
			result: 142857143,
		},
		{
			piles:  []int{805306368, 805306368, 805306368},
			h:      1000000000,
			result: 3,
		},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Running testcase %d using brute force solution", idx), func(t *testing.T) {
			if result := minEatingSpeed_Brute_Force(tc.piles, tc.h); result != tc.result {
				t.Fatalf("Expected is %d but got %d", tc.result, result)
			}
		})

		t.Run(fmt.Sprintf("Running testcase %d using binary search solution", idx), func(t *testing.T) {
			if result := minEatingSpeed_BinarySearch(tc.piles, tc.h); result != tc.result {
				t.Fatalf("Expected is %d but got %d", tc.result, result)
			}
		})
	}
}
