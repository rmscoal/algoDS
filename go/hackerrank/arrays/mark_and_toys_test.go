package arrays

import (
	"fmt"
	"testing"
)

func TestMarkAndToys(t *testing.T) {
	tests := map[string]struct {
		prices []int32
		k      int32
		answer int32
	}{
		"test_1": {
			prices: []int32{1, 12, 5, 111, 200, 1000, 10},
			k:      50,
			answer: 4,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Case %s\n", name), func(t *testing.T) {
			result := maximumToys(test.prices, test.k)
			if result != test.answer {
				t.Fatalf("Expected %d but got %d\n", test.answer, result)
			}
		})
	}
}
