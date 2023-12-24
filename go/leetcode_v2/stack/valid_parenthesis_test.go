package stack

import "testing"

func isValid(s string) bool {
	// Consists of its closing
	// and opening bracket pair
	closings := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if opening, found := closings[s[i]]; found { // If current char is a closing bracket
			// Checks if the previous opening in
			// stack is indeed its closing pair
			//
			// If there was no previous openings then we return false
			if len(stack) == 0 {
				return false
			}
			//
			// If not its pair, we return false immediately
			if prevOpening := stack[len(stack)-1]; opening != prevOpening {
				return false
			}

			// If it is, we remove from the stack
			stack = stack[:len(stack)-1]
		} else { // Else current char is opening bracket
			// Add to the stack.
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

func Test_isValid(t *testing.T) {
	if result := isValid("()"); result != true {
		t.Fatalf("expected true but got %t", result)
	}
	if result := isValid("()[]{}"); result != true {
		t.Fatalf("expected true but got %t", result)
	}
	if result := isValid("(]"); result != false {
		t.Fatalf("expected false but got %t", result)
	}
	if result := isValid("((((((()))))))"); result != true {
		t.Fatalf("expected true but got %t", result)
	}
	if result := isValid("(((((((]))))))"); result != false {
		t.Fatalf("expected true but got %t", result)
	}
}
