from typing import List


def verify_sorted(arr: List) -> bool:
    n = len(arr)

    if n == 0 or n == 1:
        return True

    return arr[0] < arr[1] and verify_sorted(arr[1:])
