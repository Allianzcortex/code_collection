package adventofcode

import (
	"sort"
)

// lol
// well : https://math.stackexchange.com/questions/113270/the-median-minimizes-the-sum-of-absolute-deviations-the-ell-1-norm
// In math we trust
func day7(positions []int) int {
	sort.Ints(positions)
	var median, res int
	middleIndex := len(positions) / 2

	if len(positions)%2 == 1 {
		median = positions[middleIndex]
	} else {
		median = (positions[middleIndex] + positions[middleIndex-1]) / 2
	}

	for _, p := range positions {
		res += abs(p, median)
	}

	return res
}
