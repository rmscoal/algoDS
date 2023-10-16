# importing sys
import sys
from typing import Tuple

# adding Folder_2/subfolder to the system path
sys.path.insert(
    0, '/Users/rifkysatyana/Developer/Github/algoDS/python/data_structure'
)

from linked_list import LinkedList

def split(linked_list: LinkedList) -> Tuple[LinkedList, LinkedList]:
    if linked_list is None or linked_list.head is None:
        left_half = linked_list
        right_half = None

        return left_half, right_half
    else:
        size = linked_list.size()
        mid = size//2

        mid_node = linked_list.node_at_index(mid-1)
        left_half = linked_list
        right_half = LinkedList()
        right_half.head = mid_node.next_node
        mid_node.next_node = None

        return left_half, right_half

def merge(left: LinkedList, right: LinkedList) -> LinkedList:
    # The result linked list
    result = LinkedList()

    # Add fake head that later will be discarded
    result.add(0)

    # Set cursor of the result head
    result_cursor = result.head

    # Obtain the head for each linked list
    left_cursor = left.head
    right_cursor = right.head

    # Iterating left and right until we reach tail of either
    while left_cursor and right_cursor:
        if left_cursor.data < right_cursor.data:
            result_cursor.next_node = left_cursor
            left_cursor = left_cursor.next_node
        else:
            result_cursor.next_node = right_cursor
            right_cursor = right_cursor.next_node

        result_cursor = result_cursor.next_node

    if left_cursor is None:
        result_cursor.next_node = right_cursor
    if right_cursor is None:
        result_cursor.next_node = left_cursor

    # Remove fake head
    result.head = result.head.next_node
        
    return result

def merge_sort_linked_list(linked_list: LinkedList):
    if not isinstance(linked_list, LinkedList):
        raise Exception("Not linked list")

    if linked_list.size() == 1:
        return linked_list
    elif linked_list.is_empty():
        return linked_list

    left_half, right_half = split(linked_list)
    left = merge_sort_linked_list(left_half)
    right = merge_sort_linked_list(right_half)

    return merge(left, right)

if __name__ == "__main__":
    l = LinkedList()
    l.add(10)
    l.add(2)
    l.add(44)
    l.add(15)
    l.add(200)

    print(l)

    sorted_linked_list = merge_sort_linked_list(l)
    print(sorted_linked_list)
