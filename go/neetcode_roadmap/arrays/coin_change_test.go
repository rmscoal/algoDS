package arrays

import (
	"fmt"
	"testing"
)

func TestCoinChange_MemoryLimit(t *testing.T) {
	tests := map[string]struct {
		coins          []int
		amount, answer int
	}{
		"test_1": {
			coins:  []int{1, 2, 5},
			amount: 11,
			answer: 3,
		},
		"test_2": {
			coins:  []int{2, 5},
			amount: 11,
			answer: 4,
		},
		"test_3": {
			coins:  []int{1},
			amount: 11,
			answer: 11,
		},
		"test_4": {
			coins:  []int{1},
			amount: 0,
			answer: 0,
		},
		"test_5": {
			coins:  []int{2},
			amount: 3,
			answer: -1,
		},
		"test_6": {
			coins:  []int{1},
			amount: 1,
			answer: 1,
		},
		"test_7": {
			coins:  []int{1, 2},
			amount: 2,
			answer: 1,
		},
		// Memory Limit Here
		"test_8": {
			coins:  []int{1, 2, 5},
			amount: 100,
			answer: 20,
		},
		// Memory Limit Here
		"test_9": {
			coins:  []int{186, 419, 83, 408},
			amount: 6249,
			answer: 20,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Coin Change %s", name), func(t *testing.T) {
			if result := coinChange_MemoryLimit(test.coins, test.amount); result != test.answer {
				t.Fatalf("Expected %s to be %d but got %d\n", name, test.answer, result)
			}
		})
	}
}

func TestCoinChange_Memoization(t *testing.T) {
	tests := map[string]struct {
		coins          []int
		amount, answer int
	}{
		"test_1": {
			coins:  []int{1, 2, 5},
			amount: 11,
			answer: 3,
		},
		"test_2": {
			coins:  []int{2, 5},
			amount: 11,
			answer: 4,
		},
		"test_3": {
			coins:  []int{1},
			amount: 11,
			answer: 11,
		},
		"test_4": {
			coins:  []int{1},
			amount: 0,
			answer: 0,
		},
		"test_5": {
			coins:  []int{2},
			amount: 3,
			answer: -1,
		},
		"test_6": {
			coins:  []int{1},
			amount: 1,
			answer: 1,
		},
		"test_7": {
			coins:  []int{1, 2},
			amount: 2,
			answer: 1,
		},
		// Memory Limit Here
		"test_8": {
			coins:  []int{1, 2, 5},
			amount: 100,
			answer: 20,
		},
		// Memory Limit Here
		"test_9": {
			coins:  []int{186, 419, 83, 408},
			amount: 6249,
			answer: 20,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Coin Change %s", name), func(t *testing.T) {
			if result := coinChange_Memoization(test.coins, test.amount); result != test.answer {
				t.Fatalf("Expected %s to be %d but got %d\n", name, test.answer, result)
			}
		})
	}
}
