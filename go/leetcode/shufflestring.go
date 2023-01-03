package leetcode

// Assumptions:
// 1. len of string and indeces are the same
// 2. All elements in indices are unique.
//
// This solution is inefficient and uses a lot of memory.
func (SS *ShuffleString) Solver(s string, indices []int) string {
	// hashMap generates the permutation according to the
	// indices given.
	hashMap := make(map[int]int)
	for i := 0; i < len(indices); i++ {
		for index, v := range indices {
			if v == i {
				hashMap[i] = index
			}
		}
	}

	var str string
	for i := 0; i < len(indices); i++ {
		str += string(s[hashMap[i]])
	}

	return str
}

func (SS *ShuffleString) FastSolver(s string, indices []int) string {
	resultByte := make([]byte, len(indices))
	for i := 0; i < len(indices); i++ {
		resultByte[indices[i]] = s[i]
	}

	return string(resultByte)
}
