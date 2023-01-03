package leetcode

import "strings"

func (IP *IPAddress) Solver(s string) string {
	arr := strings.Split(s, ".")
	return strings.Join(arr, "[.]")
}

func (IP *IPAddress) FastSolver(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
