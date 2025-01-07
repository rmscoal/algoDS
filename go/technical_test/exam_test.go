package technicaltest

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// reverseArray reverses the elements of the input slice.
func reverseArray(arr []int) []int {
	// Create a new slice to hold the reversed elements
	reversed := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-1-i]
	}
	return reversed
}

func calculation(arrInput1 []int, addition int) []int {
	var num1 int
	for i := len(arrInput1) - 1; i >= 0; i-- {
		multiplier := math.Pow(10, float64(len(arrInput1)-1-i))
		num1 += arrInput1[i] * (int(multiplier))
	}

	add := num1 + addition
	if add == 0 {
		return []int{add}
	}

	result := make([]int, 0)
	for add != 0 {
		unit := add - ((add / 10) * 10)
		result = append(result, unit)
		add /= 10
	}

	// Reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func TestCalculation(t *testing.T) {
	tests := []struct {
		arrInput1 []int
		addition  int
		expect    []int
	}{
		{
			arrInput1: []int{2, 0, 5},
			addition:  3,
			expect:    []int{2, 0, 8},
		},
		{
			arrInput1: []int{2, 0, 5},
			addition:  5,
			expect:    []int{2, 1, 0},
		},
		{
			arrInput1: []int{1, 0, 0, 9},
			addition:  11,
			expect:    []int{1, 0, 2, 0},
		},
		{
			arrInput1: []int{1, 9, 9, 9},
			addition:  1,
			expect:    []int{2, 0, 0, 0},
		},
		{
			arrInput1: []int{9, 9, 9, 9},
			addition:  1,
			expect:    []int{1, 0, 0, 0, 0},
		},
		{
			arrInput1: []int{0},
			addition:  0,
			expect:    []int{0},
		},
		{
			arrInput1: []int{1, 2},
			addition:  -1,
			expect:    []int{1, 1},
		},
		{
			arrInput1: []int{1, 0},
			addition:  -1,
			expect:    []int{9},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expect, calculation(tt.arrInput1, tt.addition))
	}
}

func twoSum(inputArr []int, targetInt int) []int {
	mapper := map[int]int{}

	for idx := 0; idx < len(inputArr); idx++ {
		partnerIdx, found := mapper[inputArr[idx]]
		if found {
			return []int{partnerIdx, idx}
		}
		mapper[targetInt-inputArr[idx]] = idx
	}

	return []int{}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		inputArr []int
		target   int
		expect   []int
	}{
		{
			inputArr: []int{2, 7, 11, 15},
			target:   9,
			expect:   []int{0, 1},
		},
		{
			inputArr: []int{2, 5, -3, 10},
			target:   -1,
			expect:   []int{0, 2},
		},
		{
			inputArr: []int{2, 2},
			target:   4,
			expect:   []int{0, 1},
		},
		{
			inputArr: []int{8, 4, 1},
			target:   5,
			expect:   []int{1, 2},
		},
		{
			inputArr: []int{4392, 9193, 1, 324156, 4567398413, 7843, 1921, 4293, -43278, -43298103},
			target:   4294,
			expect:   []int{2, 7},
		},
		{
			inputArr: []int{4392, 9193, 1, 324156, 4567398413, 7843, 1921, 4293, -43278, -43298103},
			target:   9194,
			expect:   []int{1, 2},
		},
		{
			inputArr: []int{4392, 9193, 1, 324156, 4567398413, 7843, 1921, 4293, -43278, -43298103},
			target:   -43277,
			expect:   []int{2, 8},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expect, twoSum(tt.inputArr, tt.target))
	}
}
