package leetcode

import (
	"strconv"
)

func (A *AddBinary) Solver(a string, b string) string {
	int1, _ := strconv.ParseInt(a, 2, 64)
	int2, _ := strconv.ParseInt(b, 2, 64)

	num := int64(int1 + int2)
	return strconv.FormatInt(num, 2)
}

func (A *AddBinary) FastSolver(a string, b string) string {
	var p1 int = len(a) - 1
	var p2 int = len(b) - 1
	var result string
	var carry int

	for p1 >= 0 || p2 >= 0 {
		var i1 int
		var i2 int
		var sum int = carry
		if p1 >= 0 {
			i1 = int(a[p1] - '0')
		}
		if p2 >= 0 {
			i2 = int(b[p2] - '0')
		}
		sum += i1 + i2
		switch sum {
		case 0:
			result = "0" + result
			carry = 0
		case 1:
			result = "1" + result
			carry = 0
		case 2:
			result = "0" + result
			carry = 1
		case 3:
			result = "1" + result
			carry = 1
		}
		p1 -= 1
		p2 -= 1
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}
