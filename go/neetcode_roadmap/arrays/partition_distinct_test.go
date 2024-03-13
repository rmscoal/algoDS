package arrays

import "testing"

func partitionDisjoint(nums []int) int {
	// currentMax = 57
	// traversedMax = 87
	// 32, 45, 57, 24, 19, 0, 24, 49, 67, 87, 87
	//                             i
	//                                         j

	// if nums[j] < currentMax -> move our i to j -> currentMax = traversedMax
	// if nums[j] > traversedMax -> traversedMax = nums[j]

	i := 0
	currentMax := nums[i]
	traversedMax := nums[i]
	for j := 1; j < len(nums); j++ {
		if nums[j] < currentMax {
			i = j
			currentMax = traversedMax
		}

		if nums[j] > traversedMax {
			traversedMax = nums[j]
		}
	}

	return i + 1
}

func TestPartitionDistinct(t *testing.T) {
	if result := partitionDisjoint([]int{5, 0, 3, 8, 6}); result != 3 {
		t.Fatalf("Expected 3 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{1, 1, 1, 0, 6, 12}); result != 4 {
		t.Fatalf("Expected 4 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{1, 2, 3, 4, 5, 6, 7, 8}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{8, 8, 8, 8, 8, 8, 8, 8, 9}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{8, 8, 7, 8, 8, 8, 8, 8, 9}); result != 3 {
		t.Fatalf("Expected 3 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{2, 1, 3}); result != 2 {
		t.Fatalf("Expected 2 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{-100, 100, 150, 300, 400, 500, 600, 700}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{-100, 100, 150, 300, 400, 500, 600, 700, 7, 10000}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{100, 100, 150, 300, 400, 500, 600, 700, 7, 10000}); result != 9 {
		t.Fatalf("Expected 9 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{32, 57, 24, 19, 0, 24, 49, 67, 87, 87}); result != 7 {
		t.Fatalf("Expected 7 but got %d", result)
	}
	println()
	if result := partitionDisjoint([]int{32, 45, 57, 24, 19, 0, 24, 49, 67, 87, 87}); result != 8 {
		t.Fatalf("Expected 7 but got %d", result)
	}
}
