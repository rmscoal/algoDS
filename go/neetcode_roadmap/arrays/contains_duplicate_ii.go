package arrays

// https://leetcode.com/problems/contains-duplicate-ii/description/
func containsNearbyDuplicate(nums []int, k int) bool {
	// Time: O(n)
	// Space: O(n)

	hash := make(map[int]int)

	for currIdx, value := range nums {
		if prevIdx, found := hash[value]; found {
			if currIdx-prevIdx <= k {
				return true
			}
		}

		hash[value] = currIdx
	}

	return false
}
