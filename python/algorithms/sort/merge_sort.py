from typing import List, Tuple
from utils import verify_sorted
import sys


def split(arr: List) -> Tuple[List, List]:
    '''
    Divide the unsorted list at midpoint into sublists.
    Returns two sublists - left and right.

    O(log n)
    '''

    m = (len(arr))//2
    return arr[:m], arr[m:]


def merge(left_arr: List, right_arr: List) -> List:
    '''
    Merges two sorted lists such that it returns a new sorted lists.

    O(n)
    '''

    result = []
    left_length = len(left_arr)
    right_length = len(right_arr)
    left_index = 0
    right_index = 0

    while left_index < left_length and right_index < right_length:
        if left_arr[left_index] < right_arr[right_index]:
            result.append(left_arr[left_index])
            left_index += 1
        else:
            result.append(right_arr[right_index])
            right_index += 1

    if left_index < left_length:
        result.extend(left_arr[left_index:])

    if right_index < right_length:
        result.extend(right_arr[right_index:])

    # while left_index < left_length:
    #     result.append(left_arr[left_index])
    #     left_index += 1

    # while right_index < right_length:
    #     result.append(right_arr[right_index])
    #     right_index += 1

    return result


def merge_sort(arr: List) -> List:
    """
    Sorts a list in ascending order
    Returns a new sorted list

    Divide: Find the midpoint of the list and divide into sublists
    Conquer: Recursively sort the sublists created in previous step
    Combine: Merge the sorted sublists created in previous step

    O(n log n)
    """

    if len(arr) <= 1:
        return arr

    left_half, right_half = split(arr)
    left = merge_sort(left_half)
    right = merge_sort(right_half)

    return merge(left, right)


if __name__ == "__main__":
    # Set recursion limit
    sys.setrecursionlimit(1000)

    arr = [2, 1, 5, 7, 3, 9]
    print("Before", arr)

    arr = merge_sort(arr)
    print("After Merge sort", arr)

    print("Verified", verify_sorted(arr))
