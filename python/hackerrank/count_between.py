from bisect import bisect_left, bisect_right

def count_elements_in_ranges(arr: list[int], low: list[int], high: list[int]):
    # Sort the array
    arr.sort()

    results = []
    for l, h in zip(low, high):
        # Find bounds using binary search
        lower_bound = bisect_left(arr, l)
        upper_bound = bisect_right(arr, h)
        results.append(upper_bound - lower_bound)

    return results

# Example Usage
arrr = [1, 2, 2, 3, 4]
lows = [0, 2]
highs = [2, 4]
print(count_elements_in_ranges(arrr, lows, highs))  # Output: [3, 4]
