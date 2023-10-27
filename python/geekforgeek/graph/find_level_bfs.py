from collections import defaultdict
from typing import List, Tuple, cast
import math


class Graph:
    def __init__(self):
        self.graph = defaultdict(list)
        self.max_node = -math.inf
        self.levels = defaultdict(int)

    def add_edge(self, dir: Tuple[int, int]):
        self.graph[dir[0]].append(dir[1])
        if self.max_node < dir[0] or self.max_node > dir[1]:
            self.max_node = max(dir[0], dir[1])

    def BFS_search_with_level(self, starting_node: int, target: int) -> int:
        visited = [False] * int(self.max_node + 1)
        queue = []

        queue.append(starting_node)
        visited[starting_node] = True
        self.levels[starting_node] = 0

        while queue:
            s = queue.pop(0)

            for neighbour in self.graph[s]:
                if not visited[neighbour]:
                    queue.append(neighbour)
                    visited[neighbour] = True
                    self.levels[neighbour] = self.levels[s] + 1
                if neighbour == target:
                    return self.levels[neighbour]

        return -1


def find_level(edges: List[Tuple[int, int]], x: int) -> int:
    g = Graph()
    starting_node = None

    for i in range(len(edges)):
        if i == 0:
            starting_node = edges[i][0]

        g.add_edge(edges[i])

    return g.BFS_search_with_level(cast(int, starting_node), x)


def find_level_2(V, Edges, X):
    # Create an adjacency list representation of the graph
    graph = defaultdict(list)
    for u, v in Edges:
        graph[u].append(v)
        graph[v].append(u)

    # Initialize a visited array and a queue for BFS
    visited = [False] * V
    queue = []
    level = 0

    # Start BFS from node 0
    queue.append(0)
    visited[0] = True

    while queue:
        size = len(queue)
        for _ in range(size):
            node = queue.pop(0)
            if node == X:
                return level
            for neighbor in graph[node]:
                if not visited[neighbor]:
                    queue.append(neighbor)
                    visited[neighbor] = True
        level += 1

    # If X is not found, return -1
    return -1


if __name__ == "__main__":
    print("Result", find_level(
        [(0, 1), (0, 2), (1, 3), (2, 4)], 3), "suppose to be 2")
    print("Result", find_level(
        [(0, 1), (0, 2), (1, 3), (2, 4)], 5), "suppose to be -1")
    print("Result", find_level(
        [(0, 1), (0, 2), (0, 3), (1, 4), (2, 5)], 3), "suppose to be 1")
    print("Result", find_level(
        [(0, 1), (0, 2), (0, 3), (1, 4), (2, 5)], 5), "suppose to be 2")
    print("Result", find_level(
        [(0, 1), (0, 2), (1, 3), (2, 4)], 2), "suppose to be 1")

    print()

    print("Result", find_level_2(
        6, [(0, 1), (0, 2), (0, 3), (1, 4), (2, 5)], 5), "suppose to be 2")
    print("Result", find_level_2(
        6, [(0, 1), (0, 2), (1, 3), (2, 4)], 5), "suppose to be -1")
