package leetcode

import "strings"

func (c *CollectGarbage) FastSolver(garbage []string, travel []int) int {
	// Returned variable:
	var res int = 0

	// Holds the house max index for each
	// garbage type. P stands for paper,
	// M stands for metal and G stands for
	// glass.
	var maxHouseIndexP int = 0
	var maxHouseIndexG int = 0
	var maxHouseIndexM int = 0

	// For loop to find the house max index
	// for each garbage type and adds the time
	// it takes to load garbage in each house
	// by counting the length of each string in
	// the param garbage.
	for i := 0; i < len(garbage); i++ {
		if strings.Contains(garbage[i], "M") {
			maxHouseIndexM = i
		}

		if strings.Contains(garbage[i], "P") {
			maxHouseIndexP = i
		}

		if strings.Contains(garbage[i], "G") {
			maxHouseIndexG = i
		}

		res += len(garbage[i])
	}

	// This for loop is responsible to add
	// the distance of travel to its houses
	// depending on the max house index of
	// each garbage type.
	for i := 0; i < len(travel); i++ {
		if i < maxHouseIndexG {
			res += travel[i]
		}

		if i < maxHouseIndexP {
			res += travel[i]
		}

		if i < maxHouseIndexM {
			res += travel[i]
		}
	}

	return res
}
