package binarysearch

import (
	"fmt"
	"testing"
)

func searchMatrix(matrix [][]int, target int) bool {
	// We are given a matrix m x n where m is the height
	// of the matrix and n is the length of the matrix.

	// The idea is to search first at which height our target could be.
	// Then, if we find that height, we could search for the target horizontally.

	m := len(matrix)
	n := len(matrix[0])

	// We can do sanity check earlier
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	// The vertical index that will be used to find horizontally.
	vertIdx := 0

	// Finding vertically
	left, right, mid := 0, m-1, 0

	for left < right {
		mid = (right + left) / 2

		if target >= matrix[mid][0] {
			if target < matrix[mid+1][0] {
				vertIdx = mid
				break
			} else {
				left = mid + 1
				vertIdx = left
			}
		} else if target <= matrix[mid][0] {
			if target > matrix[mid-1][0] {
				vertIdx = mid - 1
				break
			} else {
				right = mid - 1
				vertIdx = right
			}
		}
	}

	// Finding horizontally
	left, right, mid = 0, n-1, 0

	for left <= right {
		mid := (right + left) / 2

		if target == matrix[vertIdx][mid] {
			return true
		} else if target > matrix[vertIdx][mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func TestSearchMatrix(t *testing.T) {
	testcases := []struct {
		matrix [][]int
		target int
		result bool
	}{
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			result: true,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 30,
			result: true,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 11,
			result: true,
		},
		{
			matrix: [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}, {13, 14, 17}, {18, 19, 20}},
			target: 20,
			result: true,
		},
		{
			matrix: [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}, {13, 14, 17}, {18, 19, 20}},
			target: 2,
			result: true,
		},
		{
			matrix: [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}, {13, 14, 17}, {18, 19, 20}},
			target: 5,
			result: true,
		},
		{
			matrix: [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}, {13, 14, 17}, {18, 19, 20}},
			target: 21,
			result: false,
		},
		{
			matrix: [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}, {13, 14, 17}, {18, 19, 20}},
			target: 12,
			result: false,
		},
		{
			matrix: [][]int{{1}},
			target: 12,
			result: false,
		},
		{
			matrix: [][]int{{1}, {5}},
			target: 0,
			result: false,
		},
		{
			matrix: [][]int{{1}, {5}},
			target: 5,
			result: true,
		},
		{
			matrix: [][]int{{1}, {5}},
			target: 6,
			result: false,
		},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Running searchMatrix testcase %d", idx+1), func(t *testing.T) {
			if result := searchMatrix(tc.matrix, tc.target); result != tc.result {
				t.Fatalf("Expected to be %t but got %t", tc.result, result)
			}
		})
	}
}
