package arrays

import (
	"fmt"
	"sort"
	"testing"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	type kv struct {
		k, v int
	}

	arr := make([]kv, 0, len(freq))
	for key, count := range freq {
		arr = append(arr, kv{k: key, v: count})
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].v > arr[j].v
	})

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, arr[i].k)
	}

	return result
}

func TestTopKFrequent(t *testing.T) {
	fmt.Printf("Result %+v\n", topKFrequent([]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2}, 2))
}

func topKFrequent_Cleaner(nums []int, k int) []int {
	sort.Ints(nums)

	markedIndex := 0
	count := 1
	result := []int{}

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[markedIndex] {
			count++
		} else if count >= k {
			result = append(result, nums[markedIndex])
			markedIndex = i
			count = 1
		}
	}

	if count >= k {
		result = append(result, nums[markedIndex])
	}

	return result
}

func TestTopKFrequent_Cleaner(t *testing.T) {
	testcases := []struct {
		nums   []int
		k      int
		result []int
	}{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			k:      2,
			result: []int{1, 2},
		},
		{
			nums:   []int{1},
			k:      1,
			result: []int{1},
		},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Running testcase %d", idx), func(t *testing.T) {
			result := topKFrequent_Cleaner(tc.nums, tc.k)
			if len(result) != len(tc.result) {
				t.Fatalf("Expected to be %v but got %v", tc.result, result)
			}
			for i, v := range tc.result {
				if v != result[i] {
					t.Fatalf("Expected to be %v but got %v", tc.result, result)
				}
			}
		})
	}
}
