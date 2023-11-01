package graphs

import "testing"

func TestDFSNoRecurr(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 3}, {3, 3}}
	dfsNoRecurr(edges, 1)

	edges = [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 3}, {3, 3}}
	dfsNoRecurr(edges, 2)
}

func TestDFSRecurr(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 3}, {3, 3}}
	d := &DFS{}
	d.Init(edges)
	d.DFS(1)

	println()
	edges = [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 3}, {3, 3}}
	d.Init(edges)
	d.DFS(2)
}
