package stack

import (
	"fmt"
	"sort"
	"testing"
)

func carFleet_Slow(target int, positions []int, speeds []int) int {
	// Position and speed pair
	pair := make([][2]int, len(positions))
	for i := 0; i < len(positions); i++ {
		pair[i] = [2]int{positions[i], speeds[i]}
	}

	sort.SliceStable(pair, func(i, j int) bool {
		return pair[i][0] > pair[j][0]
	})

	stack := make([]float64, 0, len(positions))

	for i := range pair {
		timeTaken := float64(target-pair[i][0]) / float64(pair[i][1])

		if len(stack) == 0 || timeTaken > stack[len(stack)-1] {
			stack = append(stack, timeTaken)
		}
	}

	return len(stack)
}

func carFleet(target int, position []int, speed []int) int {
	fleets := 0

	// You are trying to actually sort the car's position
	// but using an array for it and placing the speed at
	// the position. This is possible because the position
	// is unique.
	lane := make([]int, target+1)
	for i, v := range position {
		lane[v] = speed[i]
	}

	var timeTakenForCarInFront float64
	for i := target; i >= 0; i-- {
		// If there are no cars at this position
		// then we just continue, there is nothing
		// to calculate here
		if lane[i] == 0 {
			continue
		}
		timeTakenCurrentCar := float64(target-i) / float64(lane[i])
		if timeTakenCurrentCar > timeTakenForCarInFront {
			timeTakenForCarInFront = timeTakenCurrentCar
			fleets++
		}
	}

	return fleets
}

func TestCarFleet(t *testing.T) {
	type testcase struct {
		target    int
		positions []int
		speeds    []int
		result    int
	}
	testcases := []testcase{
		{
			target:    12,
			positions: []int{10, 8, 0, 5, 3},
			speeds:    []int{2, 4, 1, 1, 3},
			result:    3,
		},
		{
			target:    10,
			positions: []int{8, 6},
			speeds:    []int{2, 6},
			result:    1,
		},
		{
			target:    10,
			positions: []int{3},
			speeds:    []int{3},
			result:    1,
		},
		{
			target:    100,
			positions: []int{0, 2, 4},
			speeds:    []int{4, 2, 1},
			result:    1,
		},
		{
			target:    50,
			positions: []int{1, 2, 3},
			speeds:    []int{1, 2, 3},
			result:    3,
		},
		{
			target:    10,
			positions: []int{8, 3, 7, 4, 6, 5},
			speeds:    []int{4, 4, 4, 4, 4, 4},
			result:    6,
		},
		{
			target:    15,
			positions: []int{14, 13},
			speeds:    []int{1, 1},
			result:    2,
		},
		{
			target:    15,
			positions: []int{14, 13},
			speeds:    []int{1, 2},
			result:    1,
		},
		{
			target:    15,
			positions: []int{14, 13},
			speeds:    []int{1, 2},
			result:    1,
		},
		{
			target:    31,
			positions: []int{5, 26, 18, 25, 29, 21, 22, 12, 19, 6},
			speeds:    []int{7, 6, 6, 4, 3, 4, 9, 7, 6, 4},
			result:    6,
		},
		{
			target:    10,
			positions: []int{8, 6},
			speeds:    []int{2, 2},
			result:    2,
		},
		{
			target:    10,
			positions: []int{8, 6, 5},
			speeds:    []int{2, 2, 4},
			result:    2,
		},
	}

	for idx, tc := range testcases {
		t.Run(fmt.Sprintf("Running car fleet slow test case %d", idx+1), func(t *testing.T) {
			if result := carFleet_Slow(tc.target, tc.positions, tc.speeds); result != tc.result {
				t.Fatalf("Expected is %d but got %d", tc.result, result)
			}
		})
		t.Run(fmt.Sprintf("Running car fleet test case %d", idx+1), func(t *testing.T) {
			if result := carFleet(tc.target, tc.positions, tc.speeds); result != tc.result {
				t.Fatalf("Expected is %d but got %d", tc.result, result)
			}
		})
	}
}
