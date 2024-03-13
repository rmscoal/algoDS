package string

import (
	"strconv"
	"strings"
)

func findBiggestLength(a, b string) int {
	if len(a) > len(b) {
		return len(a)
	} else {
		return len(b)
	}
}

func findAdditionalZeros(s string, length int) int {
	// s = "3" and length is 2 we would want to return 1
	// because we want to multiply "3" by 10 1 time.
	return length - len(s)
}

// This is cool, but it doesn't solves our problem
func largestNumberFail(nums []int) string {
	/*
		Given that we are comparing
		30 and 3...

		We would find the max length of each of them.
		So with 30 and 3... The bigger number is obviously 30
		we would want to do comparing 30 and 30.
		Then we would want to compare the length.
		The ones that are shorter will be sorted first.

		Fail case:
		111300, 111311
	*/

	strArray := make([]string, 0)

	for _, v := range nums {
		strArray = append(strArray, strconv.Itoa(v))
	}

	// Begin sorting here
	// Here let's just do bubble sort
	for i := 0; i < len(strArray); i++ {
		swapped := false

		for j := 0; j < len(strArray)-1; j++ {
			switched := false
			biggest := findBiggestLength(strArray[j], strArray[j+1])
			num1, _ := strconv.Atoi(strArray[j] + strings.Repeat("0", findAdditionalZeros(strArray[j], biggest)))
			num2, _ := strconv.Atoi(strArray[j+1] + strings.Repeat("0", findAdditionalZeros(strArray[j+1], biggest)))
			if num1 < num2 {
				switched = true
			} else if num1 == num2 && len(strArray[j]) > len(strArray[j+1]) {
				switched = true
			}

			if switched {
				strArray[j], strArray[j+1] = strArray[j+1], strArray[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return strings.Join(strArray, "")
}

func largestNumber(nums []int) string {
	// Say that we are comparing [3, 5]
	// choice1 := "35"
	// choice2 := "53"

	// Hey if choice2 > choice1  {
	// then we swap
	// }
	strArray := make([]string, 0)
	for _, v := range nums {
		strArray = append(strArray, strconv.Itoa(v))
	}

	for i := 0; i < len(strArray); i++ {
		swapped := false

		for j := 0; j < len(strArray)-1; j++ {
			choice1, _ := strconv.ParseUint(strArray[j]+strArray[j+1], 10, 64)
			choice2, _ := strconv.ParseUint(strArray[j+1]+strArray[j], 10, 64)

			if choice2 > choice1 {
				strArray[j], strArray[j+1] = strArray[j+1], strArray[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	result := strings.Join(strArray, "")
	resultAsUint, _ := strconv.ParseUint(result, 10, 64)

	if resultAsUint != 0 {
		return result
	} else {
		return "0"
	}
}
