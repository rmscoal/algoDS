package arrays

import "testing"

// https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	// Check for row duplicate
	hash := make(map[byte]bool, 0)
	for row := 0; row < 9; row++ {
		// How many filled box in a row.
		count := 0

		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				hash[board[row][col]] = true
				count++
			}
		}

		// Checks whether the length is
		// the same as count.
		if len(hash) != count {
			return false
		}

		// Reset our hash
		hash = make(map[byte]bool, 0)
	}

	// Check for column duplicate
	hash = make(map[byte]bool, 0)
	for col := 0; col < 9; col++ {
		// How many filled box in a row.
		count := 0

		for row := 0; row < 9; row++ {
			if board[row][col] != '.' {
				hash[board[row][col]] = true
				count++
			}
		}

		// Checks whether the length is
		// the same as count.
		if len(hash) != count {
			return false
		}

		// Reset our hash
		hash = make(map[byte]bool, 0)
	}

	// Check for 3x3
	for box := 1; box <= 9; box++ {
		count := 0
		hash = make(map[byte]bool)
		iMup := ((box - 1) / 3) * 3
		jMup := ((box - 1) % 3) * 3
		for i := 0 + iMup; i < 3+iMup; i++ {
			for j := 0 + jMup; j < 3+jMup; j++ {
				if board[i][j] != '.' {
					hash[board[i][j]] = true
					count++
				}
			}
		}

		// Checks whether the length is
		// the same as count.
		if len(hash) != count {
			return false
		}
	}

	return true
}

func TestIsValidSudoku(t *testing.T) {
	if shouldBeTrue := isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}); !shouldBeTrue {
		t.Fatalf("Should be true")
	}
}
