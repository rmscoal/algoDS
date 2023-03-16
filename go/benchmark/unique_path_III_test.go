// UP3 := &leetcode.UniquePathThree{}
// // UP3.Solver([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}})
// UP3.Solver([][]int{{0, 1}, {2, 0}})

package benchmark

import (
	"competitive/leetcode"
	"fmt"
	"testing"
)

var uniquePathIIITestCases = []struct {
	test [][]int
}{
	{
		test: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}},
	},
	{
		test: [][]int{{0, 1}, {2, 0}},
	},
}

// Use -bench=. to run all benchmark
//
// To avoid executing any test functions in the test files,
// pass a regular expression to the -run flag
//
// Preferred commands:
// go test -bench=UniquePathThree -count 5 -benchtime=100x -benchmem
// go test -bench=UniquePathThree -count 5 -benchtime=10s -benchmem

func BenchmarkUniquePathThree(b *testing.B) {
	UP3 := &leetcode.UniquePathThree{}
	for i, v := range uniquePathIIITestCases {
		b.Run(fmt.Sprintf("test case no-%d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UP3.Solver(v.test)
			}
		})
	}
}
