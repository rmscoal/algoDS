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
