package arrays

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	x := strings.Split(s, "")
	sort.Strings(x)
	return strings.Join(x, "")
}

func groupAnagramsSlow(strs []string) [][]string {
	// There is always at least 1 result
	// and the first anagram group is the
	// first element in strs.
	result := make([][]string, 1)
	result[0] = []string{strs[0]}

	// Stores the sorted anagram group and
	// its index group in the result slice.
	hash := make(map[int]string)
	hash[0] = sortString(strs[0])

	for i := 1; i < len(strs); i++ {
		sortedCurrentString := sortString(strs[i])
		groupFound := false
		j := 0
		for j < len(hash) {
			if sortedCurrentString == hash[j] {
				groupFound = true
				result[j] = append(result[j], strs[i])
				break
			}
			j++
		}
		if !groupFound {
			result = append(result, []string{strs[i]})
			hash[j] = sortedCurrentString
		}
	}

	return result
}

func groupAnagrams(strs []string) [][]string {
	freqs := make(map[[26]byte][]string)

	for _, word := range strs {
		charCounter := [26]byte{}

		for _, char := range word {
			charCounter[char-'a']++
		}

		freqs[charCounter] = append(freqs[charCounter], word)
	}

	result := [][]string{}
	for _, words := range freqs {
		result = append(result, words)
	}

	return result
}
