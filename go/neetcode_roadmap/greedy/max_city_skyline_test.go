package greedy

import (
	"fmt"
	"testing"
)

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	tests := map[string]struct {
		grid   [][]int
		answer int
	}{
		"test_1": {
			grid:   [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}},
			answer: 35,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Max Increase Keeping Skyline %s\n", name), func(t *testing.T) {
			result := maxIncreaseKeepingSkyline(test.grid)
			if result != test.answer {
				t.Fatalf("Expected %d but got %d\n", test.answer, result)
			}
		})
	}
}
