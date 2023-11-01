package graphs

import "testing"

func TestBFS(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 3}, {3, 3}}
	bfs(edges, 1)
	println()
	edges = [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 3}, {3, 3}}
	bfs(edges, 2)
}