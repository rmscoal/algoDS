package greedy

func maxIncreaseKeepingSkyline(grid [][]int) int {
	increasedBy := 0
	/*
		Cache the max height of horizontal and vertical skylines

		Example:
		cache := [
			[8, 7, 9, 3] -> horizontal
			[9, 4, 8, 7] -> Vertical
		]

		If we land on arr[0,0]
		We check cache[0][0] and check cache[1][0]

		If we land on arr[2,3]
		We check cache[0][2] and check cache[1][3]
	*/
	cache := [2][]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < len(grid); j++ {
			cache[i] = append(cache[i], -1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if cache[0][i] == -1 {
				for k := 0; k < len(grid); k++ {
					if cache[0][i] < grid[i][k] {
						cache[0][i] = grid[i][k]
					}
				}
			}

			if cache[1][j] == -1 {
				for k := 0; k < len(grid); k++ {
					if cache[1][j] < grid[k][j] {
						cache[1][j] = grid[k][j]
					}
				}
			}

			if cache[0][i] < cache[1][j] {
				increasedBy += cache[0][i] - grid[i][j]
			} else {
				increasedBy += cache[1][j] - grid[i][j]
			}
		}
	}

	return increasedBy
}
