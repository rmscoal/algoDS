package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func containsDuplicate(nums []int) bool {
	maps := map[int]int{}

	for _, v := range nums {
		if _, ok := maps[v]; ok {
			return true
		} else {
			maps[v] = 1
		}
	}

	return false
}

func TestContainsDuplicate_1(t *testing.T) {
	result := containsDuplicate([]int{1, 2, 3, 4})
	assert.False(t, result)
}

func TestContainsDuplicate_2(t *testing.T) {
	result := containsDuplicate([]int{1, 2, 1})
	assert.True(t, result)
}
