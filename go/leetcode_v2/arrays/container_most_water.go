package arrays

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1

	for left < right {
		if max < (right-left)*min(height[left], height[right]) {
			max = (right - left) * min(height[left], height[right])
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
