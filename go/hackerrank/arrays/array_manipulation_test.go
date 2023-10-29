package arrays

import (
	"fmt"
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	tests := map[string]struct {
		n       int32
		queries [][]int32
		answer  int64
	}{
		"test_1": {
			n:       10,
			queries: [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}},
			answer:  10,
		},
		"test_2": {
			n:       10,
			queries: [][]int32{{1, 5, 3}, {4, 8, 8}, {6, 9, 1}},
			answer:  11,
		},
		"test_3": {
			n:       10,
			queries: [][]int32{{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15}},
			answer:  31,
		},
		"test_4": {
			n:       10,
			queries: [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}, {8, 10, 20}},
			answer:  28,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Array Manipulation %s", name), func(t *testing.T) {
			result := arrayManipulationWow(test.n, test.queries)
			if result != test.answer {
				t.Fatalf("Expectation is %d but got %d\n", test.answer, result)
			}
		})
	}
}
