package benchmark

import (
	"fmt"
	"testing"

	"algoDS/go/leetcode"
)

var decodeXORedarrayTestCases = []struct {
	encoded []int
	first   int
}{
	{encoded: []int{1, 2, 3}, first: 1},
	{encoded: []int{6, 2, 7, 3}, first: 4},
}

// Use -bench=. to run all benchmark
//
// To avoid executing any test functions in the test files,
// pass a regular expression to the -run flag
//
// Preferred commands:
// go test -bench=DecodeXOR -count 5 -benchtime=100x -benchmem
// go test -bench=DecodeXOR -count 5 -benchtime=10s -benchmem

func BenchmarkDecodeXORedArray(b *testing.B) {
	C := &leetcode.DecodeXORedArray{}
	for i, v := range decodeXORedarrayTestCases {
		b.Run(fmt.Sprintf("Test Case No %d", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				C.Solver(v.encoded, v.first)
			}
		})
	}
}
