package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rotate-array/

func rotate(nums []int, k int) {
	k = k % len(nums)

	right := make([]int, 0, k)
	for i := len(nums) - k; i < len(nums); i++ {
		right = append(right, nums[i])
	}

	left := make([]int, 0, len(nums)-k)
	for i := 0; i < len(nums)-k; i++ {
		left = append(left, nums[i])
	}

	for i := 0; i < len(left); i++ {
		nums[i+k] = left[i]
	}

	for i := 0; i < k; i++ {
		nums[i] = right[i]
	}
}

func TestRotateArray(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			nums:     []int{-1},
			k:        2,
			expected: []int{-1},
		},
		{
			nums:     []int{-1, 2},
			k:        2,
			expected: []int{-1, 2},
		},
	}

	for _, tt := range tests {
		rotate(tt.nums, tt.k)
		assert.Equal(t, tt.expected, tt.nums)
	}
}
