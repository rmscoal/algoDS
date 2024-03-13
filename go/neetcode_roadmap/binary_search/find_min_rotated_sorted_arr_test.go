package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMin_Naive(nums []int) int {
	for {
		if nums[0] > nums[len(nums)-1] {
			nums = append(nums[1:], nums[0])
		} else {
			break
		}
	}

	return nums[0]
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	result := 5000

	for l <= r {
		if nums[l] < nums[r] {
			result = min(result, nums[l])
			return result
		}

		mid := (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid - 1
			result = min(result, nums[mid])
		} else { // it is impossible for right bigger than left when mid < right
			l = mid + 1
		}

		result = min(result, nums[mid])
	}

	return result
}

func TestFindMin(t *testing.T) {
	testcases := []struct {
		nums   []int
		result int
	}{
		{
			nums:   []int{3, 4, 5, 1, 2},
			result: 1,
		},
		{
			nums:   []int{4, 1, 3},
			result: 1,
		},
		{
			nums:   []int{3, 4, 5, 6, 7, 1, 2},
			result: 1,
		},
		{
			nums:   []int{5, 1, 4},
			result: 1,
		},
		{
			nums:   []int{4, 1, 3},
			result: 1,
		},
		{
			nums:   []int{5, 1, 2, 3, 4},
			result: 1,
		},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Running findMin_Naive test number %d", idx), func(t *testing.T) {
			result := findMin_Naive(tc.nums)
			assert.Equal(t, tc.result, result)
		})
		t.Run(fmt.Sprintf("Running findMin test number %d", idx), func(t *testing.T) {
			result := findMin(tc.nums)
			assert.Equal(t, tc.result, result)
		})
	}
}
