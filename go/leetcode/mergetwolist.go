package leetcode

// list1 and list2 is the head node.
func (M *MergeTwoList) Solver(list1 *ListNode, list2 *ListNode) *ListNode {
	result := ListNode{}
	tail := &result

	for list1 != nil || list2 != nil {
		/*
			Checks whether either list1 or list2 is empty. If it is
			then sets the remaining tail to the remain available
			linked list.
		*/
		if list1 == nil {
			tail.Next = list2
			break
		}

		if list2 == nil {
			tail.Next = list1
			break
		}

		/*
			Compares the linked list's values to append next to the
			list. Once compared, set the lists pointer to its respective
			next pointer.
		*/
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		// Start fresh of new tail
		tail = tail.Next
	}

	// Returns the head node
	return result.Next
}
