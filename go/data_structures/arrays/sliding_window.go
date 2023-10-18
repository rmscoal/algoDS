package arrays

/*
Question 1: Given an array of positive integers nums and an integer k, find the
length of the longest subarray whose sum is less than or equal to k. This is the
problem we have been talking about above. We will now formally solve it.
*/

func SlidingWindowQuestionOneSolutionOne(s []int, k int) int {
	left, max, cursor := 0, 0, 0
	n := len(s)

	for right := 0; right < n; right++ {
		cursor += s[right]

		for cursor > k {
			cursor -= s[left]
			left++
		}

		if max <= right-left+1 {
			max = right - left + 1
		}
	}

	return max
}

/*
Example 2: You are given a binary string s (a string containing only "0" and
"1"). You may choose up to one "0" and flip it to a "1". What is the length of
the longest substring achievable that contains only "1"?
*/

func SlidingWindowQuestionTwoSolutionOne(s string) int {
	left, max, cursor := 0, 0, 0

	for right := 0; right < len(s); right++ {
		if s[right] == '0' {
			cursor++
		}

		for cursor >= 2 {
			if s[left] == '0' {
				cursor--
			}
			left++
		}

		if max <= right-left+1 {
			max = right - left + 1
		}
	}

	return max
}

/*
Given an array of positive integers nums and an integer k, return the number of
subarrays where the product of all the elements in the subarray is strictly less
than k.
*/

func SlidingWindowQuestionThreeSolutionOne(arr []int, k int) int {
	if k == 1 {
		return 0
	}

	left, cursor, total := 0, 1, 0
	n := len(arr)

	for right := 0; right < n; right++ {
		cursor *= arr[right]

		for cursor >= k {
			cursor /= arr[left]
			left++
		}

		total += right - left + 1
	}

	return total
}

/*
Given an integer array nums and an integer k, find the sum of the subarray with
the largest sum whose length is k.
*/

func SlidingWindowQuestionFourSolutionOne(arr []int, k int) int {
	max, cursor := 0, 0

	for right := 0; right < k; right++ {
		cursor += arr[right]
	}

	max = cursor

	for right := k; right < len(arr); right++ {
		cursor += arr[right] - arr[right-k]

		if max <= cursor {
			max = cursor
		}
	}

	return max
}
