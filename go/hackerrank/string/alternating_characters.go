package string

func alternatingCharacters(s string) int32 {
	deletionCount := 0
	comparisonIndex := 0

	for i := 1; i < len(s); i++ {
		if s[comparisonIndex] == s[i] {
			deletionCount++
		} else {
			comparisonIndex = i
		}
	}

	return int32(deletionCount)
}
