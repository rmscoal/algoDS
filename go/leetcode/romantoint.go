package leetcode

func (R *RomanToInt) Solver(s string) int {
	type mapper struct {
		value  int
		before byte
	}
	dict := map[byte]mapper{
		'I': {1, ' '},
		'V': {5, 'I'},
		'X': {10, 'I'},
		'L': {50, 'X'},
		'C': {100, 'X'},
		'D': {500, 'C'},
		'M': {1000, 'C'},
	}

	var total int

	for index, char := range s {
		total += dict[byte(char)].value
		if index != 0 && s[index-1] == dict[byte(char)].before {
			total -= dict[byte(s[index-1])].value * 2
		}
	}

	return total
}

func (R *RomanToInt) FastSolver(s string) int {
	symbolToValue := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var total int

	for i := 0; i < len(s)-1; i++ {
		if symbolToValue[s[i]] < symbolToValue[s[i+1]] {
			total -= symbolToValue[s[i]]
		} else {
			total += symbolToValue[s[i]]
		}
	}

	total += symbolToValue[s[len(s)-1]]
	return total
}
