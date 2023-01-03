package leetcode

func (CS *ClimbingStairs) Solver(n int) int {

	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}

	return CS.Solver(n-1) + CS.Solver(n-2)
}

func (CS *ClimbingStairs) FastSolver(n int) int {
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1

	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n]
}
