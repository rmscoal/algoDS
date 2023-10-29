package arrays

// Solution to
// https://www.hackerrank.com/challenges/crush/problem
// Too slow
func arrayManipulation(n int32, queries [][]int32) int64 {
	var max int64 = 0

	arr := make([]int64, n, n)

	for _, query := range queries {
		k := query[2]
		for i := query[0] - 1; i < query[1]; i++ {
			arr[i] += int64(k)

			if arr[i] > max {
				max = arr[i]
			}
		}
	}

	return max
}

// Solution to
// https://www.hackerrank.com/challenges/crush/problem
func arrayManipulationWow(n int32, queries [][]int32) int64 {
	var max int64
	arr := make([]int64, n+1)

	for _, q := range queries {
		start := q[0] - 1
		end := q[1]
		k := q[2]

		arr[start] += int64(k)
		arr[end] -= int64(k)
	}

	var tmp int64

	for _, v := range arr {
		tmp += v
		if tmp > max {
			max = tmp
		}
	}

	return max
}
