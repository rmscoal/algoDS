package sorting

func split(arr []int32) ([]int32, []int32) {
	mid := len(arr) / 2

	return arr[:mid], arr[mid:]
}

func merge(lArr, rArr []int32, count *int64) []int32 {
	var result []int32

	left := 0
	right := 0

	for left < len(lArr) && right < len(rArr) {
		if lArr[left] <= rArr[right] {
			result = append(result, lArr[left])
			left++
		} else {
			result = append(result, rArr[right])
			right++
			*count += int64(len(lArr) - left)
		}
	}

	result = append(result, lArr[left:]...)
	result = append(result, rArr[right:]...)

	return result
}

func mergesort(arr []int32, count *int64) []int32 {
	if len(arr) <= 1 {
		return arr
	}

	left, right := split(arr)

	left = mergesort(left, count)
	right = mergesort(right, count)

	return merge(left, right, count)
}

func countInversions(arr []int32) int64 {
	var count int64 = 0

	mergesort(arr, &count)

	return count
}
