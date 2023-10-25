package string

import (
	"strings"
)

/*
Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
*/

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] == "" {
			arr = append(arr[:i], arr[i+1:]...)
			i--
			n--
		}
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return strings.Join(arr, " ")
}
