package leetcode

import (
	"strconv"
	"strings"
)

func (ES *ExcelSheet) WrongSolver(s string) []string {
	// Constraints:
	// 1. s.length == 5
	// Column is from A - Z
	// Row is from 1 - 9

	res := make([]string, 0)

	split := strings.Split(s, ":")
	org, dest := split[0], split[1]
	splitOrg := strings.Split(org, "")
	orgCol, orgRow := splitOrg[0], splitOrg[1]
	splitDest := strings.Split(dest, "")
	destCol, destRow := splitDest[0], splitDest[1]

	orgRowInt, _ := strconv.Atoi(orgRow)
	destRowInt, _ := strconv.Atoi(destRow)

	var j int = orgRowInt

	if []byte(orgCol)[0] == []byte(destCol)[0] {
		for ; j <= destRowInt; j++ {
			temp := orgCol + strconv.Itoa(j)
			res = append(res, temp)
		}
	} else {
		for i := []byte(orgCol)[0]; i <= []byte(destCol)[0]; i++ {
			if i == []byte(destCol)[0] {
				for ; j <= destRowInt; j++ {
					temp := string(i) + strconv.Itoa(j)
					res = append(res, temp)
				}
				break
			} else {
				for ; j <= 9; j++ {
					temp := string(i) + strconv.Itoa(j)
					res = append(res, temp)
				}
			}
			j = 1
		}
	}

	return res
}

func (ES *ExcelSheet) Solver(s string) []string {
	// Constraints:
	// 1. s.length == 5
	// Column is from A - Z
	// Row is from 1 - 9

	res := make([]string, 0)

	// Splitting process
	split := strings.SplitN(s, ":", 2)
	orgCol, orgRow := split[0][0], split[0][1]
	destCol, destRow := split[1][0], split[1][1]

	for i := orgCol; i <= destCol; i++ {
		for j := orgRow; j <= destRow; j++ {
			temp := string(i) + string(j)
			res = append(res, temp)
		}
	}

	return res
}
