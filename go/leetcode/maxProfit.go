package leetcode

import (
	"math"
	"sort"
)

func (MP *MaxProfit) Solver(prices []int) int {
	// Variables
	// max_profit holds the max profit
	var max_profit int = 0
	// smallest holds the smallest stock prices in the array
	var smallest float64 = math.Pow10(4)
	// smallestIndex holds the index of the smallest stock
	// prices in the array.
	var smallestIndex int = 0

	// Find the smallest stock prices in the array.
	// Then keeps the index to the smallestIndex
	for i := 0; i < len(prices); i++ {
		if prices[i] < int(smallest) {
			smallest = float64(prices[i])
			smallestIndex = i
		}
	}

	// Find the biggest stock prices after the smallest stock
	// prices in the array. For doing so, sort the array after
	// the smallest prices. Then take the last element of the array
	// assuming that it is the biggest. Getting these two values,
	// calculate the max profit as biggest - smallest.
	sort.Ints(prices[smallestIndex+1:])
	max_profit = prices[smallestIndex:][len(prices[smallestIndex:])-1] - int(smallest)

	// For this case, if the length of the prices is bigger than 2
	// and the smallestIndex is located at 1, return the max_profit
	// immediately such that there no use for looking from the array
	// before the previous of the smallestIndex.
	// Ex: [2,1,3]
	if smallestIndex == 1 {
		if len(prices) > 2 {
			return max_profit
		}
	}

	// Iterate through the array before the smallestIndex to look
	// for a bigger max_profit. Two pointers method.
	// Ex: [2,4,5,6,1,2] => max_profit = 2 instead of 1.
	for index, val1 := range prices[:smallestIndex] {
		for _, val2 := range prices[index:] {
			if val2-val1 > max_profit {
				max_profit = val2 - val1
			} else {
				continue
			}
		}
	}

	return max_profit
}

func (MP *MaxProfit) FastSolver(prices []int) int {
	max_profit, buyPrice := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		sellPrice := prices[i]
		if max_profit < sellPrice-buyPrice {
			max_profit = sellPrice - buyPrice
		}
		if sellPrice < buyPrice {
			buyPrice = sellPrice
		}
	}

	return max_profit
}
