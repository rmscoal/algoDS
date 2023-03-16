package leetcode

func (u *UniquePathThree) Solver(grid [][]int) int {
	remaining := 0
	startX := 0
	startY := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				remaining += 1
			} else if grid[i][j] == 1 {
				startX = i
				startY = j
			}
		}
	}

	return u.helper(grid, remaining, startX, startY)
}

func (u *UniquePathThree) helper(grid [][]int, remaining, i, j int) int {
	result := 0

	// grid[i][j] = -1 means that either box is visited
	// 								or obstacle
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[i]) || grid[i][j] == -1 {
		return 0
	} else if grid[i][j] == 0 || grid[i][j] == 1 {
		grid[i][j] = -1
		result = u.helper(grid, remaining-1, i-1, j) + u.helper(grid, remaining-1, i+1, j) + u.helper(grid, remaining-1, i, j-1) + u.helper(grid, remaining-1, i, j+1)
		grid[i][j] = 0 // return to its normal so other recursived function can use this path
	} else if grid[i][j] == 2 && remaining < 0 {
		return 1
	}

	return result
}
