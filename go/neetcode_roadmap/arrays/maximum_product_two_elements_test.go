package arrays

import (
	"fmt"
	"testing"
)

func TestMaxProducts(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		answer int
	}{
		"test_1": {
			nums:   []int{3, 4, 5, 2},
			answer: 12,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Maximum Product of Two Elements in Array %s\n", name), func(t *testing.T) {
			result := maxProduct(test.nums)
			if result != test.answer {
				t.Fatalf("Expected %d but got result %d\n", test.answer, result)
			}

			result = maxProductV2(test.nums)
			if result != test.answer {
				t.Fatalf("Expected %d but got result %d\n", test.answer, result)
			}
		})
	}
}
