package stack

import (
	"testing"
)

type MinStack struct {
	// The stack as an array
	stack []int
	// The length of the stack
	len int

	// Holds the links of the current minimum index at the
	// stack with its previous at the period of operation.
	// For example:
	// Given these insertions: 4 -> 1 -> 2 -> -1 -> -2
	// we'll have the stack as [4, 1, 2, -1, -2].
	// Hence, the minLink is:
	// { 1: 0, 3: 1, 4: 3 }
	minLink map[int]int
	// Holds the current minimum index pointing at the stack
	currMinIndex int
}

func Constructor() MinStack {
	return MinStack{
		stack:   make([]int, 0),
		minLink: make(map[int]int),
	}
}

// Push appends the value to the stack and evaluate the minimum value.
func (this *MinStack) Push(val int) {
	// When the stack is empty, we want our
	// current minimum index to be that new
	// value. On the other hand, if the new
	// value is less than the stack's current
	// min index, we want to create a new link
	// to it.
	if this.len == 0 {
		this.currMinIndex = 0
	} else if val < this.stack[this.currMinIndex] {
		this.minLink[this.len] = this.currMinIndex
		this.currMinIndex = this.len
	}

	// Append new value to the stack
	this.stack = append(this.stack, val)
	this.len++
}

// Pop removes the value from the stack and evaluate the minimum value.
func (this *MinStack) Pop() {
	// If the value being popped is the currMinIndex's value
	// then we want to make currMinIndex now pointing
	// to its previous index in minLink and then remove
	// from minLink.
	if prevMinIndex, ok := this.minLink[this.len-1]; ok {
		this.currMinIndex = prevMinIndex
		delete(this.minLink, this.len-1)
	}

	// Removes from stack
	this.stack = this.stack[:this.len-1]
	this.len--
}

func (this *MinStack) Top() int {
	return this.stack[this.len-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.currMinIndex]
}

func TestMinStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(4)
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(-1)
	minStack.Push(-2)

	if top := minStack.Top(); top != -2 {
		t.Fatalf("expected Top() to be -2 but got %d", top)
	}

	if min := minStack.GetMin(); min != -2 {
		t.Fatalf("expected GetMin() to be -2 but got %d", min)
	}

	minStack.Pop()

	if min := minStack.GetMin(); min != -1 {
		t.Fatalf("expected GetMin() to be -1 but got %d", min)
	}

	minStack.Pop()

	if min := minStack.GetMin(); min != 1 {
		t.Fatalf("expected GetMin() to be 1 but got %d", min)
	}

	minStack.Push(5)

	if min := minStack.GetMin(); min != 1 {
		t.Fatalf("expected GetMin() to be 1 but got %d", min)
	}

	if top := minStack.Top(); top != 5 {
		t.Fatalf("expected Top() to be 5 but got %d", top)
	}

	minStack.Push(0)

	if min := minStack.GetMin(); min != 0 {
		t.Fatalf("expected GetMin() to be 0 but got %d", min)
	}

	if top := minStack.Top(); top != 0 {
		t.Fatalf("expected Top() to be 0 but got %d", top)
	}
}

type MinStack2 struct {
	// The original stack
	stack []int
	// Holds the minimum indexes in the stack
	mininmums []int
}

func Constructor2() MinStack2 {
	return MinStack2{stack: make([]int, 0), mininmums: make([]int, 0)}
}

func (this *MinStack2) Push(val int) {
	if len(this.stack) == 0 {
		this.mininmums = []int{0}
	} else {
		currentMinimum := this.stack[this.mininmums[len(this.mininmums)-1]]
		if val < currentMinimum {
			this.mininmums = append(this.mininmums, len(this.stack))
		}
	}

	this.stack = append(this.stack, val)
}

func (this *MinStack2) Pop() {
	poppedIndex := len(this.stack) - 1

	this.stack = this.stack[:poppedIndex]

	if this.mininmums[len(this.mininmums)-1] == poppedIndex {
		this.mininmums = this.mininmums[:len(this.mininmums)-1]
	}
}

func (this *MinStack2) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack2) GetMin() int {
	return this.stack[this.mininmums[len(this.mininmums)-1]]
}

func TestMinStack2(t *testing.T) {
	minStack := Constructor2()

	minStack.Push(4)
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(-1)
	minStack.Push(-2)

	if top := minStack.Top(); top != -2 {
		t.Fatalf("expected Top() to be -2 but got %d", top)
	}

	if min := minStack.GetMin(); min != -2 {
		t.Fatalf("expected GetMin() to be -2 but got %d", min)
	}

	minStack.Pop()

	if min := minStack.GetMin(); min != -1 {
		t.Fatalf("expected GetMin() to be -1 but got %d", min)
	}

	minStack.Pop()

	if min := minStack.GetMin(); min != 1 {
		t.Fatalf("expected GetMin() to be 1 but got %d", min)
	}

	minStack.Push(5)

	if min := minStack.GetMin(); min != 1 {
		t.Fatalf("expected GetMin() to be 1 but got %d", min)
	}

	if top := minStack.Top(); top != 5 {
		t.Fatalf("expected Top() to be 5 but got %d", top)
	}

	minStack.Push(0)

	if min := minStack.GetMin(); min != 0 {
		t.Fatalf("expected GetMin() to be 0 but got %d", min)
	}

	if top := minStack.Top(); top != 0 {
		t.Fatalf("expected Top() to be 0 but got %d", top)
	}
}

type MinStackLL struct {
	head *node
}

type node struct {
	val, minVal int
	prev        *node
}

func ConstructorLL() MinStackLL {
	return MinStackLL{}
}

func (this *MinStackLL) Push(val int) {
	newNode := &node{val: val, prev: this.head}

	if this.head == nil {
		newNode.minVal = val
	} else if val < this.head.minVal {
		newNode.minVal = val
	} else {
		newNode.minVal = this.head.minVal
	}

	this.head = newNode
}

func (this *MinStackLL) Pop() {
	this.head = this.head.prev
}

func (this *MinStackLL) Top() int {
	return this.head.val
}

func (this *MinStackLL) GetMin() int {
	return this.head.minVal
}

func TestMinStackLL(t *testing.T) {
	minStack := ConstructorLL()

	minStack.Push(4)
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(-1)
	minStack.Push(-2)

	if top := minStack.Top(); top != -2 {
		t.Fatalf("expected Top() to be -2 but got %d", top)
	}

	if min := minStack.GetMin(); min != -2 {
		t.Fatalf("expected GetMin() to be -2 but got %d", min)
	}

	minStack.Pop()

	if min := minStack.GetMin(); min != -1 {
		t.Fatalf("expected GetMin() to be -1 but got %d", min)
	}

	minStack.Pop()

	if min := minStack.GetMin(); min != 1 {
		t.Fatalf("expected GetMin() to be 1 but got %d", min)
	}

	minStack.Push(5)

	if min := minStack.GetMin(); min != 1 {
		t.Fatalf("expected GetMin() to be 1 but got %d", min)
	}

	if top := minStack.Top(); top != 5 {
		t.Fatalf("expected Top() to be 5 but got %d", top)
	}

	minStack.Push(0)

	if min := minStack.GetMin(); min != 0 {
		t.Fatalf("expected GetMin() to be 0 but got %d", min)
	}

	if top := minStack.Top(); top != 0 {
		t.Fatalf("expected Top() to be 0 but got %d", top)
	}
}
