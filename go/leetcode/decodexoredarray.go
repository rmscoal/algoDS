package leetcode

/*
There is a hidden integer array arr that
consists of nnon-negative integers.

It was encoded into another integer array
encoded of length n - 1, such that encoded[i]
= arr[i] XOR arr[i + 1]. For example, if
arr = [1,0,2,1], then encoded = [1,2,3].

You are given the encoded array. You are also
given an integer first, that is the first element
of arr, i.e. arr[0].

Return the original array arr. It can be proved
that the answer exists and is unique.

Time-complexity: O(n)
Space-complexity: O(n)
*/
func (s *DecodeXORedArray) Solver(encoded []int, first int) []int {
	expected_input := make([]int, len(encoded)+1)
	expected_input[0] = first

	for i := 0; i < len(encoded); i++ {
		expected_input[i+1] = encoded[i] ^ expected_input[i]
	}

	return expected_input
}
