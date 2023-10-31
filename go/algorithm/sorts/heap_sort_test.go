package sorts

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{4, 3, 2, 5, 1, 9}

	heapsort(arr)

	if !isSorted(arr) {
		t.Fatalf("Not sorted %v\n", arr)
	}
}
