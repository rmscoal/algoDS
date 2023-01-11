package leetcode

import "math"

/*
	Problem:
	There is a city composed of n x n blocks, where each block
	contains a single building shaped like a vertical square prism.
	You are given a 0-indexed n x n integer matrix grid where grid[r][c]
	represents the height of the building located in the block at row r
	and column c.

	A city's skyline is the the outer contour formed by all the building
	when viewing the side of the city from a distance. The skyline from
	each cardinal direction north, east, south, and west may be different.

	We are allowed to increase the height of any number of buildings by
	any amount (the amount can be different per building). The height of
	a 0-height building can also be increased. However, increasing the
	height of a building should not affect the city's skyline from any
	cardinal direction.

	Return the maximum total sum that the height of the buildings can
	be increased by without changing the city's skyline from any cardinal
	direction.

	Constraints:
	- n == grid.length
	- n == grid[r].length
	- 2 <= n <= 50
	- 0 <= grid[r][c] <= 100
*/

// Helper function to find the greatest horizontal and vertical integer
// in the grid according to the given index on the grid.
func findIndexGreatestHeight(grid [][]int, gridSize, i, j int) (int, int) {
	// Returned variable
	maxHorizontal, maxVertical := grid[i][j], grid[i][j]

	// Finds the max horizontal grid
	for k := 0; k < gridSize; k++ {
		if grid[k][j] > maxHorizontal {
			maxHorizontal = grid[k][j]
		}
	}

	// Finds the max vertical grid
	for k := 0; k < gridSize; k++ {
		if grid[i][k] > maxVertical {
			maxVertical = grid[i][k]
		}
	}

	return maxHorizontal, maxVertical
}

func (s *KeepCitySkyline) Solver(grid [][]int) int {
	// Returned variable
	increaseInTotalHeight := 0

	// Grid's size
	gridSize := len(grid)

	/*
		How to solve:
		Firstly, iterate through all the elements in the grid. Then, find its
		corresponding horizontal and vertical array's greatest height. From here,
		find the minimum of those two values then increment the returned variable
		`increaseInTotalHeight' with the difference of the current index and the
		minimum value found.
	*/
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			maxHorizontal, maxVertical := findIndexGreatestHeight(grid, gridSize, i, j)
			minOfMaxs := int(math.Min(float64(maxHorizontal), float64(maxVertical)))
			if grid[i][j] < minOfMaxs {
				increaseInTotalHeight += minOfMaxs - grid[i][j]
			}
		}
	}

	return increaseInTotalHeight
}
