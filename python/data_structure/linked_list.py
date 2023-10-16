from typing import Optional, cast


class Node:
    data = None
    next_node: Optional['Node'] = None

    def __init__(self, data):
        self.data = data
        self.next_node = None

    def __repr__(self) -> str:
        return "<Node data: %s>" % self.data


class LinkedList:
    """
    Singly linked list
    """
    head: Optional['Node'] = None

    def __init__(self):
        self.head = None

    def __repr__(self) -> str:
        """
        Print self representation as a string.

        O(n)
        """
        nodes = []
        current = self.head

        while current:
            if current is self.head:
                nodes.append("[Head: %s]" % current.data)
            elif current.next_node is None:
                nodes.append("[Tail: %s]" % current.data)
            else:
                nodes.append("[%s]" % current.data)

            current = current.next_node

        return '-> '.join(nodes)

    def is_empty(self) -> bool:
        """
        Indicates whether the instances does not have any nodes.

        O(1)
        """
        return self.head is None

    def size(self) -> int:
        """
        Calculate length of linked list.

        O(n)
        """
        count = 0
        current = self.head

        while current:
            count += 1
            current = current.next_node

        return count

    def add(self, data):
        """
        Adds prepends new node data.

        O(1)
        """
        new_node = Node(data)
        new_node.next_node = self.head
        self.head = new_node

    def search(self, data) -> Optional[Node]:
        """
        Search a node containing given data.

        O(n)
        """
        current = self.head

        while current:
            if current.data == data:
                return current

            current = current.next_node

        return None

    def insert(self, data, index) -> None:
        """
        Inserts a new node at given index

        O(n)
        """
        if index == 0:
            self.add(data)
            return

        if index < 0:
            raise Exception("Index is less than zero")

        if self.size() - 1 < index:
            raise Exception("Index out of range")

        new_node = Node(data)
        current = cast(Node, self.head)
        pos = 0

        while pos != index - 1:
            current = cast(Node, current.next_node)
            pos += 1

        new_node.next_node = current.next_node
        current.next_node = new_node

    def remove(self, data) -> None:
        """
        Remove a node based on the first match of the given data.

        O(n)
        """
        if self.is_empty():
            raise Exception("Instance is empty")

        current = cast(Node, self.head)

        if current.data == data:
            self.head = current.next_node

        while current.next_node:
            if current.next_node.data == data:
                current.next_node = current.next_node.next_node
                break


    def node_at_index(self, index) -> Node:
        if self.head is None:
            raise Exception("Instance is empty")

        current = cast(Node, self.head)

        while index != 0:
            current = cast(Node, current.next_node)
            if current is None:
                raise Exception("Index out of range")
            index -= 1

        return current


if __name__ == "__main__":
    l = LinkedList()
    l.add(1)
    l.add(2)
    l.add(3)
    l.add(4)
    print(l)

    n = l.search(3)
    print(n)

    l.insert(45, 1)
    print(l)
    l.insert(50, 4)
    print(l)

    l.remove(45)
    print(l)

    print("Node at index 2")
    c = l.node_at_index(2)
    print(c)
