package arrays

import "testing"

func TestSlidingWindowQuestionOneSolutionOne(t *testing.T) {
	num1 := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	k1 := 8
	ans1 := 4

	t.Run("Question 1 Solution 1 Case 1", func(t *testing.T) {
		result := SlidingWindowQuestionOneSolutionOne(num1, k1)
		if result != ans1 {
			t.Fatalf("Expected is %d but result was %d", ans1, result)
		}
	})
}

func TestSlidingWindowQuestionTwoSolutionOne(t *testing.T) {
	str1 := "1101100111"
	ans1 := 5

	t.Run("Question 2 Solution 1 Case 1", func(t *testing.T) {
		result := SlidingWindowQuestionTwoSolutionOne(str1)
		if result != ans1 {
			t.Fatalf("Expected is %d but result was %d", ans1, result)
		}
	})
}

func TestSlidingWindowQuestionThreeSolutionOne(t *testing.T) {
	arr1 := []int{10, 5, 2, 6}
	k := 100
	ans1 := 8

	t.Run("Question 3 Solution 1 Case 1", func(t *testing.T) {
		result := SlidingWindowQuestionThreeSolutionOne(arr1, k)
		if result != ans1 {
			t.Fatalf("Expected is %d but result was %d", ans1, result)
		}
	})
}

func TestSlidingWindowQuestionFourSolutionOne(t *testing.T) {
	arr1 := []int{10, 5, 2, 6}
	k := 2
	ans1 := 15

	t.Run("Question 4 Solution 1 Case 1", func(t *testing.T) {
		result := SlidingWindowQuestionFourSolutionOne(arr1, k)
		if result != ans1 {
			t.Fatalf("Expected is %d but result was %d", ans1, result)
		}
	})
}
