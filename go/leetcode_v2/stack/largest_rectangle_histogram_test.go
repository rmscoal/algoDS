package stack

import (
	"fmt"
	"testing"
)

func largestRectangleArea(heights []int) int {
	// Here we want to have the idea that the heights will just be multiplied
	// witht the widths that is possible for expansion. So the idea is, we calculate
	// how many widths in front of each heights that can form rectangle + how many
	// width are behind each heights that can form a rectangle too. Then for each
	// heights and width pair, we find the maximum multiplication result.
	widths := make([]int, len(heights))
	for i := 0; i < len(widths); i++ {
		widths[i] = 1
	}

	stack := []int{0}
	for i := 1; i < len(heights); i++ {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if heights[i] < heights[top] {
				stack = stack[:len(stack)-1]
				widths[top] += i - top - 1
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	for _, top := range stack {
		widths[top] += len(heights) - top - 1
	}

	stack = []int{len(heights) - 1}
	for i := len(heights) - 2; i >= 0; i-- {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if heights[i] < heights[top] {
				stack = stack[:len(stack)-1]
				widths[top] += top - i - 1
			} else {
				break
			}
		}
		stack = append(stack, i)
	}

	for _, top := range stack {
		widths[top] += top
	}

	result := 0
	for i := 0; i < len(heights); i++ {
		result = max(result, heights[i]*widths[i])
	}

	return result
}

func largestRectangleArea_FROM_LEETCODE_BEST_SOLUTION(heights []int) int {
	// i think we will have to keep a stack here which is increasing order.. so if we get an element which is greater than
	// the top, then we will append the value to the stack.. and if we get an element which is smaller than the top , then
	// top is sandwhiched between 2 small values. and we can basically pop the value and then find the diff between top and
	// current and multiply by the element we just popped out.
	// similarly we can pop out as long as the next is smaller than top .
	// finally we will have values in ascending order.. since they were not popped means that we didnt receive any value
	// that was smaller than that.. so we can take the difference between the n and the previous element in the stack
	// also we can add -1 to stack first

	stack := []int{-1}
	n := 1
	ans := 0
	for i, val := range heights {
		for n > 1 && heights[stack[n-1]] > val { // the top of the stack is greater than the next element // and i am considering n>1 , because then only the top is sandwhiched between 2 elments
			top := stack[n-1]
			n--
			stack = stack[:n]
			ans = max(ans, heights[top]*(i-stack[n-1]-1))
		}
		stack = append(stack, i)
		n++
	}
	for len(stack) > 1 { // pop out all the elements except the bottom element which was -1
		top := stack[n-1]
		n--
		stack = stack[:n]
		ans = max(ans, heights[top]*(len(heights)-stack[n-1]-1))
	}

	return ans
}

func TestLargestRectangleArea(t *testing.T) {
	type testcase struct {
		heights []int
		result  int
	}

	testcases := []testcase{
		{
			heights: []int{1, 2, 3, 4},
			result:  6,
		},
		{
			heights: []int{2, 1, 2},
			result:  3,
		},
		{
			heights: []int{2, 0, 2},
			result:  2,
		},
		{
			heights: []int{2, 1, 5, 6, 2, 3},
			result:  10,
		},
		{
			heights: []int{2, 4},
			result:  4,
		},
		{
			heights: []int{4, 3, 2, 1},
			result:  6,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Running largestRectangleArea for test case %d", i+1), func(t *testing.T) {
			if result := largestRectangleArea(tc.heights); result != tc.result {
				t.Fatalf("Expected %d but got %d", tc.result, result)
			}
		})
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("Running largestRectangleArea_Faster for test case %d", i+1), func(t *testing.T) {
			if result := largestRectangleArea_FROM_LEETCODE_BEST_SOLUTION(tc.heights); result != tc.result {
				t.Fatalf("Expected %d but got %d", tc.result, result)
			}
		})
	}
}
