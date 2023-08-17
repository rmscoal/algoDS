package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateSingleLinkedListFromSlice(arr []int) *ListNode {
	head := &ListNode{}
	headptr := head

	for i, v := range arr {
		if i != 0 {
			headptr.Next = &ListNode{}
			headptr = headptr.Next
		}

		headptr.Val = v
	}

	return head
}
