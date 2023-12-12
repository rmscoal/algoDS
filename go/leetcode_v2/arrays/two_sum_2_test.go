package arrays

import (
	"testing"
)

func twoSum_II_Hash(numbers []int, target int) []int {
	// numbers = [2,7,11,15], target = 9
	// output = [1,2]

	hash := make(map[int]int)

	for i, v := range numbers {
		if _, found := hash[v]; found {
			return []int{hash[v] + 1, i + 1}
		}

		hash[target-v] = i
	}

	return []int{}
}

func TestTwoSum_II_Hash(t *testing.T) {
	if result := twoSum_II_Hash([]int{2, 7, 11, 15}, 9); result[0] != 1 && result[1] != 2 {
		t.Fatalf("Expected to be [1,2] but got %+v", result)
	}
	if result := twoSum_II_Hash([]int{2, 3, 4}, 6); result[0] != 1 && result[1] != 3 {
		t.Fatalf("Expected to be [1,3] but got %+v", result)
	}
	if result := twoSum_II_Hash([]int{-1, 0}, -1); result[0] != 1 && result[1] != 2 {
		t.Fatalf("Expected to be [1,2] but got %+v", result)
	}
}

func twoSum_II_Two_Pointer(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		current := numbers[left] + numbers[right]
		if current == target {
			return []int{left + 1, right + 1}
		}

		if current > target {
			right--
		} else {
			left++
		}
	}

	return []int{0, 1}
}

func TestTwoSum_II_Two_Pointer(t *testing.T) {
	if result := twoSum_II_Two_Pointer([]int{2, 7, 11, 15}, 9); result[0] != 1 && result[1] != 2 {
		t.Fatalf("Expected to be [1,2] but got %+v", result)
	}
	if result := twoSum_II_Two_Pointer([]int{2, 3, 4}, 6); result[0] != 1 && result[1] != 3 {
		t.Fatalf("Expected to be [1,3] but got %+v", result)
	}
	if result := twoSum_II_Two_Pointer([]int{-1, 0}, -1); result[0] != 1 && result[1] != 2 {
		t.Fatalf("Expected to be [1,2] but got %+v", result)
	}
	if result := twoSum_II_Two_Pointer([]int{1, 2, 3, 4}, 6); result[0] != 2 && result[1] != 4 {
		t.Fatalf("Expected to be [2,4] but got %+v", result)
	}
	if result := twoSum_II_Two_Pointer([]int{-1, 2, 4, 5, 6}, 7); result[0] != 2 && result[1] != 4 {
		t.Fatalf("Expected to be [2,4] but got %+v", result)
	}
	if result := twoSum_II_Two_Pointer([]int{-10, -5, -1, 1, 2, 6}, 7); result[0] != 4 && result[1] != 6 {
		t.Fatalf("Expected to be [4,6] but got %+v", result)
	}
	if result := twoSum_II_Two_Pointer([]int{3, 24, 50, 79, 88, 150, 345}, 200); result[0] != 3 && result[1] != 6 {
		t.Fatalf("Expected to be [3,6] but got %+v", result)
	}
	if result := twoSum_II_Two_Pointer([]int{5, 25, 75}, 100); result[0] != 2 && result[1] != 3 {
		t.Fatalf("Expected to be [2,3] but got %+v", result)
	}
}
