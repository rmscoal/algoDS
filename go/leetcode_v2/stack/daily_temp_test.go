package stack

import (
	"fmt"
	"testing"
)

func dailyTemperatures_TimeLimit_Exceeded(temperatures []int) []int {
	// Resulting array
	result := make([]int, len(temperatures))
	// We could store the index of the temperatures that
	// the next day are lesser than its current day
	stack := map[int]bool{}

	// We only need to analyze until second last element
	// of the temperatures since, the last will always
	// be 0 in the resulting array.
	for i := 0; i < len(temperatures); i++ {
		// Off loads whats in the stack
		for idx := range stack {
			if temperatures[idx] < temperatures[i] {
				result[idx] = i - idx
				delete(stack, idx)
			}
		}

		// Do not need to check last element
		if i == len(temperatures)-1 {
			continue
		}

		// Then analyze current day with regards of tomorrow
		if temperatures[i] >= temperatures[i+1] {
			stack[i] = true
		} else {
			result[i] = 1
		}
	}

	return result
}

func dailyTemperatures(temperatures []int) []int {
	// IDEA:
	// * Start the stack with having the first index
	// of temperatures.
	// * Then iterate through temperatures from index = 1
	// till n-1.
	// * Do a while loop for the stack is not empty.
	// * Retrieve the top element from the stack (that
	// is the last element of the stack array) and checks
	// if the element of this particular temperature is
	// lesser than the current temperature.
	// 		* If it does, calculate the difference in indexes
	// 		then pop from the stack. Repeat.
	// 		*If not, break the loop
	// Appends the current index to the stack.
	//
	// The key here is that, IT IS GUARANTEED THAT THE
	// TEMPERATURE AT THE TOP OF THE STACK IS THE LOWEST
	// IN THE STACK.
	result := make([]int, len(temperatures))
	stack := []int{0}

	for i := 1; i < len(temperatures); i++ {
		for len(stack) != 0 {
			topIdx := stack[len(stack)-1]
			if temperatures[topIdx] < temperatures[i] {
				result[topIdx] = i - topIdx
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	return result
}

func Test_DailyTemp_All(t *testing.T) {
	testcases := []struct {
		temperatures []int
		result       []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			result:       []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{30, 60, 90},
			result:       []int{1, 1, 0},
		},
		{
			temperatures: []int{30, 40, 50, 60},
			result:       []int{1, 1, 1, 0},
		},
		{
			temperatures: []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70},
			result:       []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{77, 77, 77, 77, 77, 41, 77, 41, 41, 77},
			result:       []int{0, 0, 0, 0, 0, 1, 0, 2, 1, 0},
		},
	}

	for idx, tc := range testcases {
		t.Run("dailyTemperatures_TimeLimit_Exceeded", func(t *testing.T) {
			t.Run(fmt.Sprintf("Running test case %d", idx), func(t *testing.T) {
				result := dailyTemperatures_TimeLimit_Exceeded(tc.temperatures)
				for idx, v := range result {
					if v != tc.result[idx] {
						t.Fatalf("expected result at index %d is %d\n", idx, tc.result[idx])
					}
				}
			})
		})
		t.Run("dailyTemperatures", func(t *testing.T) {
			t.Run(fmt.Sprintf("Running test case %d", idx), func(t *testing.T) {
				result := dailyTemperatures(tc.temperatures)
				for idx, v := range result {
					if v != tc.result[idx] {
						t.Fatalf("expected result at index %d is %d\n", idx, tc.result[idx])
					}
				}
			})
		})
	}
}
