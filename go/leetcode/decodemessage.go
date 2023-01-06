package leetcode

import (
	"strings"
)

/*
You are given the strings key and message,
which represent a cipher key and a secret
message, respectively. The steps to decode
message are as follows:

Use the first appearance of all 26 lowercase
English letters in key as the order of the
substitution table. Align the substitution
table with the regular English alphabet.
Each letter in message is then substituted
using the table. Spaces ' ' are transformed
to themselves.

Input: key = "the quick brown fox jumps over
the lazy dog", message = "vkbs bs t suepuv"
Output: "this is a secret"
Explanation: The diagram above shows the
substitution table. It is obtained by taking
the first appearance of each letter in "the
quick brown fox jumps over the lazy dog".

Assumption:
The input key contains all 26 English lowercase
letters.
*/

// Done without assumptions
func (s *DecodeMessage) Solver(key, message string) string {
	var str string = ""
	var hashMap map[byte]byte = make(map[byte]byte)

	// This for loop will fail if input key does not
	// follow the assumption.
	for i, j := 0, 97; i < len(key) || j <= 'z'; i++ {
		if key[i] != ' ' {
			if _, exist := hashMap[key[i]]; !exist {
				hashMap[key[i]] = byte(j)
				j++
			}
		}
	}

	// Not needed due to assumptions
	// for j := 97; j <= 'z'; j++ {
	// 	if _, exist := hashMap[byte(j)]; !exist {
	// 		hashMap[byte(j)] = byte(j)
	// 	}
	// }

	for _, r := range message {
		if byte(r) == ' ' {
			str += string(r)
		} else {
			str += string(hashMap[byte(r)])
		}
	}

	return str
}

func (s *DecodeMessage) FastSolver(key, message string) string {
	m := make(map[rune]byte)
	i := 0
	for _, c := range key {
		if c == ' ' {
			continue
		}
		if _, ok := m[c]; !ok {
			m[c] = 'a' + byte(i)
			i++
		}
	}
	builder := strings.Builder{}
	for _, c := range message {
		if c == ' ' {
			builder.WriteRune(c)
		} else {
			builder.WriteByte(m[c])
		}
	}
	return builder.String()
}
