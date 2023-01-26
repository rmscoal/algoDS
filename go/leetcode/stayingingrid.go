package leetcode

/*
Problem: There is an n x n grid, with the top-left cell at (0, 0) and
the bottom-right cell at (n - 1, n - 1). You are given the
integer n and an integer array startPos where

startPos = [startrow, startcol]

indicates that a robot is initially at cell (startrow, startcol).

You are also given a 0-indexed string s of length m where
s[i] is the ith instruction for the robot: 'L' (move left),
'R' (move right), 'U' (move up), and 'D' (move down).

The robot can begin executing from any ith instruction in s.
It executes the instructions one by one towards the end of s
but it stops if either of these conditions is met:

The next instruction will move the robot off the grid.
There are no more instructions left to execute.
Return an array answer of length m where answer[i] is the
number of instructions the robot can execute if the robot
begins executing from the ith instruction in s.

Example:
Input: n = 3, startPos = [0,1], s = "RRDDLU"
Output: [1,5,4,3,1,0]
Explanation: Starting from startPos and beginning execution
from the ith instruction:
- 0th: "RRDDLU". Only one instruction "R" can be executed
before it moves off the grid.
- 1st:  "RDDLU". All five instructions can be executed while
it stays in the grid and ends at (1, 1).
- 2nd:   "DDLU". All four instructions can be executed while
it stays in the grid and ends at (1, 0).
- 3rd:    "DLU". All three instructions can be executed while
it stays in the grid and ends at (0, 0).
- 4th:     "LU". Only one instruction "L" can be executed before
it moves off the grid.
- 5th:      "U". If moving up, it would move off the grid.

Constraints:
m == s.length
1 <= n, m <= 500
startPos.length == 2
0 <= startrow, startcol < n
s consists of 'L', 'R', 'U', and 'D'.

Solver:
Do a double loop to read the input str. The outer loop is to
read all the input. Meanwhile, the input is to read the characters
onwards to calculcate how many instructions can the robot take.
In the inner loop, after each loop, we will alter the robot's
current position according to the instructions. Next, validate
whether the robot is still staying in grid after oding the instruction.
If it doesn't then break the inner loop. If it is, increment answer[i]
value.

Time-complexity: O(n^2)
Space-complexity: O(n)
*/
func (s *StayingInGrid) Solver(n int, startPos []int, str string) []int {
	var answer []int = make([]int, len(str))

	for i := 0; i < len(str); i++ {
		currPos := startPos
		for j := i; j < len(str); j++ {
			switch byte(str[j]) {
			case 'D':
				currPos = []int{currPos[0] + 1, currPos[1]}
			case 'U':
				currPos = []int{currPos[0] - 1, currPos[1]}
			case 'L':
				currPos = []int{currPos[0], currPos[1] - 1}
			case 'R':
				currPos = []int{currPos[0], currPos[1] + 1}
			}

			if currPos[0] > n-1 || currPos[0] < 0 || currPos[1] > n-1 || currPos[1] < 0 {
				break
			}

			answer[i] += 1
		}
	}
	return answer
}
