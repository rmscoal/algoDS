package string

import (
	"fmt"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		answer string
	}{
		"test_1": {
			nums:   []int{10, 2},
			answer: "210",
		},
		"test_2": {
			nums:   []int{3, 30, 34, 5, 9},
			answer: "9534330",
		},
		"test_3": {
			nums:   []int{432, 43243},
			answer: "43243432",
		},
		"test_4": {
			nums:   []int{3, 30},
			answer: "330",
		},
		"test_5": {
			nums:   []int{1113, 111311},
			answer: "1113111311",
		},
		"test_6": {
			nums:   []int{0, 0},
			answer: "0",
		},
		"test_7": {
			nums:   []int{0, 10},
			answer: "100",
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Largest Number %s", name), func(t *testing.T) {
			result := largestNumber(test.nums)
			if result != test.answer {
				t.Fatalf("Expected is %s but got %s\n", test.answer, result)
			}
		})
	}
}
