package leetcode

func hashMapper(jewels string) map[byte]int {
	hashMap := make(map[byte]int, len(jewels))
	for _, j := range jewels {
		hashMap[byte(j)] = 1
	}

	return hashMap
}

func (J *Jewel) Solver(jewels string, stones string) int {
	hashMap := hashMapper(jewels)

	total := 0

	for _, val := range stones {
		if v, exist := hashMap[byte(val)]; exist {
			total += v
		}
	}

	return total
}
