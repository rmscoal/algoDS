package arrays

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	tests := map[string]struct {
		inputs   []int
		expected int
	}{
		"test_1": {
			inputs:   []int{1, 1, 2},
			expected: 2,
		},
		"test_2": {
			inputs:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for name, cases := range tests {
		t.Run(fmt.Sprintf("Remove duplicate v1 %s", name), func(t *testing.T) {
			input := make([]int, len(cases.inputs))
			copy(input, cases.inputs)
			result := removeDuplicates(input)
			if result != cases.expected {
				t.Fatalf("Expected %d but got %d", cases.expected, result)
			}
		})
	}

	for name, cases := range tests {
		t.Run(fmt.Sprintf("Remove duplicate v2 %s", name), func(t *testing.T) {
			input := make([]int, len(cases.inputs))
			copy(input, cases.inputs)
			result := removeDuplicatesV2(input)
			if result != cases.expected {
				t.Fatalf("Expected %d but got %d", cases.expected, result)
			}
		})
	}
}
