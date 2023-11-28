package dp

import "testing"

func TestFindMax(t *testing.T) {
	if result := findMax(1, 2); result != 2 {
		t.Fatal("Expcted 2 but got", result)
	}

	if result := findMax(2, 1); result != 2 {
		t.Fatal("Expcted 2 but got", result)
	}

	if result := findMax(-2, 1); result != 1 {
		t.Fatal("Expcted 1 but got", result)
	}
}

func TestFindMin(t *testing.T) {
	if result := findMin(1, 2); result != 1 {
		t.Fatal("Expcted 1 but got", result)
	}

	if result := findMin(2, 1); result != 1 {
		t.Fatal("Expcted 1 but got", result)
	}

	if result := findMin(-2, 1); result != -2 {
		t.Fatal("Expcted -2 but got", result)
	}
}
