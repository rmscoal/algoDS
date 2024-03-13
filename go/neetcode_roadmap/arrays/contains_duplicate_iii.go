package arrays

import (
	"sort"
)

func removeValue(dup []int, value int) []int {
	for i := 0; i < len(dup); i++ {
		if dup[i] == value {
			dup = append(dup[:i], dup[i+1:]...)
			return dup
		}
	}
	return dup
}

func insertValue(dup []int, value int) ([]int, int) {
	for i := 0; i < len(dup)-1; i++ {
		if value <= dup[i] {
			dup = append(dup[:i], append([]int{value}, dup[i:]...)...)
			return dup, i
		}
	}

	dup = append(dup, value)
	return dup, len(dup) - 1
}

// Hard
func containsNearbyAlmostDuplicateSlow(nums []int, indexDiff int, valueDiff int) bool {
	if indexDiff == 0 {
		return false
	}

	lengthOfDup := findMin(len(nums), indexDiff+1)
	dup := make([]int, findMin(len(nums), lengthOfDup))
	for i := 0; i < lengthOfDup; i++ {
		dup[i] = nums[i]
	}
	sort.Ints(dup)

	for i := 0; i < lengthOfDup-1; i++ {
		if dup[i+1]-dup[i] <= valueDiff {
			return true
		}
	}

	if lengthOfDup >= len(nums) {
		return false
	}

	dup = removeValue(dup, nums[0])
	left := 1
	right := indexDiff + 1
	var idx int

	for right < len(nums) {
		dup, idx = insertValue(dup, nums[right])
		if idx+1 < len(dup) {
			if dup[idx+1]-dup[idx] <= valueDiff {
				return true
			}
		}
		if idx-1 >= 0 {
			if dup[idx]-dup[idx-1] <= valueDiff {
				return true
			}
		}

		dup = removeValue(dup, nums[left])
		right++
		left++
	}

	return false
}

func containsNearbyAlmostDuplicateSolution(nums []int, indexDiff, valueDiff int) bool {
	buckets := make(map[int]int, 0)

	for i, n := range nums {
		bucketNo := n / (valueDiff + 1)

		if n < 0 {
			bucketNo -= 1
		}

		if _, ok := buckets[bucketNo]; ok {
			return true
		}

		buckets[bucketNo] = n

		if prev, ok := buckets[bucketNo-1]; ok {
			if (n - prev) <= valueDiff {
				return true
			}
		}

		if next, ok := buckets[bucketNo+1]; ok {
			if (next - n) <= valueDiff {
				return true
			}
		}

		if len(buckets) > indexDiff {
			removeBucketNo := nums[i-indexDiff] / (valueDiff + 1)
			if nums[i-indexDiff] < 0 {
				removeBucketNo -= 1
			}

			delete(buckets, removeBucketNo)
		}

	}

	return false
}
