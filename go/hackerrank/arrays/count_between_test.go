package arrays

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// findMinimumIndex finds i such that arr[i] < k and arr[i+1:] >= k.
func findMinimumIndex(arr []int32, k int32) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		curr := arr[mid]

		if curr >= k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func TestFindMinimumIndex(t *testing.T) {
	tests := []struct {
		arr      []int32
		k        int32
		expected int
	}{
		{
			arr:      []int32{1, 2, 2, 3, 4},
			k:        2,
			expected: 1,
		},
		{
			arr:      []int32{1, 2, 2, 3, 4},
			k:        0, // Out of bound
			expected: 0,
		},
		{
			arr:      []int32{1, 2, 2, 3, 3, 3, 3, 3, 4, 5, 6},
			k:        3,
			expected: 3,
		},
		{
			arr:      []int32{1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 5, 6},
			k:        1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected,
			findMinimumIndex(tt.arr, tt.k))
	}
}

// findMaximumIndex finds i such that arr[i] > k and arr[:i] <= k.
func findMaximumIndex(arr []int32, k int32) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		curr := arr[mid]

		if curr <= k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func TestFindMaximumIndex(t *testing.T) {
	tests := []struct {
		arr      []int32
		k        int32
		expected int
	}{
		{
			arr:      []int32{1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 5, 6},
			k:        2,
			expected: 3,
		},
		{
			arr:      []int32{1, 2, 2, 5, 6},
			k:        6,
			expected: 5,
		},
		{
			arr:      []int32{1, 2, 2, 5, 6},
			k:        7, // Out of bound
			expected: 5,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected,
			findMaximumIndex(tt.arr, tt.k))
	}
}

// countBetween is finding the number of elements between the given ranges
// from the array.
func countBetween(arr []int32, low []int32, high []int32) []int32 {
	// Binary search

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	result := make([]int32, len(low))
	for i := 0; i < len(low); i++ {
		lowerIndex := findMinimumIndex(arr, low[i])
		upperIndex := findMaximumIndex(arr, high[i])
		result[i] = int32(upperIndex - lowerIndex)
	}

	return result
}

func TestCountBetween(t *testing.T) {
	tests := []struct {
		arr     []int32
		low     []int32
		high    []int32
		expects []int32
	}{
		{
			arr:     []int32{1, 2, 2, 3, 4},
			low:     []int32{0, 2},
			high:    []int32{2, 4},
			expects: []int32{3, 4},
		},
		{
			arr:     []int32{10},
			low:     []int32{5, 10, 15},
			high:    []int32{10, 10, 20},
			expects: []int32{1, 1, 0},
		},
		{
			arr:     []int32{5, 10, 15, 20, 25},
			low:     []int32{0},
			high:    []int32{30},
			expects: []int32{5},
		},
		{
			arr:     []int32{1, 3, 5, 7, 9},
			low:     []int32{10, 20},
			high:    []int32{15, 25},
			expects: []int32{0, 0},
		},
		{
			arr:     []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
			low:     []int32{2, 5, 3},
			high:    []int32{6, 9, 4},
			expects: []int32{5, 5, 2},
		},
		{
			arr:     []int32{2, 2, 3, 3, 3, 5, 6},
			low:     []int32{2, 3},
			high:    []int32{3, 6},
			expects: []int32{5, 5},
		},
		{
			arr:     []int32{10, 20, 30, 40, 50},
			low:     []int32{15, 35},
			high:    []int32{25, 45},
			expects: []int32{1, 1},
		},
		{
			arr:     []int32{1, 1, 1, 1, 1},
			low:     []int32{1, 0},
			high:    []int32{1, 2},
			expects: []int32{5, 5},
		},
		{
			arr:     []int32{},
			low:     []int32{1, 2},
			high:    []int32{3, 4},
			expects: []int32{0, 0},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expects, countBetween(tt.arr, tt.low, tt.high))
	}
}

func BenchmarkCountBetween(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tests := []struct {
			arr     []int32
			low     []int32
			high    []int32
			expects []int32
		}{
			{
				arr:     []int32{1, 2, 2, 3, 4},
				low:     []int32{0, 2},
				high:    []int32{2, 4},
				expects: []int32{3, 4},
			},
		}

		for _, tt := range tests {
			countBetween(tt.arr, tt.low, tt.high)
		}
	}
}
