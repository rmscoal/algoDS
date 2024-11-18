package leetcode

import "testing"

// Webpage: https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/?envType=daily-question&envId=2024-11-16

func powerOfKSubarray(nums []int, k int) []int {
	// The idea is to use sliding window to iterate the array
	// and use a queue to store the indexes that disrupted the
	// condition.
	N := len(nums)
	result := make([]int, N-k+1)
	queue := make([]int, 0, k)

	left, right := 0, 0
	for left <= right && right < N {
		// Checks whether the current window is consecutive and in ascending order.
		if k > 1 && right > 0 {
			if nums[right] != nums[right-1]+1 {
				queue = append(queue, right-1)
			}
		}

		// Checks if the current window lenght is k. This could happen
		// during the inilization of the sliding window
		if right-left+1 != k {
			right++
			continue
		}

		// Checks if the queue is empty
		if len(queue) == 0 {
			result[left] = nums[right]
		} else {
			result[left] = -1
		}

		if len(queue) > 0 {
			if queue[0] == left {
				queue = queue[1:]
			}
		}
		left++
		right++
	}

	return result
}

func TestPowerOfKSubarray(t *testing.T) {
	testcases := []struct {
		nums   []int
		k      int
		result []int
	}{
		{
			nums:   []int{1, 3, 4},
			k:      2,
			result: []int{-1, 4},
		},
		{
			nums:   []int{1, 1, 1, 2},
			k:      2,
			result: []int{-1, -1, 2},
		},
		{
			nums:   []int{1, 4},
			k:      1,
			result: []int{1, 4},
		},
		{
			nums:   []int{2, 2, 2, 2, 2},
			k:      4,
			result: []int{-1, -1},
		},
		{
			nums:   []int{1, 2, 3, 4, 3, 2, 5},
			k:      3,
			result: []int{3, 4, -1, -1, -1},
		},
	}

	for _, testcase := range testcases {
		t.Run("Running testcase", func(t *testing.T) {
			result := powerOfKSubarray(testcase.nums, testcase.k)

			if len(result) != len(testcase.result) {
				t.Fatal("Expected result to be", testcase.result, "but got", result)
			}

			for i := 0; i < len(result); i++ {
				if result[i] != testcase.result[i] {
					t.Fatal("Expected result to be", testcase.result, "but got", result)
				}
			}
		})
	}
}
