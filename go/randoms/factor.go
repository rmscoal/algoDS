package randoms

import "sort"

func pthFactor(num, p int) int {
	factors := []int{}

	for i := 1; i < 15/2; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			factors = append(factors, num/i)
		}
	}

	sort.Ints(factors)

	for _, factor := range factors {
		p--

		if p == 0 {
			return factor
		}
	}

	return num
}
