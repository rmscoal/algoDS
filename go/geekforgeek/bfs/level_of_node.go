package bfs

func nodeContainsInVisited(visited []int, node int) bool {
	for _, v := range visited {
		if v == node {
			return true
		}
	}

	return false
}

func nodeLevel(edges [][]int, startingNode, target int) int {
	graph := map[int][]int{}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	queue := []int{startingNode}
	visited := []int{}
	level := 0

	for len(queue) >= 1 {
		// Since we are using BFS we want ot exhaust all current level node.
		currentLevelNodeCount := len(queue)
		for currentLevelNodeCount > 0 {
			node := queue[0]

			if node == target {
				return level
			}

			queue = queue[1:]
			if !nodeContainsInVisited(visited, node) {
				visited = append(visited, node)
			}

			for _, neighbour := range graph[node] {
				if neighbour == target {
					return level + 1
				}

				if !nodeContainsInVisited(visited, neighbour) {
					visited = append(visited, neighbour)
					queue = append(queue, neighbour)
				}
			}

			currentLevelNodeCount--
		}
		level++
	}

	return -1
}
