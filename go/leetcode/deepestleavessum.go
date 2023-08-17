package leetcode

import (
	"sync"
)

var mu sync.Mutex

func (p *DeepestLeavesSum) Solver(root *TreeNode) int {
	if root == nil {
		return 0
	}

	slc := []int{}

	p.helper(root, &slc, 0)

	return p.sum(&slc)
}

func (p *DeepestLeavesSum) helper(node *TreeNode, slc *[]int, curr int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		mu.Lock()
		*slc = append(*slc, node.Val, curr)
		mu.Unlock()
		return
	}

	if node.Left == nil {
		p.helper(node.Right, slc, curr+1)
		return
	}

	if node.Right == nil {
		p.helper(node.Left, slc, curr+1)
		return
	}

	p.helper(node.Left, slc, curr+1)
	p.helper(node.Right, slc, curr+1)
}

func (p *DeepestLeavesSum) sum(slc *[]int) int {
	// Find the max
	max := 0
	for i := 1; i < len(*slc); i += 2 {
		if (*slc)[i] > max {
			max = (*slc)[i]
		}
	}

	sum := 0
	for i := 0; i < len(*slc); i += 2 {
		if (*slc)[i+1] == max {
			sum += (*slc)[i]
		}
	}

	return sum
}

// It turns out, by benchmark result, that this solution is
// slower when it comes to larger tree.
func (p *DeepestLeavesSum) FasterSolver(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	sum := 0
	for len(queue) > 0 {

		nodesIncurrentHeight := len(queue)
		sum = 0
		for i := 0; i < nodesIncurrentHeight; i++ {
			curr := queue[0]
			queue = queue[1:]
			sum += curr.Val
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

	}
	return sum
}
