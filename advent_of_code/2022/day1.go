package adventofcode

import "strconv"

func day1Part1(energy []string) int {
	curEnergy, maxEnergy := 0, 0

	for _, val := range energy {
		if (len(val)) == 0 {
			maxEnergy = max(curEnergy, maxEnergy)
			curEnergy = 0
			continue
		}
		item, _ := strconv.Atoi(val)
		curEnergy += item
	}

	return maxEnergy
}

func day1Part2(energy []string) int {
	curEnergy := 0
	maxEnergy1, maxEnergy2, maxEnergy3 := 0, 0, 0

	for _, val := range energy {
		if (len(val)) == 0 {
			// if curEnergy is the biggest
			if curEnergy >= maxEnergy1 {
				maxEnergy3 = maxEnergy2
				maxEnergy2 = maxEnergy1
				maxEnergy1 = curEnergy
				// if curEnergy is the 2nd biggest
			} else if curEnergy >= maxEnergy2 {
				maxEnergy3 = maxEnergy2
				maxEnergy2 = curEnergy
				// if curEnergy is the 3rd biggest
			} else if curEnergy >= maxEnergy3 {
				maxEnergy3 = curEnergy
			}

			curEnergy = 0
			continue
		}
		item, _ := strconv.Atoi(val)
		curEnergy += item
	}

	return maxEnergy1 + maxEnergy2 + maxEnergy3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
