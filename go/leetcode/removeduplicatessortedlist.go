package leetcode

func (RSL *RemoveDuplicateSortedList) Solver(head *ListNode) *ListNode {
	result := ListNode{}
	tail := &result

	// Checks if there are no linked list or only one chain of list.
	if head == nil || head.Next == nil {
		return head
	}

	// Two pointers.
	tail.Next = head
	tail = tail.Next
	head = head.Next

	for head != nil {

		if head.Val > tail.Val {
			tail.Next = head
			tail = tail.Next
		}

		head = head.Next

	}

	tail.Next = nil

	return result.Next
}

func (RSL *RemoveDuplicateSortedList) FastSolver(head *ListNode) *ListNode {
	temp := head

	if head == nil {
		return head
	}

	for {
		if temp.Next != nil && temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}

		if temp == nil {
			break
		}
	}

	return head
}
