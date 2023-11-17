package arrays

import "testing"

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	if result := containsNearbyAlmostDuplicateSlow([]int{1, 14, 23, 45, 56, 2, 3}, 1, 10); result != true {
		t.Fatalf("Expected is %t but got %t\n", true, result)
	}
	if result := containsNearbyAlmostDuplicateSlow([]int{-3, 3}, 2, 4); result != false {
		t.Fatalf("Expected is %t but got %t\n", false, result)
	}
	if result := containsNearbyAlmostDuplicateSlow([]int{-3, 3, 9, 15, 21}, 5, 4); result != false {
		t.Fatalf("Expected is %t but got %t\n", false, result)
	}

	if result := containsNearbyAlmostDuplicateSolution([]int{1, 14, 23, 45, 56, 2, 3}, 1, 10); result != true {
		t.Fatalf("Expected is %t but got %t\n", true, result)
	}
	if result := containsNearbyAlmostDuplicateSolution([]int{-3, 3}, 2, 4); result != false {
		t.Fatalf("Expected is %t but got %t\n", false, result)
	}
	if result := containsNearbyAlmostDuplicateSolution([]int{-3, 3, 9, 15, 21}, 5, 4); result != false {
		t.Fatalf("Expected is %t but got %t\n", false, result)
	}
}
