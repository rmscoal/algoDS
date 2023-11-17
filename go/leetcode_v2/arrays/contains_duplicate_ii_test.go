package arrays

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	if result := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3); result != true {
		t.Fatalf("Expected is %t but got %t\n", true, result)
	}
}
