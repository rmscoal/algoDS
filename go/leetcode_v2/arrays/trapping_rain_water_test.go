package arrays

import (
	"testing"
)

func trap_slow(heights []int) int {
	if len(heights) <= 2 {
		return 0
	}

	rights := heights[2:]
	water := make([]int, len(heights))
	maxLeft := heights[0]
	maxRight := findMaxInSlice(rights)

	for i := 1; i < len(heights)-1; i++ {
		if heights[i-1] > maxLeft {
			maxLeft = heights[i-1]
		}
		rights = heights[i+1:]
		if heights[i] == maxRight {
			maxRight = findMaxInSlice(rights)
		}

		minHeight := min(maxLeft, maxRight)
		if minHeight >= heights[i] {
			water[i] = minHeight - heights[i]
		}
	}

	total := 0
	for _, v := range water {
		total += v
	}

	return total
}

func findMaxInSlice(arr []int) int {
	m := 0

	for _, v := range arr {
		m = max(m, v)
	}

	return m
}

func Test_TrapSlow(t *testing.T) {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	if result := trap_slow(heights); result != 6 {
		t.Fatalf("expected 6 but got %d", result)
	}

	heights = []int{4, 2, 0, 3, 2, 5}
	if result := trap_slow(heights); result != 9 {
		t.Fatalf("expected 9 but got %d", result)
	}

	heights = []int{4, 2, 3}
	if result := trap_slow(heights); result != 1 {
		t.Fatalf("expected 1 but got %d", result)
	}
}

func trap_faster(heights []int) int {
	// 4, 2, 3

	// 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1
	//    i
	// x     y

	// 4, 2, 0, 3, 2, 5
	//    i
	// x              y

	// 5, 0, 1, 2, 0, 2
	//             i
	//          x     y

	total := 0

	for i := 1; i < len(heights)-1; i++ {
		left := i - 1
		right := i + 1

		if heights[i] < heights[left] {
			continue
		}

		if heights[right] < heights[left] {
			for j := i + 2; j < len(heights); j++ {
				if heights[j] >= heights[left] {
					right = j
					break
				} else if heights[j] > heights[right] {
					right = j
				}
			}

			// That means we didnt find any right that is greater than
			// left thus we are unable to store any water here.
			if right == i+1 && heights[i] >= heights[right] {
				continue
			}
		}

		for i < right {
			trapped := min(heights[left], heights[right]) - heights[i]
			if trapped > 0 {
				total += trapped
			}
			if i == right-1 {
				break
			}
			i++
		}
	}

	return total
}

func TestTrap_Faster(t *testing.T) {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	if result := trap_faster(heights); result != 6 {
		t.Fatalf("expected 6 but got %d", result)
	}

	heights = []int{4, 2, 0, 3, 2, 5}
	if result := trap_faster(heights); result != 9 {
		t.Fatalf("expected 9 but got %d", result)
	}

	heights = []int{4, 2, 3}
	if result := trap_faster(heights); result != 1 {
		t.Fatalf("expected 1 but got %d", result)
	}

	heights = []int{4, 2, 3, 2}
	if result := trap_faster(heights); result != 1 {
		t.Fatalf("expected 1 but got %d", result)
	}

	heights = []int{0, 7, 1, 4, 6}
	if result := trap_faster(heights); result != 7 {
		t.Fatalf("expected 7 but got %d", result)
	}

	heights = []int{5, 0, 1, 2, 0, 2}
	if result := trap_faster(heights); result != 5 {
		t.Fatalf("expected 5 but got %d", result)
	}
}

func trap_O_of_n_mem(heights []int) int {
	maxLefts := make([]int, len(heights))
	maxRights := make([]int, len(heights))

	maxLeft := heights[0]
	for i := 1; i < len(heights); i++ {
		if heights[i-1] > maxLeft {
			maxLeft = heights[i-1]
		}

		maxLefts[i] = maxLeft
	}

	maxRight := heights[len(heights)-1]
	for i := len(heights) - 2; i >= 0; i-- {
		if heights[i+1] > maxRight {
			maxRight = heights[i+1]
		}

		maxRights[i] = maxRight
	}

	total := 0
	for i := 0; i < len(heights); i++ {
		trapped := min(maxLefts[i], maxRights[i]) - heights[i]
		if trapped > 0 {
			total += trapped
		}
	}

	return total
}

func TestTrap_On_Mem(t *testing.T) {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	if result := trap_O_of_n_mem(heights); result != 6 {
		t.Fatalf("expected 6 but got %d", result)
	}

	heights = []int{4, 2, 0, 3, 2, 5}
	if result := trap_O_of_n_mem(heights); result != 9 {
		t.Fatalf("expected 9 but got %d", result)
	}

	heights = []int{4, 2, 3}
	if result := trap_O_of_n_mem(heights); result != 1 {
		t.Fatalf("expected 1 but got %d", result)
	}

	heights = []int{4, 2, 3, 2}
	if result := trap_O_of_n_mem(heights); result != 1 {
		t.Fatalf("expected 1 but got %d", result)
	}

	heights = []int{0, 7, 1, 4, 6}
	if result := trap_O_of_n_mem(heights); result != 7 {
		t.Fatalf("expected 7 but got %d", result)
	}

	heights = []int{5, 0, 1, 2, 0, 2}
	if result := trap_O_of_n_mem(heights); result != 5 {
		t.Fatalf("expected 5 but got %d", result)
	}
}

func trap_two_pointer_best(heights []int) int {
	var (
		total             int = 0
		left, right       int = 0, len(heights) - 1
		leftMax, rightMax int = 0, 0
	)

	for left < right {
		if heights[left] < heights[right] {
			if heights[left] >= leftMax {
				leftMax = heights[left]
			} else {
				total += leftMax - heights[left]
			}
			left++
		} else {
			if heights[right] >= rightMax {
				rightMax = heights[right]
			} else {
				total += rightMax - heights[right]
			}
			right--
		}
	}

	return total
}

func TestTrap_Two_Pointer_Best(t *testing.T) {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	if result := trap_two_pointer_best(heights); result != 1 {
		t.Fatalf("expected 1 but got %d", result)
	}

	heights = []int{4, 2, 3, 2}
	if result := trap_two_pointer_best(heights); result != 1 {
		t.Fatalf("expected 1 but got %d", result)
	}

	heights = []int{0, 7, 1, 4, 6}
	if result := trap_two_pointer_best(heights); result != 7 {
		t.Fatalf("expected 7 but got %d", result)
	}

	heights = []int{5, 0, 1, 2, 0, 2}
	if result := trap_two_pointer_best(heights); result != 5 {
		t.Fatalf("expected 5 but got %d", result)
	}
}
