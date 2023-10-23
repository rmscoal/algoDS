package greedy

import (
	"fmt"
	"testing"
)

func TestMinimizePairSum(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		answer int
	}{
		"test_1": {
			nums:   []int{3, 5, 2, 3},
			answer: 7,
		},
		"test_2": {
			nums:   []int{3, 5, 4, 2, 4, 6},
			answer: 8,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Minimize maximum Pair Sum in Array %s\n", name), func(t *testing.T) {
			result := minPairSum(test.nums)
			if result != test.answer {
				t.Fatalf("Expected %d but got %d\n", test.answer, result)
			}
		})
	}
}
