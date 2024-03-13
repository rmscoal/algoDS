package arrays

import "sort"

// Max Heapify
func heapify(arr []int, length int, index int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left < length && arr[largest] < arr[left] {
		largest = left
	}

	if right < length && arr[largest] < arr[right] {
		largest = right
	}

	if largest != index {
		// Swap our current index with the largest one
		arr[index], arr[largest] = arr[largest], arr[index]
		heapify(arr, length, largest)
	}
}

func heapsort(arr []int) {
	N := len(arr)

	for i := (N / 2) - 1; i > -1; i-- {
		heapify(arr, N, i)
	}

	for i := N - 1; i > N-3; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
}

func maxProduct(nums []int) int {
	if len(nums) == 2 {
		return (nums[0] - 1) * (nums[1] - 1)
	}

	heapsort(nums)

	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}

func maxProductV2(nums []int) int {
	if len(nums) == 2 {
		return (nums[0] - 1) * (nums[1] - 1)
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	n := len(nums)
	return (nums[n-1] - 1) * (nums[n-2] - 1)
}
