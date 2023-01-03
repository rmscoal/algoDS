package leetcode

func (M *MoveBallToBox) Solver(boxes string) []int {
	res := make([]int, len(boxes))

	for i, j := 0, len(boxes)-1; i < len(boxes); i, j = i+1, j-1 {
		for k := 0; k < i; k++ {
			if boxes[k] == '1' {
				res[i] += i - k
			}
		}

		for k := len(boxes) - 1; k > j; k-- {
			if boxes[k] == '1' {
				res[j] += k - j
			}
		}
	}

	return res
}

// Brilliant
func (M *MoveBallToBox) FastSolver(boxes string) []int {
	minOps, left, right := 0, 0, 0

	for i, box := range boxes {
		if box == '1' {
			minOps += i
			right++
		}
	}

	res := make([]int, len(boxes))
	for i, box := range boxes {
		res[i] = minOps
		right -= int(box - '0')
		left += int(box - '0')
		minOps -= right
		minOps += left
	}

	return res
}
