package benchmark

import (
	"fmt"
	"testing"

	"algoDS/go/leetcode"
)

var (
	tree1leave1 = &leetcode.TreeNode{Val: 1, Left: nil, Right: nil}
	tree1leave2 = &leetcode.TreeNode{Val: 3, Left: nil, Right: nil}
	tree1       = &leetcode.TreeNode{Val: 2, Left: tree1leave1, Right: tree1leave2}

	tree2leave1 = &leetcode.TreeNode{Val: 7, Left: nil, Right: nil}
	tree2leave2 = &leetcode.TreeNode{Val: 8, Left: nil, Right: nil}
	tree2node5  = &leetcode.TreeNode{Val: 6, Left: nil, Right: tree2leave2}
	tree2node4  = &leetcode.TreeNode{Val: 5, Left: nil, Right: nil}
	tree2node3  = &leetcode.TreeNode{Val: 4, Left: tree2leave1, Right: nil}
	tree2node2  = &leetcode.TreeNode{Val: 3, Left: nil, Right: tree2node5}
	tree2node1  = &leetcode.TreeNode{Val: 2, Left: tree2node3, Right: tree2node4}
	tree2       = &leetcode.TreeNode{Val: 1, Left: tree2node2, Right: tree2node1}
)

var deepestLeavesSumTestCases = []*leetcode.TreeNode{tree1, tree2}

// Use -bench=. to run all benchmark
//
// To avoid executing any test functions in the test files,
// pass a regular expression to the -run flag
//
// Preferred commands:
// go test -bench=DeepestLeavesSum -count 5 -benchtime=100x -benchmem
// go test -bench=DeepestLeavesSum -count 5 -benchtime=10s -benchmem

func BenchmarkDeepestLeavesSum(b *testing.B) {
	DLS := &leetcode.DeepestLeavesSum{}
	for i, v := range deepestLeavesSumTestCases {
		b.Run(fmt.Sprintf("Solver-Test Case %d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DLS.Solver(v)
			}
		})
		b.Run(fmt.Sprintf("FasterSolver-Test Case %d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DLS.FasterSolver(v)
			}
		})
	}
}
