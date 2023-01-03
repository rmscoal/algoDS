package leetcode

func (PO *PlusOne) FastSolver(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i] + 1
		if digit == 10 {
			digits[i] = 0
		} else {
			digits[i] = digit
			return digits
		}
	}

	// This is for the int such as 999...
	newDigits := []int{}
	newDigits = append(newDigits, 1)
	newDigits = append(newDigits, digits...) // Destructure digits
	return newDigits
}
