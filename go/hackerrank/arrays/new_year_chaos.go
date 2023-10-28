package arrays

// Solution to
// https://www.hackerrank.com/challenges/new-year-chaos/problem
func minimumBribes(q []int32) int64 {
	var count int64 = 0

	for i := 0; i < len(q); i++ {
		swapped := false
		chaotic := false
		swapCount := 0

		for j := 0; j < len(q)-1; j++ {
			if q[j] > q[j+1] {
				q[j], q[j+1] = q[j+1], q[j]
				count++
				swapCount++
				swapped = true
			} else {
				swapCount = 0
			}

			if swapCount > 2 {
				chaotic = true
			}
		}

		if chaotic {
			println("Too chaotic")
			return -1
		}

		if !swapped {
			break
		}
	}

	// for i := 0; i < len(q)-1; i++ {
	// 	pos := int(q[i] - 1)
	// 	if pos > i {
	// 		if pos-i > 2 {
	// 			println("Too chaotic")
	// 			return -1
	// 		}

	// 		count += int64(pos - i)
	// 	}
	// }

	return count
}
