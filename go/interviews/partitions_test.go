package interviews

import (
	"testing"
)

// Partition Question
//
// Given a number `target`. How many ways can you partition it such that, no
// each part are greater than k.
// Example:
// total = 8
// k = 2
//
// - [1,1,1,1,1,1,1,1]
// - [1,1,1,1,1,1,2]
// - [1,1,1,1,2,2]
// - [1,1,2,2,2]
// - [2,2,2,2]

// Solution
// Let's see some pattern => p_k(n) = p_k(n-k) + p_k-1(n) where p_k(n), k >= n = p(n) = n
// target = 1, k = 1 -> 1

// target = 2, k = 1 -> 1 (trivial)
// target = 2, k = 2 -> 2
//		1,1
//		2

// target = 3, k = 1 -> 1 (trivial)
// target = 3, k = 2 -> 2 => [3,1] + [1,2]
//		1,1,1
//		1,2
// target = 3, k = 3 -> 3
//		1,1,1
//		1,2
//		3

// target = 4, k = 1 -> 1 (trivial)
// target = 4, k = 2 -> 3 => [4,1] + [2,2]
//		1,1,1,1
//		1,1,2
//		2,2
// target = 4, k = 3 -> 4 => [4,2] + [1,3]
//		1,1,1,1
//		1,1,2
//		1,3
//		2,2
// target = 4, k = 4 -> 5 => [4,3] + [1,4]
//		1,1,1,1
//		1,1,2
//		1,3
//		2,2
//		4

// target = 5, k = 1 -> 1 (trivial)
// target = 5, k = 2 -> 3 => [5,1] + [3,2]
//
//	1,1,1,1,1
//	1,1,1,2
//	1,2,2
//
// target = 5, k = 3 -> 5 => [5,2] + [2,3]
//
//	1,1,1,1,1
//	1,1,1,2
//	1,2,2
//	1,1,3
//	2,3
//
// target = 5, k = 4 -> 5 => [5,4] + [1,3]
//
//	1,1,1,1,1
//	1,1,1,2
//	1,2,2
//	1,1,3
//	2,3
//	1,4
//
// target = 5, k = 5 -> 5
//
//	1,1,1,1,1
//	1,1,1,2
//	1,2,2
//	1,1,3
//	2,3
//	1,4
//	5
func partitions_recursive(target, k int) int {
	if k == 1 {
		return 1
	} else if target <= k {
		return target
	} else {
		return partitions_recursive(target-k, k) + partitions_recursive(target, k-1)
	}
}

func TestPartitions(t *testing.T) {
	if result := partitions_recursive(8, 2); result != 5 {
		t.Errorf("Expected 5 but got %d", result)
	}
	if result := partitions_recursive(5, 3); result != 5 {
		t.Errorf("Expected 5 but got %d", result)
	}
	if result := partitions_recursive(2, 1); result != 1 {
		t.Errorf("Expected 1 but got %d", result)
	}
	if result := partitions_recursive(4, 3); result != 4 {
		t.Errorf("Expected 4 but got %d", result)
	}
}

func BenchmarkPartitions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partitions_recursive(100, 10)
	}
}
