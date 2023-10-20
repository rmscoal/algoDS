package benchmark

import (
	"fmt"
	"testing"

	"algoDS/go/leetcode"
)

var addTwoNumbersTestCase = []struct {
	l1 *leetcode.ListNode
	l2 *leetcode.ListNode
}{
	{l1: leetcode.GenerateSingleLinkedListFromSlice([]int{2, 4, 3}), l2: leetcode.GenerateSingleLinkedListFromSlice([]int{5, 6, 4})},
	{l1: leetcode.GenerateSingleLinkedListFromSlice([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}), l2: leetcode.GenerateSingleLinkedListFromSlice([]int{5, 6, 4})},
	{l1: leetcode.GenerateSingleLinkedListFromSlice([]int{9, 9, 9, 9, 9, 9, 9}), l2: leetcode.GenerateSingleLinkedListFromSlice([]int{9, 9, 9, 9})},
}

// Use -bench=. to run all benchmark
//
// To avoid executing any test functions in the test files,
// pass a regular expression to the -run flag
//
// Preferred commands:
// go test -bench=StayingInGridSolver -count 5 -benchtime=100x -benchmem
// go test -bench=AddTwoNumbers -count 5 -benchtime=10s -benchmem

func BenchmarkAddTwoNumbers(b *testing.B) {
	CG := &leetcode.AddTwoNumbers{}
	for i, v := range addTwoNumbersTestCase {
		b.Run(fmt.Sprintf("test case no-%d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CG.Solver(v.l1, v.l2)
			}
		})
	}
}
