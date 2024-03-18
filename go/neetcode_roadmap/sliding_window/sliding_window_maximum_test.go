package slidingwindow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxSlidingWindow_Timeout(nums []int, k int) []int {
	var results []int
	// We maintain a chain whose first element is always the last element in the window
	// and the last element in the chain is the biggest element in the window and each
	// element in the chain is bigger than the previous element in the chain.
	chain := []int{nums[k-1]}

	// Example 1:
	// 2, 1, 4, 5, 3, 4, 1, 2  | k = 4
	// window = 2, 1, 4, 5
	// chain = [5] -> chain is built

	// window = 1, 4, 5, 3
	// chain = [3, 5]

	// window = 4, 5, 3, 4
	// chain = [4, 5]

	// window = 5, 3, 4, 1
	// chain = [1, 4, 5]

	// window = 3, 4, 1, 2
	// chain = [1, 4]

	// Example 2:
	// 1, 3, -1, -3, 5, 3, 6, 7  | k = 3
	// window = 1, 3, -1
	// chain = [-1, 3]
	//
	// window = 3, -1, -3
	// chain = [-1, 3]
	//
	// window = -1, -3, 5
	// chain = [5]
	//
	// window = -3, 5, 3
	// chain = [3, 5]
	//
	// window = 5, 3, 6
	// chain = [6]

	for i := k - 2; i >= 0; i-- {
		if nums[i] >= chain[len(chain)-1] {
			chain = append(chain, nums[i])
		}
	}

	fmt.Printf("init chain %+v\n", chain)
	results = append(results, chain[len(chain)-1])
	left := 0

	for right := k; right < len(nums); right++ {
		if chain[len(chain)-1] == nums[left] {
			chain = chain[:len(chain)-1]
		}
		left++

		for len(chain) > 0 {
			if chain[0] < nums[right] {
				chain = chain[1:]
			} else {
				break
			}
		}
		fmt.Printf("after removal chain %+v\n", chain)
		chain = append([]int{nums[right]}, chain...) // The bottleneck lies in here...
		fmt.Printf("after add chain %+v\n", chain)
		println()
		results = append(results, chain[len(chain)-1])
	}

	return results
}

func maxSlidingWindow(nums []int, k int) []int {
	var results []int
	chain := make([]int, 0, k) // Here instead we use monotonic queue

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			// Remove those smaller than our new element
			for len(chain) > 0 && chain[len(chain)-1] < nums[i] {
				chain = chain[:len(chain)-1]
			}
			chain = append(chain, nums[i])
		} else {
			// Similarly, remove those smaller than our new element
			for len(chain) > 0 && chain[len(chain)-1] < nums[i] {
				chain = chain[:len(chain)-1]
			}
			chain = append(chain, nums[i])
			results = append(results, chain[0])
			// If our biggest element is at the beginning of the window, just remove it
			if chain[0] == nums[i-k+1] {
				chain = chain[1:]
			}
		}
	}

	return results
}

func maxSlidingWindow_Others_Solution(nums []int, k int) []int {
	var res []int

	var queue MonotonicQueue

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			queue.push(nums[i])
		} else {
			queue.push(nums[i])
			res = append(res, queue[0])
			queue.pop(nums[i-k+1])
		}
	}

	return res
}

type MonotonicQueue []int

func (m *MonotonicQueue) length() int {
	return len(*m)
}

func (m *MonotonicQueue) back() int {
	if m.length() == 0 {
		panic("empty queue")
	}

	return (*m)[len(*m)-1]
}

func (m *MonotonicQueue) pop(item int) {
	if m.length() > 0 && (*m)[0] == item {
		*m = (*m)[1:]
	}
}

func (m *MonotonicQueue) popBack() int {
	if len(*m) == 0 {
		panic("empty queue")
	}

	poppedItem := (*m)[len(*m)-1]

	*m = (*m)[:len(*m)-1]

	return poppedItem
}

func (m *MonotonicQueue) push(item int) {
	for m.length() > 0 && m.back() < item {
		m.popBack()
	}

	*m = append(*m, item)
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4, []int{7, 7, 7, 7, 7}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, maxSlidingWindow_Timeout(tt.nums, tt.k))
		assert.Equal(t, tt.want, maxSlidingWindow(tt.nums, tt.k))
		assert.Equal(t, tt.want, maxSlidingWindow_Others_Solution(tt.nums, tt.k))
	}
}
