package arrays

import (
	"fmt"
	"sort"
	"testing"
)

func threesum(nums []int) [][]int {
	sort.Ints(nums)
	answer := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target, left, right := -nums[i], i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				answer = append(answer, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	}

	return answer
}

func TestThreeSum(t *testing.T) {
	fmt.Printf("%+v\n", threesum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("%+v\n", threesum([]int{0, 0, 0}))
	fmt.Printf("%+v\n", threesum([]int{-2, 0, 0, 2, 2}))
	fmt.Printf("%+v\n", threesum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
