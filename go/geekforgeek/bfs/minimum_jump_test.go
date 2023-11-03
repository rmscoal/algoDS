package bfs

import (
	"fmt"
	"testing"
)

func TestMinimumJumps(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		answer int
	}{
		"test_1": {
			nums:   []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404},
			answer: 3,
		},
		"test_2": {
			nums:   []int{7, 6, 9, 6, 9, 6, 9, 7},
			answer: 1,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Minimum Jumps %s", name), func(t *testing.T) {
			result := minimumJumps(test.nums)
			if result != test.answer {
				t.Fatalf("Expected is %d but got %d\n", test.answer, result)
			}
		})
	}
}
