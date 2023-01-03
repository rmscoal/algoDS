package leetcode

import "strings"

func (LS *LongestSentence) Solver(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		arr := strings.Split(sentence, " ")
		if len(arr) > max {
			max = len(arr)
		}
	}
	return max
}
