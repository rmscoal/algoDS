package arrays

import (
	"fmt"
	"testing"
)

func TestMinimumBribes(t *testing.T) {
	tests := map[string]struct {
		q      []int32
		answer int64
	}{
		"test_1": {
			q:      []int32{2, 1, 5, 3, 4},
			answer: 3,
		},
		"test_2": {
			q:      []int32{2, 5, 1, 3, 4},
			answer: -1,
		},
		"test_3": {
			q:      []int32{1, 2, 5, 3, 4, 7, 8, 6},
			answer: 4,
		},
		"test_4": {
			q:      []int32{1, 2, 5, 3, 7, 8, 6, 4},
			answer: 7,
		},
		"test_5": {
			q:      []int32{5, 1, 2, 3, 7, 8, 6, 4},
			answer: -1,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Minimum Bribes %s", name), func(t *testing.T) {
			result := minimumBribes(test.q)
			if result != test.answer {
				t.Fatalf("Expected is %d but got %d\n", test.answer, result)
			}
		})
	}
}
