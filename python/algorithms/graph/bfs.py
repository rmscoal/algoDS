from collections import defaultdict


class Graph:
    def __init__(self):
        self.graph = defaultdict(list)

    def add_edge(self, u: int, v: int):
        self.graph[u].append(v)

    def BFS(self, starting_node: int):
        visited = [False] * (max(self.graph) + 1)

        queue = []

        queue.append(starting_node)
        visited[starting_node] = True

        while queue:
            s = queue.pop(0)

            print(s, end=' ')

            for neighbour in self.graph[s]:
                if not visited[neighbour]:
                    queue.append(neighbour)
                    visited[neighbour] = True


if __name__ == "__main__":
    g = Graph()

    g.add_edge(0, 1)
    g.add_edge(0, 2)
    g.add_edge(1, 2)
    g.add_edge(2, 0)
    g.add_edge(2, 3)
    g.add_edge(3, 3)

    g.BFS(2)
