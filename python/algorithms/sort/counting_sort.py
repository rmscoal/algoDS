from typing import List


def find_largest_value_in_list(arr: List[int]) -> int:
    # Largest value in the input array
    largest = 0
    for v in arr:
        largest = max(largest, v)

    return largest


def prefix_sum(arr: List[int]) -> List[int]:
    cursor = 0
    for i in range(0, len(arr)):
        cursor += arr[i]

        arr[i] = cursor

    return arr


def counting_sort(arr: List[int]) -> List[int]:
    # Our result array
    result = [0 for _ in range(len(arr))]

    largest = find_largest_value_in_list(arr)

    # Make a temporary count
    count_arr = [0 for _ in range(largest+1)]

    # Fill in the count of occurence
    for v in arr:
        count_arr[v] += 1

    # Next do prefix sum
    count_arr = prefix_sum(count_arr)

    for i in range(len(arr) - 1, -1, -1):
        result[count_arr[arr[i]] - 1] = arr[i]
        count_arr[arr[i]] -= 1

    return result


if __name__ == "__main__":
    print(counting_sort([2, 5, 3, 0, 2, 3, 0, 3, 4,
          5, 2, 5, 7, 20, 194, 2823, 482, 12843]))
