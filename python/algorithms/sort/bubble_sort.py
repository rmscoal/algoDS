from typing import List


def bubble_sort(arr: List[int]):
    '''
    bubble_sort implements a bubble sort algorithm. Passing a list
    will change the underlying value.
    '''

    n = len(arr)

    # We are traversing to the whole array
    for i in range(n):
        swapped = False

        # For each iteration we would want to swap the left hand side
        # element with the right hand side if the left is bigger.
        #
        # Hence that is why we gradually traverse near the beginning
        # of the array.
        for j in range(n-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = True

        print("Iteration at i =", i)
        print("Array:", arr)

        if swapped == False:
            break


if __name__ == "__main__":
    arr = [14, 33, 27, 35, 10]

    print("Before: ", arr)

    bubble_sort(arr)

    print("Result: ", arr)
