package arrays

import (
	"fmt"
	"testing"
)

func TestContainerMostWater(t *testing.T) {
	tests := map[string]struct {
		heights []int
		answer  int
	}{
		"test_1": {
			heights: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			answer:  49,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Container Most Water %s", name), func(t *testing.T) {
			result := maxArea(test.heights)
			if result != test.answer {
				t.Fatalf("Expected %d but got %d\n", test.answer, result)
			}
		})
	}
}
