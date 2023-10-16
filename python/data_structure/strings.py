from typing import List


def gatherAllSubstring(s: str) -> List[str]:
    n = len(s)
    arr = []

    for i in range(0, n, 1):
        for j in range(i, n, 1):
            substr = ""
            for k in range(i, j+1):
                substr += s[k]
            arr.append(substr)

    return arr


def bubble_sort_by_len_str(arr: List[str]) -> List[str]:
    n = len(arr)

    for i in range(n):
        swapped = False
        for j in range(n-1-i):
            if len(arr[j]) > len(arr[j+1]):
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = True

        if swapped == False:
            break

    return arr


if __name__ == "__main__":
    arr = bubble_sort_by_len_str(gatherAllSubstring("hello"))

    for i in arr:
        print(i)
