package arrays

import (
	"math"
	"sort"
)

func coinChange_MemoryLimit(coins []int, amount int) int {
	// BFS
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)

	var startIndex int

	for i := len(coins) - 1; i > -1; i-- {
		if coins[i] == amount {
			return 1
		} else if coins[i] < amount {
			startIndex = i
			break
		} else {
			startIndex = -1
		}
	}

	if startIndex == -1 {
		return -1
	}

	if amount%coins[startIndex] == 0 {
		return amount / coins[startIndex]
	}

	level := 1
	queue := []int{coins[startIndex]}

	for len(queue) > 0 {
		size := len(queue)
		level++

		for size > 0 {
			sum := queue[0]
			queue = queue[1:]
			size--

			for x := 0; x <= startIndex; x++ {
				if sum+coins[x] == amount {
					return level
				} else if sum+coins[x] < amount {
					queue = append(queue, sum+coins[x])
				}
			}
		}
	}

	return -1
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange_Memoization(coins []int, amount int) int {
	memo := make([]int, amount+1)

	for i := 1; i < len(memo); i++ {
		memo[i] = math.MaxInt32
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				memo[i] = findMin(memo[i-coins[j]]+1, memo[i])
			}
		}
	}

	if memo[amount] == math.MaxInt32 {
		return -1
	}

	return memo[amount]
}
