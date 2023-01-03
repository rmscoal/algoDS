package leetcode

import "strings"

func (GP *GoalParser) Solver(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	command = strings.ReplaceAll(command, "()", "o")
	return command
}
