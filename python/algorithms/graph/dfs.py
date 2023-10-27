from collections import defaultdict
from typing import Set


class Graph:
    def __init__(self):
        self.graph = defaultdict(list)

    def add_edge(self, u: int, v: int):
        self.graph[u].append(v)

    def DFS_util(self, v: int, visited: Set):
        visited.add(v)
        print(v, end=' ')

        for neighbour in self.graph[v]:
            if neighbour not in visited:
                self.DFS_util(neighbour, visited)

    def DFS(self, v):
        visited = set()

        self.DFS_util(v, visited)


if __name__ == "__main__":
    g = Graph()

    g.add_edge(0, 1)
    g.add_edge(0, 2)
    g.add_edge(1, 2)
    g.add_edge(2, 0)
    g.add_edge(2, 3)
    g.add_edge(3, 3)

    g.DFS(2)
