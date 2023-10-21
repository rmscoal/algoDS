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

        if swapped == False:
            break

def bubble_sort_v2(arr: List[int]):
    n = len(arr)

    for i in range(n):
        swapped = False
        for j in range(n-1-i):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = True

        if not swapped:
            break



if __name__ == "__main__":
    arr = [14, 33, 27, 35, 10]

    print("Before: ", arr)
    bubble_sort(arr)
    print("Result: ", arr)

    arr = [5,4,3,2,1]

    print("Before: ", arr)
    bubble_sort_v2(arr)
    print("Result: ", arr)
