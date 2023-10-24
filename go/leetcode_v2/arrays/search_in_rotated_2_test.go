package arrays

import (
	"fmt"
	"testing"
)

func TestSearchInRotated2(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		answer bool
	}{
		"test_1": {
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			answer: true,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Search in Rotated Sorted Array II %s", name), func(t *testing.T) {
			result := search(test.nums, test.target)
			if result != test.answer {
				t.Fatalf("Expected %t but got %t\n", test.answer, result)
			}
		})
	}
}
