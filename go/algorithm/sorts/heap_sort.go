package sorts

func heapify(arr []int, n, currentIndex int) {
	largestIndex := currentIndex

	left := 2*currentIndex + 1
	right := 2*currentIndex + 2

	if left < n && arr[largestIndex] < arr[left] {
		largestIndex = left
	}

	if right < n && arr[largestIndex] < arr[right] {
		largestIndex = right
	}

	if largestIndex != currentIndex {
		arr[largestIndex], arr[currentIndex] = arr[currentIndex], arr[largestIndex]
		heapify(arr, n, largestIndex)
	}
}

func heapsort(arr []int) {
	n := len(arr)

	for i := (n / 2) - 1; i > -1; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}
