package bfs

import (
	"fmt"
	"testing"
)

func TestMinMaxAdjHeightDiff(t *testing.T) {
	tests := map[string]struct {
		matrix [][]int
		answer int
	}{
		"test_1": {
			matrix: [][]int{{1, 2}, {2, 1}},
			answer: 1,
		},
		"test_2": {
			matrix: [][]int{{1, 2, 1}, {2, 8, 2}, {2, 4, 2}},
			answer: 1,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Min Max Adjacent Height Diff %s", name), func(t *testing.T) {
			result := minAdjHeightDiff(test.matrix)
			if result != test.answer {
				t.Fatalf("Expected is %d but got %d\n", test.answer, result)
			}
		})
	}
}
