package arrays

import "testing"

func TestProductExceptSelf(t *testing.T) {
	if result := productExceptSelfExtraMem([]int{1, 2, 3, 4}); !assertArray[int]([]int{24, 12, 8, 6}, result) {
		t.Fatalf("Wrong bro")
	}
	if result := productExceptSelf([]int{1, 2, 3, 4}); !assertArray[int]([]int{24, 12, 8, 6}, result) {
		t.Fatalf("Wrong bro")
	}
}
