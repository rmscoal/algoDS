package sorting

import (
	"fmt"
	"testing"
)

func TestCountInversion(t *testing.T) {
	tests := map[string]struct {
		input  []int32
		answer int64
	}{
		"test_1": {
			input:  []int32{2, 4, 1},
			answer: 2,
		},
		"test_2": {
			input:  []int32{7, 5, 3, 1},
			answer: 6,
		},
		"test_3": {
			input:  []int32{1, 1, 1, 2, 2},
			answer: 0,
		},
		"test_4": {
			input:  []int32{2, 1, 3, 1, 2},
			answer: 4,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Counting Inversion %s", name), func(t *testing.T) {
			result := countInversions(test.input)
			if result != test.answer {
				t.Fatalf("Expectation is %d but got %d\n", test.answer, result)
			}
		})
	}
}
