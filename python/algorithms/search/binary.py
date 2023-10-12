from typing import List


def binary_search(arr: List[int], x: int) -> int:
    l = 0
    r = len(arr) - 1

    while (l < r):
        m = l + ((r-l) // 2)

        if x == arr[m]:
            return m
        elif x > arr[m]:
            l = m + 1
        else:
            r = m - 1

    return -1


def binary_search_recursive(arr: List[int], x: int) -> int:
    if len(arr) == 0:
        return - 1

    m = (len(arr) - 1)//2

    if x == arr[m]:
        return m
    elif x > arr[m]:
        return binary_search_recursive(arr[m+1:], x)
    else:
        return binary_search_recursive(arr[:m], x)


if __name__ == "__main__":
    print("Binary search", binary_search([1, 2, 3, 4, 5, 6, 7, 8], 6))
    print("Binary search", binary_search([1, 2, 3, 4, 5, 6, 7, 8], 2))

    print("Binary search recursive", binary_search_recursive(
        [1, 2, 3, 4, 5, 6, 7, 8], 2))
    print("Binary search recursive", binary_search_recursive(
        [1, 2, 3, 4, 5, 6, 7, 8], 10))
