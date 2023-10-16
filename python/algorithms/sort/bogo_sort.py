import random
import sys


def is_sorted(values: list) -> bool:
    for index in range(len(values) - 1):
        if values[index] > values[index + 1]:
            return False
    return True


def bogo_sort(values: list):
    attempts = 0
    while not is_sorted(values):
        random.shuffle(values)
        attempts += 1

    print("Number of attempts", attempts)
    return values


if __name__ == "__main__":
    numbers = [1, 54, 7, 3, 43, 53]
    print(bogo_sort(numbers))
