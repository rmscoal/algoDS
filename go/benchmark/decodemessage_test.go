package benchmark

import (
	"fmt"
	"testing"

	"algoDS/go/leetcode"
)

var decodeMessageTestCases = []struct {
	input [2]string
}{
	{input: [2]string{"the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"}},
	{input: [2]string{"eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"}},
}

func BenchmarkDecodeMessageSolver(b *testing.B) {
	DM := &leetcode.DecodeMessage{}
	for i, v := range decodeMessageTestCases {
		b.Run(fmt.Sprintf("test-case_%d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DM.Solver(v.input[0], v.input[1])
			}
		})
	}
}

func BenchmarkDecodeMessageFastSolver(b *testing.B) {
	DM := &leetcode.DecodeMessage{}
	for i, v := range decodeMessageTestCases {
		b.Run(fmt.Sprintf("test-case_%d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DM.FastSolver(v.input[0], v.input[1])
			}
		})
	}
}
