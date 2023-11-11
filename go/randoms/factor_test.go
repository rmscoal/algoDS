package randoms

import (
	"fmt"
	"testing"
)

func TestPthFactor(t *testing.T) {
	tests := map[string]struct {
		num, p, answer int
	}{
		"test_1": {
			num:    20,
			p:      3,
			answer: 4,
		},
		"test_2": {
			num:    15,
			p:      2,
			answer: 3,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Pth Factor %s", name), func(t *testing.T) {
			if result := pthFactor(test.num, test.p); result != test.answer {
				t.Fatalf("Expectation %s is %d but got %d\n", name, test.answer, result)
			}
		})
	}
}
