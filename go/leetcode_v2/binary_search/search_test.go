package binarysearch

import (
	"fmt"
	"testing"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func TestSearch(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		result int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			result: 4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			result: -1,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 12,
			result: 5,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -1,
			result: 0,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 15,
			result: -1,
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			result: 2,
		},
		{
			nums:   []int{2, 4, 6, 8, 10},
			target: 8,
			result: 3,
		},
		{
			nums:   []int{-5, -3, 0, 7, 10},
			target: 0,
			result: 2,
		},
		{
			nums:   []int{-10, -5, 0, 5, 10},
			target: -10,
			result: 0,
		},
		{
			nums:   []int{1, 3, 5, 7, 9},
			target: 4,
			result: -1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Running search test case %d", i), func(t *testing.T) {
			if result := search(tc.nums, tc.target); result != tc.result {
				t.Fatalf("Expected %d but got %d", tc.result, result)
			}
		})
	}
}
