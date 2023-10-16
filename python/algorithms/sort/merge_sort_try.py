from typing import List, Tuple


def split(arr: List) -> Tuple[List, List]:
    m = len(arr) // 2

    return arr[:m], arr[m:]


def merge(left: List, right: List) -> List:
    result = []

    left_length = len(left)
    left_index = 0

    right_length = len(right)
    right_index = 0

    while left_index < left_length and right_index < right_length:
        if left[left_index] < right[right_index]:
            result.append(left[left_index])
            left_index += 1
        else:
            result.append(right[right_index])
            right_index += 1

    if left_index < left_length:
        result.extend(left[left_index:])

    if right_index < right_length:
        result.extend(right[right_index:])

    return result


def merge_sort(arr: List) -> List:
    if len(arr) <= 1:
        return arr

    left, right = split(arr)
    left = merge_sort(left)
    right = merge_sort(right)

    return merge(left, right)


if __name__ == "__main__":
    numbers = [4, 3, 6, 1, 8, 2, 9, 14, 10, 23, 3, 21, 20]

    print("Result", merge_sort(numbers))
