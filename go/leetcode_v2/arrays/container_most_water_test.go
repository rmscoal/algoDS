package arrays

import (
	"fmt"
	"testing"
)

func maxArea(heights []int) int {
	result := 0

	left, right := 0, len(heights)-1

	for left < right {
		result = max(result, (right-left)*min(heights[left], heights[right]))

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

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
