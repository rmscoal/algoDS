// https://www.geeksforgeeks.org/minimize-maximum-adjacent-difference-in-a-path-from-top-left-to-bottom-right/

package bfs

import (
	"math"
	"sort"
)

func helper(matrix [][]int, visited [][]int, max, i, j int, maxArr *[]int) {
	if i < len(matrix) && i >= 0 && j < len(matrix[0]) && j >= 0 {
		for _, coor := range visited {
			if i == coor[0] && j == coor[1] {
				return
			}
		}
		prevCoor := visited[len(visited)-1]
		visited = append(visited, []int{i, j})
		currentValue := matrix[i][j]
		prevValue := matrix[prevCoor[0]][prevCoor[1]]
		diff := math.Abs(float64(prevValue) - float64(currentValue))
		if max < int(diff) {
			max = int(diff)
		}
		if i == len(matrix)-1 && j == len(matrix[i])-1 {
			*maxArr = append(*maxArr, max)
		} else {
			helper(matrix, visited, max, i+1, j, maxArr)
			helper(matrix, visited, max, i-1, j, maxArr)
			helper(matrix, visited, max, i, j+1, maxArr)
			helper(matrix, visited, max, i, j-1, maxArr)
		}
	}
}

func minAdjHeightDiff(matrix [][]int) int {
	// Max variable
	max := -1
	maxArr := make([]int, 0)

	helper(matrix, [][]int{{0, 0}}, max, 0, 1, &maxArr)
	helper(matrix, [][]int{{0, 0}}, max, 1, 0, &maxArr)

	sort.Ints(maxArr)

	return maxArr[0]
}
