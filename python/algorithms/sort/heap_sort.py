import heapq  # just testing


def heapify(arr, N, current_index):
    largest_index = current_index

    left_index = 2 * current_index + 1
    right_index = 2 * current_index + 2

    if left_index < N and arr[largest_index] < arr[left_index]:
        largest_index = left_index

    if right_index < N and arr[largest_index] < arr[right_index]:
        largest_index = right_index

    if largest_index != current_index:
        arr[current_index], arr[largest_index] = arr[largest_index], arr[current_index]
        heapify(arr, N, largest_index)


def heapsort(arr):
    N = len(arr)

    for i in range(N//2 - 1, -1, -1):
        heapify(arr, N, i)

    for i in range(N-1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapify(arr, i, 0)


if __name__ == "__main__":
    heapify_arr_trial = [4, 7, 3, 4, 1, 9, 6, 7]
    heapq.heapify(heapify_arr_trial)

    x = [5, 6, 7, 4, 1, 5, 4]
    heapsort(x)

    print(x)
