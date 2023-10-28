package arrays

import (
	"fmt"
	"testing"
)

func TestMinimumSwaps(t *testing.T) {
	tests := map[string]struct {
		arr    []int32
		answer int32
	}{
		"test_1": {
			arr:    []int32{4, 3, 1, 2},
			answer: 3,
		},
		"test_2": {
			arr:    []int32{7, 1, 3, 2, 4, 5, 6},
			answer: 5,
		},
		"test_3": {
			arr:    []int32{1, 3, 5, 2, 4, 6, 7},
			answer: 3,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Minimum Swap %s", name), func(t *testing.T) {
			result := minimumSwaps(test.arr)
			if result != test.answer {
				t.Fatalf("Expectation is %d but got %d\n", test.answer, result)
			}
		})
	}
}
