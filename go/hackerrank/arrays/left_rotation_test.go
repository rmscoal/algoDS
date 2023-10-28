package arrays

import (
	"fmt"
	"testing"
)

func TestLeftRotation(t *testing.T) {
	tests := map[string]struct {
		array  []int32
		d      int32
		answer []int32
	}{
		"test_1": {
			array:  []int32{1, 2, 3, 4, 5},
			d:      4,
			answer: []int32{5, 1, 2, 3, 4},
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Left Rotation %s", name), func(t *testing.T) {
			result := rotLeft(test.array, test.d)
			for i, v := range result {
				if v != test.answer[i] {
					t.Fatalf("Expected at index %d is %d but got %d\n", i, test.answer[i], v)
				}
			}
		})
	}
}
