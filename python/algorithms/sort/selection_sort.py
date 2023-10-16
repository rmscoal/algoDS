from typing import List
import math


def selection_sort(values: List) -> List:
    result = []

    while len(values) != 0:
        smallest = math.inf
        index = 0

        # Loops through to find the index of smallest value
        for i in range(len(values)):
            if values[i] < smallest:
                smallest = values[i]
                index = i

        # Remove that shit and add to our result array
        result.append(values.pop(index))

    return result


if __name__ == "__main__":
    numbers = [45, 23, 42, 65, 2, 4, 6]

    print(selection_sort(numbers))
