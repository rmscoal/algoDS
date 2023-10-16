from typing import List, Tuple

import random


def split_with_pivot(arr: List) -> Tuple[List, int, List]:
    # Randomly selects our pivot
    pivot = arr.pop(random.randint(0, len(arr) - 1))

    # The left contains elements which are less than the pivot
    left = []
    # The right contains elements which are more than the pivot
    right = []

    while len(arr) != 0:
        if arr[0] < pivot:
            left.append(arr.pop(0))
        else:
            right.append(arr.pop(0))

    return left, pivot, right


def merge_with_pivot(left: List, pivot: int, right: List) -> List:
    result = []

    left.append(pivot)
    result.extend(left)
    result.extend(right)

    return result


def quicksort(arr: List) -> List:
    if len(arr) <= 1:
        return arr

    left, pivot, right = split_with_pivot(arr)

    left = quicksort(left)
    right = quicksort(right)

    return merge_with_pivot(left, pivot, right)


if __name__ == "__main__":
    numbers = [4, 3, 6, 1, 8, 2, 9, 14, 10, 23, 3, 21, 20]

    print("Result", quicksort(numbers))
