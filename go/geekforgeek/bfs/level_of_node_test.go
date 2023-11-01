package bfs

import (
	"fmt"
	"testing"
)

func TestLevelOfNode(t *testing.T) {
	tests := map[string]struct {
		edges                        [][]int
		startingNode, target, answer int
	}{
		"test_1": {
			edges:        [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}},
			startingNode: 0,
			target:       4,
			answer:       2,
		},
		"test_2": {
			edges:        [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}},
			startingNode: 0,
			target:       0,
			answer:       0,
		},
		"test_3": {
			edges:        [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}},
			startingNode: 0,
			target:       1,
			answer:       1,
		},
		"test_4": {
			edges:        [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}},
			startingNode: 0,
			target:       1,
			answer:       1,
		},
		"test_5": {
			edges:        [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}},
			startingNode: 0,
			target:       2,
			answer:       1,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Level of Node %s", name), func(t *testing.T) {
			result := nodeLevel(test.edges, test.startingNode, test.target)
			if test.answer != result {
				t.Fatalf("Expectation is %d but got %d\n", test.answer, result)
			}
		})
	}
}
