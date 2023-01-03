package leetcode

func (NL *NewLanguage) Solver(operations []string) int {
	total := 0

	for _, operation := range operations {
		if operation == "X++" || operation == "++X" {
			total += 1
		} else {
			total -= 1
		}
	}

	return total
}
