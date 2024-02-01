package array

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findPair_BruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{}
}

func findPair_Sort(nums []int, target int) []int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}

func findPair_Hash(nums []int, target int) []int {
	// We want to store the target - nums[i] here.
	// So that when we are iterating through the nums
	// we can easily identify existing pair.
	hash := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if pair, found := hash[nums[i]]; found {
			return []int{nums[i], pair}
		} else {
			hash[target-nums[i]] = nums[i]
		}
	}

	return []int{}
}

func TestFindPair_BruteForce(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		found  bool
		result []int
	}{
		{
			nums:   []int{8, 7, 2, 5, 4, 1},
			target: 10,
			found:  true,
			result: []int{8, 2},
		},
		{
			nums:   []int{5, 2, 6, 8, 1, 9},
			target: 12,
			found:  false,
			result: []int{},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Running testcase %d", i), func(t *testing.T) {
			result := findPair_BruteForce(tc.nums, tc.target)
			if tc.found {
				assert.Contains(t, result, tc.result[0])
				assert.Contains(t, result, tc.result[1])
			} else {
				assert.Empty(t, result)
			}
		})
		t.Run(fmt.Sprintf("Running testcase %d using sort", i), func(t *testing.T) {
			result := findPair_Sort(tc.nums, tc.target)
			if tc.found {
				assert.Contains(t, result, tc.result[0])
				assert.Contains(t, result, tc.result[1])
			} else {
				assert.Empty(t, result)
			}
		})
		t.Run(fmt.Sprintf("Running testcase %d using hash", i), func(t *testing.T) {
			result := findPair_Sort(tc.nums, tc.target)
			if tc.found {
				assert.Contains(t, result, tc.result[0])
				assert.Contains(t, result, tc.result[1])
			} else {
				assert.Empty(t, result)
			}
		})
	}
}
