package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	prev := head
	cursor := head.Next
	head.Next = nil
	for cursor != nil {
		tmp := cursor.Next
		cursor.Next = prev
		prev = cursor
		cursor = tmp
	}

	return prev
}

func printList(head *ListNode) {
	for head != nil {
		println("VALUE", head.Val)
		head = head.Next
	}
}

func TestReverseList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	head = reverseList(head)
	printList(head)

	assert.Equal(t, 5, head.Val)
	assert.Equal(t, 4, head.Next.Val)
	assert.Equal(t, 3, head.Next.Next.Val)
	assert.Equal(t, 2, head.Next.Next.Next.Val)
	assert.Equal(t, 1, head.Next.Next.Next.Next.Val)
}
