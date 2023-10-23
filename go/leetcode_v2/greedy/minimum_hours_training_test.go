package greedy

import (
	"fmt"
	"testing"
)

func TestMinimumTrainingHour(t *testing.T) {
	tests := map[string]struct {
		initialEnergy, initialExperience, answer int
		energies, experiences                    []int
	}{
		"test_1": {
			initialEnergy:     5,
			initialExperience: 3,
			energies:          []int{1, 4, 3, 2},
			experiences:       []int{2, 6, 3, 1},
			answer:            8,
		},
		"test_2": {
			initialEnergy:     2,
			initialExperience: 4,
			energies:          []int{1},
			experiences:       []int{3},
			answer:            0,
		},
		"test_3": {
			initialEnergy:     5,
			initialExperience: 3,
			energies:          []int{1, 5, 3, 2},
			experiences:       []int{2, 7, 3, 1},
			answer:            10,
		},
	}

	for name, test := range tests {
		t.Run(fmt.Sprintf("Test Minimum Training Hour %s", name), func(t *testing.T) {
			energies := make([]int, len(test.energies))
			exps := make([]int, len(test.experiences))
			copy(energies, test.energies)
			copy(exps, test.experiences)

			result := minNumberOfHours(test.initialEnergy, test.initialExperience, energies, exps)
			if result != test.answer {
				t.Fatalf("Expected %d but got %d\n", test.answer, result)
			}

			result = minNumberOfHoursV2(test.initialEnergy, test.initialExperience, test.energies, test.experiences)
			if result != test.answer {
				t.Fatalf("Expected %d but got %d\n", test.answer, result)
			}
		})
	}
}
