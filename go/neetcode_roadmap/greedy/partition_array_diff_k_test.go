package greedy

import (
	"fmt"
	"testing"
)

func TestPartitionArray(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		k      int
		answer int
	}{
		"test_1": {
			nums:   []int{3, 6, 1, 2, 5},
			k:      2,
			answer: 2,
		},
		"test_2": {
			nums:   []int{1, 2, 3},
			k:      1,
			answer: 2,
		},
		"test_3": {
			nums:   []int{2, 2, 4, 5},
			k:      0,
			answer: 3,
		},
		"test_4": {
			nums:   []int{2, 100},
			k:      1,
			answer: 2,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Partition Array Such That Maximum Difference is K %s\n", name), func(t *testing.T) {
			result := partitionArray(test.nums, test.k)
			if result != test.answer {
				t.Fatalf("Expected is %d but got %d\n", test.answer, result)
			}
		})
	}
}
