package leetcode

func (ad *AddTwoNumbers) Solver(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	l3ptr := l3

	var carry int

	for {
		var n1 int
		var n2 int

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		n3 := n1 + n2 + carry

		if n3 >= 10 {
			carry = n3 / 10
			l3ptr.Val = n3 - 10
		} else {
			carry = 0
			l3ptr.Val = n3
		}

		if l1 != nil || l2 != nil || carry != 0 {
			l3ptr.Next = &ListNode{}
			l3ptr = l3ptr.Next
		} else {
			break
		}
	}

	return l3
}
