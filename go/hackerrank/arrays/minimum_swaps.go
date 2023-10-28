package arrays

// Solution for
// https://www.hackerrank.com/challenges/minimum-swaps-2/problem
func minimumSwaps(arr []int32) int32 {
	var count int32 = 0

	for i := 1; i <= len(arr); i++ {
		if arr[i-1] != int32(i) {
			for j := i + 1; j <= len(arr); j++ {
				if arr[j-1] == int32(i) {
					arr[i-1], arr[j-1] = arr[j-1], arr[i-1]
					count++
					break
				}
			}
		}
	}

	return count
}

// Cases?

// [7, 1, 3, 2, 4, 5, 6]
// -
// [1, 7, 3, 2, 4, 5, 6]
// [1, 2, 3, 7, 4, 5, 6]
// [1, 2, 3, 4, 7, 5, 6]
// [1, 2, 3, 4, 5, 7, 6]
// [1, 2, 3, 4, 5, 6, 7]

// [4, 3, 1, 2]
// -
// [1, 3, 4, 2]
// [1, 2, 4, 3]
// [1, 2, 3, 4]

// [2, 3, 4, 1, 5]
// -
// [1, 3, 4, 2, 5]
// [1, 2, 4, 3, 5]
// [1, 2, 3, 4, 5]

// [1, 3, 5, 2, 4, 6, 7]
// -
// [1, 2, 5, 3, 4, 6, 7]
// [1, 2, 3, 5, 4, 6, 7]
// [1, 2, 3, 4, 5, 6, 7]
