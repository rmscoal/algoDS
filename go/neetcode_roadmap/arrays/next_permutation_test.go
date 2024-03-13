package arrays

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		answer []int
	}{
		"test_1": {
			nums:   []int{1, 2, 3},
			answer: []int{1, 3, 2},
		},
		"test_2": {
			nums:   []int{2, 3, 1},
			answer: []int{3, 1, 2},
		},
		"test_3": {
			nums:   []int{4, 2, 0, 2, 3, 2, 0},
			answer: []int{4, 2, 0, 3, 0, 2, 2},
		},
		"test_4": {
			nums:   []int{3, 2, 1},
			answer: []int{1, 2, 3},
		},
		"test_5": {
			nums:   []int{1, 1, 4, 3, 2},
			answer: []int{1, 2, 1, 3, 4},
		},
		"test_6": {
			nums:   []int{1, 4, 6, 5, 4},
			answer: []int{1, 5, 4, 4, 6},
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Next Permutation %s", name), func(t *testing.T) {
			nextPermutation(test.nums)
			for i, v := range test.nums {
				if test.answer[i] != v {
					t.Fatalf("Expected %v but got %v\n", test.answer, test.nums)
				}
			}
		})
	}
}
