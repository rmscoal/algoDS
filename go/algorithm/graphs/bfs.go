package graphs

import "fmt"

func bfs(edges [][]int, startingNode int) {
	// Firstly transform our edges to a map to represent a graph
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	visited := make([]int, 0)
	queue := []int{startingNode}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if nodeContainsInVisited(visited, node) {
			continue
		}
		fmt.Printf("%d\t", node)
		visited = append(visited, node)

		queue = append(queue, graph[node]...)
	}
}
