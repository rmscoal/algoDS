package leetcode

func (PDB *PartitioningDeciBinary) Solver(n string) int {
	max := 0
	if len(n) == 0 {
		return max
	}

	for _, c := range n {
		// c - '0' will convert c (rune) to int holding back
		// to original value in the string.
		max = getMax(max, int(c-'0'))
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (PDB *PartitioningDeciBinary) FastSolver(n string) int {
	var max byte = 0

	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}

	return int(max - '0')
}
