package leetcode

func (LL *LengthLastWord) Solver(s string) int {
	// Variable to hold the start of the last string
	var j int = -1
	var i int

	// Start iterating from the last of the string
	for i = len(s) - 1; 0 <= i; i-- {
		if byte(s[i]) != ' ' {
			if j == -1 {
				j = i
			} else {
				continue
			}
		} else {
			if j == -1 {
				continue
			} else {
				break
			}
		}
	}

	if j == len(s)-1 {
		return len(s[i+1:])
	}
	return len(s[i+1 : j+1])
}
