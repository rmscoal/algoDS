package benchmark

import (
	"competitive/leetcode"
	"fmt"
	"testing"
)

// Benchmarking with various inputs
var collectGarbageTestCases = []struct {
	garbage []string
	travel  []int
}{
	{garbage: []string{"G", "P", "GP", "GG"}, travel: []int{2, 4, 3}},
	{garbage: []string{"G", "M", "P"}, travel: []int{3, 10}},
}

// Use -bench=. to run all benchmark
//
// To avoid executing any test functions in the test files,
// pass a regular expression to the -run flag
//
// Preferred commands:
// go test -bench=CollectGarbageFastSolver -count 5 -benchtime=100x -benchmem -run=^#
// go test -bench=CollectGarbageFastSolver -count 5 -benchtime=10s -benchmem -run=^#

func BenchmarkCollectGarbageFastSolver(b *testing.B) {
	CG := &leetcode.CollectGarbage{}
	for _, v := range collectGarbageTestCases {
		b.Run(fmt.Sprintf("garbage_%+v and travel_%+v", v.garbage, v.travel), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CG.FastSolver(v.garbage, v.travel)
			}
		})
	}
}
