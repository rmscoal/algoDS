package arrays

import (
	"math"
	"sort"
	"testing"
)

func longestConsecutive_Timeout_Exceeded(nums []int) int {
	// Hashmap

	// hash = {}
	// [100, 4, 200, 1, 3, 2]
	// i = 100
	// 	check if 99 or 101 is in the hash
	// 		hash[100] = 1
	//
	// i = 4
	// 		hash[4] = 1
	//
	// i = 200
	// 		hash[200] = 1
	//
	// i = 1
	// 	check if -1 or 2 is in the hash
	// 		hash[1] = 1
	//
	// i =  3
	// 	check if 2 or 4 is in the hash
	// 		hash[3] = 1
	// 		hash[4]++
	// 			check if 5 is in the hash
	//
	// i = 2
	// 	check if 1 or 3 is in the hash
	// 		hash[2] = 1
	// 		hash[1]++
	// 			check if 0 is in the hash
	// 		hash[3]++
	// 			check if 4 is in the hash
	// 				hash[4]++
	//
	// hence we have
	// hash = {100: 1, 4: 3, 200: 1, 1: 2, 3: 2, 2: 1}

	hash := map[int]int{}

	// Maps our nums to a set
	for i, j := 0, len(nums); i <= j; i, j = i+1, j-1 {
		hash[nums[i]] = 1
		hash[nums[j]] = 1
	}

	// Traverse through our set
	for num := range hash {
		lowerValue := num - 1
		upperValue := num + 1

		// Finding lower values
		for {
			_, found := hash[lowerValue]
			if !found {
				break
			}
			hash[num]++
			lowerValue--
		}

		// Find upper values
		for {
			_, found := hash[upperValue]
			if !found {
				break
			}
			hash[num]++
			upperValue++
		}
	}

	max := 0

	for _, value := range hash {
		if max < value {
			max = value
		}
	}

	return max
}

func TestLongestConsecutiveTimeoutExceeded(t *testing.T) {
	if result := longestConsecutive_Timeout_Exceeded([]int{100, 4, 200, 1, 3, 2}); result != 4 {
		t.Fatalf("Expected 4 but got %d", result)
	}
	if result := longestConsecutive_Timeout_Exceeded([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}); result != 9 {
		t.Fatalf("Expected 9 but got %d", result)
	}
}

func longestConsecutive_Sort_nlogn(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	sort.Ints(nums)

	max := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i]-1 == nums[i-1]:
			count++
		case nums[i] == nums[i-1]:
		default:
			count = 1
		}

		if max < count {
			max = count
		}
	}

	return max
}

func TestLongestConsecutiveSort(t *testing.T) {
	if result := longestConsecutive_Sort_nlogn([]int{100, 4, 200, 1, 3, 2}); result != 4 {
		t.Fatalf("Expected 4 but got %d", result)
	}
	if result := longestConsecutive_Sort_nlogn([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}); result != 9 {
		t.Fatalf("Expected 9 but got %d", result)
	}
	if result := longestConsecutive_Sort_nlogn([]int{0}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
}

func longestConsecutiveHashTable_Timeout(nums []int) int {
	set := make(map[int]int, 0)
	min, max := math.MaxInt, math.MinInt

	// Converting our nums into a set while also finding
	// our minimum and maximum value because we want to
	// iterate by sequence.
	for _, num := range nums {
		set[num] = 1
		if min > num {
			min = num
		}

		if max < num {
			max = num
		}
	}

	// If out set length is less than or equals to 1,
	// we just return the length of our set.
	if len(set) <= 1 {
		return len(set)
	}

	// Make a sequence group always starting from the
	// mininum number of a sequence.
	seq := make(map[int]int, 0)
	seq[min] = 1
	previousGroup := min
	hasGroup := true
	maxLen := 0
	for i := min + 1; i <= max; i++ {
		// Check if i is in our set
		_, found := set[i]

		switch {
		// In the case where i is in the set and the chain
		// has not cut off, we add the count to our previous
		// group.
		case found && hasGroup:
			seq[previousGroup]++
			// In the case where i is in the set but the previous
			// chain has been cut off, we are going to make a new
			// sequence starting from i.
		case found && !hasGroup:
			seq[i] = 1
			previousGroup = i
			hasGroup = true
			// In the case where i is not found on our set, then
			// we have to cut off the chain.
		default:
			hasGroup = false
		}

		// Here we also find the greatest sequence length
		if maxLen < seq[previousGroup] {
			maxLen = seq[previousGroup]
		}
	}

	return maxLen
}

func TestLongestConsecutiveHash(t *testing.T) {
	if result := longestConsecutiveHashTable_Timeout([]int{100, 4, 200, 1, 3, 2}); result != 4 {
		t.Fatalf("Expected 4 but got %d", result)
	}
	if result := longestConsecutiveHashTable_Timeout([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}); result != 9 {
		t.Fatalf("Expected 9 but got %d", result)
	}
	if result := longestConsecutiveHashTable_Timeout([]int{0}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	if result := longestConsecutiveHashTable_Timeout([]int{0, 0}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	if result := longestConsecutiveHashTable_Timeout([]int{}); result != 0 {
		t.Fatalf("Expected 0 but got %d", result)
	}
}

/* FINAL */

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, 0)

	for _, num := range nums {
		set[num] = true
	}

	if len(set) <= 1 {
		return len(set)
	}

	group := make(map[int]int, 0)
	maxLen := 0
	for num := range set {
		// If num is not the start of a sequence
		// we want to just continue to the next num.
		if _, found := set[num-1]; found {
			continue
		}

		// Hence here, num is the start of a sequence
		group[num] = 1
		// We want to find the consecutive sequence
		nextSeq := num + 1
		for {
			if _, found := set[nextSeq]; found {
				group[num]++
				nextSeq++
			} else {
				break
			}
		}

		if maxLen < group[num] {
			maxLen = group[num]
		}
	}

	return maxLen
}

func TestLongestConsecutive(t *testing.T) {
	if result := longestConsecutive([]int{100, 4, 200, 1, 3, 2}); result != 4 {
		t.Fatalf("Expected 4 but got %d", result)
	}
	if result := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}); result != 9 {
		t.Fatalf("Expected 9 but got %d", result)
	}
	if result := longestConsecutive([]int{0}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	if result := longestConsecutive([]int{0, 0}); result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
	if result := longestConsecutive([]int{}); result != 0 {
		t.Fatalf("Expected 0 but got %d", result)
	}
}
