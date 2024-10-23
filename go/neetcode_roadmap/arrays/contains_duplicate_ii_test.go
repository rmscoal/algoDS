package arrays

import "testing"

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	// Time: O(n)
// 	// Space: O(n)

// 	hash := make(map[int]int)

// 	for currIdx, value := range nums {
// 		if prevIdx, found := hash[value]; found {
// 			if currIdx-prevIdx <= k {
// 				return true
// 			}
// 		}

// 		hash[value] = currIdx
// 	}

// 	return false
// }

func containsNearbyDuplicate(nums []int, k int) bool {
	i, j := 0, 1
	maps := map[int]int{nums[i]: 1}

	for j < len(nums) {
		count, found := maps[nums[j]]
		if !found {
			maps[nums[j]] = 1
		} else if count > 0 {
			return true
		} else {
			maps[nums[j]]++
		}

		j++
		if j-i > k {
			maps[nums[i]]--
			i++
		}
	}

	return false
}

func TestContainsNearbyDuplicate(t *testing.T) {
	if result := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3); result != true {
		t.Fatalf("Expected is %t but got %t\n", true, result)
	}

	if result := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2); result != false {
		t.Fatalf("Expected is %t but got %t\n", false, result)
	}
}
