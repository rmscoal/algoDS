package leetcode

import (
	"sort"
)

func (s *MinMovesToSeat) Solver(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	total := 0
	for i := 0; i < len(seats); i++ {
		val := seats[i] - students[i]
		if val < 0 {
			total += -1 * val
		} else {
			total += val
		}
	}

	return total
}
