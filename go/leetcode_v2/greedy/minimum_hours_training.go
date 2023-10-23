package greedy

/*
You are entering a competition, and are given two positive integers initialEnergy
and initialExperience denoting your initial energy and initial experience
respectively.

You are also given two 0-indexed integer arrays energy and experience, both of
length n.

You will face n opponents in order. The energy and experience of the ith opponent
is denoted by energy[i] and experience[i] respectively. When you face an opponent,
you need to have both strictly greater experience and energy to defeat them
and move to the next opponent if available.

Defeating the ith opponent increases your experience by experience[i], but
decreases your energy by energy[i].

Before starting the competition, you can train for some number of hours. After
each hour of training, you can either choose to increase your initial experience
by one, or increase your initial energy by one.

Return the minimum number of training hours required to defeat all n opponents.
*/

func minNumberOfHours(initialEnergy int, initialExperience int, energies []int, experiences []int) int {
	/*
		Input: initialEnergy = 5, initialExperience = 3, energy = [1,4,3,2], experience = [2,6,3,1]

		Solution:
		energyRequired = sum(energy) + 1 			-> 10 + 1
		if initialEnergy < energyRequired {   -> 5 < 11
			trainingHour += energyRequired - initialEnergy   -> 11 - 5 = 6
		}

		for _, exp := range experience {
			if initialExperience <= exp {			-> 11 < 3
				expRequired := exp - initialExperience + 1		-> 6 - 5 + 1 = 2
				trainingExp += expRequired										-> 2
				initialExp += expRequired											-> initialExp = 7
			}

			initialExperience += exp					-> initialExperience = 11
		}

		return trainingExp + trainingHour
	*/

	var trainingHour int

	// Prefix Sum our energies
	for i := 1; i < len(energies); i++ {
		energies[i] += energies[i-1]
	}

	if initialEnergy <= energies[len(energies)-1] {
		trainingHour += energies[len(energies)-1] - initialEnergy + 1
	}

	for _, exp := range experiences {
		if initialExperience <= exp {
			expRequired := exp - initialExperience + 1
			trainingHour += expRequired
			initialExperience += expRequired
		}

		initialExperience += exp
	}

	return trainingHour
}

func minNumberOfHoursV2(initialEnergy int, initialExperience int, energies []int, experiences []int) int {
	var trainingHour int

	for i := 0; i < len(energies); i++ {
		exp := experiences[i]
		energy := energies[i]

		if initialExperience <= exp {
			expRequired := exp - initialExperience + 1
			trainingHour += expRequired
			initialExperience += expRequired
		}

		if initialEnergy <= energy {
			energyRequired := energy - initialEnergy + 1
			trainingHour += energyRequired
			initialEnergy += energyRequired
		}

		initialEnergy -= energy
		initialExperience += exp
	}

	return trainingHour
}
