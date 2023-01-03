package leetcode

func (V *ValidParenthesis) Solver(s string) bool {
	pair := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Variable to hold to openings. At the end of the loop
	// this list should be empty.
	var list []int

	// Since open bracket must be close, the length of the
	// string must be even.
	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if value, close := pair[s[i]]; close {
			// This block is for a close parenthesis
			//
			// During a close parenthesis, checks whether there
			// was an opening at first.
			if len(list) == 0 {
				return false
			}

			// Checks whether it had the correct opening.
			if value != s[list[len(list)-1]] {
				return false
			} else {
				// Pop the last element
				list = list[:len(list)-1]
			}
		} else {
			// This block is for an open parenthesis
			//
			// Add an opening.
			list = append(list, i)
		}
	}

	// Checks whether all open parenthesis has been closed.
	return len(list) == 0
}
