package leetcode

func BinarySearch(arr []int, target, low, high int) int {
	if low > high {
		return -1
	} else {
		mid := low + (high-low)/2
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			return BinarySearch(arr, target, mid+1, high)
		} else {
			return BinarySearch(arr, target, low, mid-1)
		}
	}
}
