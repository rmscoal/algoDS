package string

import (
	"fmt"
	"testing"
)

func TestCommonChild(t *testing.T) {
	tests := map[string]struct {
		s1, s2 string
		answer int32
	}{
		"test_1": {
			s1:     "ABCD",
			s2:     "ABDC",
			answer: 3,
		},
		"test_2": {
			s1:     "HARRY",
			s2:     "SALLY",
			answer: 2,
		},
		"test_3": {
			s1:     "AA",
			s2:     "BB",
			answer: 0,
		},
		"test_4": {
			s1:     "SHINCHAN",
			s2:     "NOHARAAA",
			answer: 3,
		},
		"test_5": {
			s1:     "ABCDEF",
			s2:     "FBDAMN",
			answer: 2,
		},
		"test_6": {
			s1:     "AGGTAB",
			s2:     "GXTXAY",
			answer: 3,
		},
		"test_7": {
			s1:     "WEWOUCUIDGCGTRMEZEPXZFEJWISRSBBSYXAYDFEJJDLEBVHHKS",
			s2:     "FDAGCXGKCTKWNECHMRXZWMLRYUCOCZHJRRJBOAJOQJZZVUYXIC",
			answer: 15,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Common Child %s", name), func(t *testing.T) {
			if result := commonChild(test.s1, test.s2); result != test.answer {
				t.Fatalf("Testing %s expected %d but got result %d\n", name, test.answer, result)
			}
		})
	}
}
