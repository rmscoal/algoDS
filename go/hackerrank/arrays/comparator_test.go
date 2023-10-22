package arrays

import (
	"fmt"
	"testing"
)

func TestCompareAlphabet(t *testing.T) {
	result := compareAlphabet("Smith", "Jones")
	if result {
		t.FailNow()
	}

	result = compareAlphabet("amy", "david")
	if !result {
		t.Log("Expected is true but got false")
		t.FailNow()
	}

	result = compareAlphabet("amy", "amya")
	if !result {
		t.Log("Expected is true but got false")
		t.FailNow()
	}

	result = compareAlphabet("ade", "c")
	if !result {
		t.Log("Expected is true but got false")
		t.FailNow()
	}

	result = compareAlphabet("bda", "cab")
	if !result {
		t.Log("Expected is true but got false")
		t.FailNow()
	}
}

func TestComparator(t *testing.T) {
	tests := map[string]struct {
		names        []string
		scores       []int
		answerNames  []string
		answerScores []int
	}{
		"test_1": {
			names:        []string{"Smith", "Jones", "Jones"},
			scores:       []int{20, 15, 20},
			answerNames:  []string{"Jones", "Smith", "Jones"},
			answerScores: []int{20, 20, 15},
		},
		"test_2": {
			names:        []string{"amy", "david", "heraldo", "aakansha", "aleksa"},
			scores:       []int{100, 100, 50, 75, 150},
			answerNames:  []string{"aleksa", "amy", "david", "aakansha", "heraldo"},
			answerScores: []int{150, 100, 100, 75, 50},
		},
		"test_3": {
			names:        []string{"ab", "bcc", "ade", "cab", "dee", "bda", "c", "db", "a", "cbb"},
			scores:       []int{6, 0, 5, 2, 0, 2, 5, 2, 1, 1},
			answerNames:  []string{"ab", "ade", "c", "bda", "cab", "db", "a", "cbb", "bcc", "dee"},
			answerScores: []int{6, 5, 5, 2, 2, 2, 1, 1, 0, 0},
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Comparator %s\n", name), func(t *testing.T) {
			names := make([]string, len(test.names))
			scores := make([]int, len(test.scores))
			copy(names, test.names)
			copy(scores, test.scores)
			comparator(names, scores)

			for index, v := range names {
				if v != test.answerNames[index] {
					t.Fatalf("Expected name at index %d is %s but got %s\n", index, test.answerNames[index], v)
				}
			}
			for index, v := range scores {
				if v != test.answerScores[index] {
					t.Fatalf("Expected score at index %d is %d but got %d\n", index, test.answerScores[index], v)
				}
			}
		})
	}
}
