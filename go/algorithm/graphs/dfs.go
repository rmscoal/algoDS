package graphs

import "fmt"

func nodeContainsInVisited(visited []int, node int) bool {
	for _, v := range visited {
		if v == node {
			return true
		}
	}

	return false
}

func dfsNoRecurr(edges [][]int, startingNode int) {
	// DFS uses a stack to maintain the buffer of node
	// to visit next...

	// Firstly transform our edges to a map to represent a graph
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	fmt.Printf("Graph %v\n", graph)

	// Visited is used to store the visited node.. If it is visited
	// there is no use to visit that particular node again.
	visited := make([]int, 0)
	stack := make([]int, 0)

	// We append our statingNode to the stack
	stack = append(stack, startingNode)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		// Pop from our stack
		stack = stack[:len(stack)-1]

		if !nodeContainsInVisited(visited, node) {
			visited = append(visited, node)
			// Append our neighbour to our stack
			for _, neighbour := range graph[node] {
				if !nodeContainsInVisited(stack, neighbour) {
					stack = append(stack, neighbour)
				}
			}
			// stack = append(stack, graph[node]...)
			println(node)
		}
	}
}

type DFS struct {
	graph map[int][]int
}

func (d *DFS) Init(edges [][]int) {
	d.graph = make(map[int][]int)
	for _, edge := range edges {
		d.graph[edge[0]] = append(d.graph[edge[0]], edge[1])
	}
}

func (d *DFS) DFSUtil(v int, visited *[]int) {
	*visited = append(*visited, v)
	fmt.Printf("%d\t", v)

	for _, neighbour := range d.graph[v] {
		if !nodeContainsInVisited(*visited, neighbour) {
			d.DFSUtil(neighbour, visited)
		}
	}
}

func (d *DFS) DFS(startingNode int) {
	visited := make([]int, 0)

	d.DFSUtil(startingNode, &visited)
}
