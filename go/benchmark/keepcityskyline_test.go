package benchmark

import (
	"competitive/leetcode"
	"fmt"
	"testing"
)

// Benchmarking with various inputs
var keepCitySkylineTestCases = []struct {
	input [][]int
}{
	{input: [][]int{{3, 0, 8, 4, 3, 0, 8, 4, 3, 0, 8, 4},
		{2, 4, 5, 7, 2, 4, 5, 7, 2, 4, 5, 7}, {9, 2, 6, 3, 9, 2, 6, 3, 9, 2, 6, 3},
		{0, 3, 1, 0, 0, 3, 1, 0, 0, 3, 1, 0}}},
	{input: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
}

// Use -bench=. to run all benchmark
//
// To avoid executing any test functions in the test files,
// pass a regular expression to the -run flag
//
// Preferred commands:
// go test -bench=KeepCitySkylineSolver -count 5 -benchtime=100x -benchmem
// go test -bench=KeepCitySkylineSolver -count 5 -benchtime=10s -benchmem

func BenchmarkKeepCitySkylineSolver(b *testing.B) {
	CG := &leetcode.KeepCitySkyline{}
	for i, v := range keepCitySkylineTestCases {
		b.Run(fmt.Sprintf("keep city skyline test case %d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CG.Solver(v.input)
			}
		})
	}
}
