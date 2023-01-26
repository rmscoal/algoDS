package benchmark

import (
	"competitive/leetcode"
	"fmt"
	"testing"
)

var stayingInGridTestCases = []struct {
	n        int
	startPos []int
	str      string
}{
	{n: 3, startPos: []int{0, 1}, str: "RRDDLU"},
	{n: 2, startPos: []int{1, 1}, str: "LURD"},
	{n: 1, startPos: []int{0, 0}, str: "LRUD"},
	{n: 100, startPos: []int{60, 52}, str: "LLLDURDLRUURRLLDLDLRUUDDDLRLRUUDDLRRLRRLDUDLLLRDLRUDL"},
}

// Use -bench=. to run all benchmark
//
// To avoid executing any test functions in the test files,
// pass a regular expression to the -run flag
//
// Preferred commands:
// go test -bench=StayingInGridSolver -count 5 -benchtime=100x -benchmem
// go test -bench=StayingInGridSolver -count 5 -benchtime=10s -benchmem

func BenchmarkStayingInGridSolver(b *testing.B) {
	CG := &leetcode.StayingInGrid{}
	for i, v := range stayingInGridTestCases {
		b.Run(fmt.Sprintf("test case no-%d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CG.Solver(v.n, v.startPos, v.str)
			}
		})
	}
}
