package adventofcode

import (
	"math"
	"sort"
)

// well : https://math.stackexchange.com/questions/113270/the-median-minimizes-the-sum-of-absolute-deviations-the-ell-1-norm
// In math we trust
func day7Part1(positions []int) int {
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

// use brute force
func day7Part2(positions []int) int {
	sort.Ints(positions)
	startPos, endPos := positions[0], positions[len(positions)-1]
	res := math.MaxInt64

	for ; startPos <= endPos; startPos++ {
		fuel := calculateFuel(startPos, positions)
		res = min(res, fuel)
	}

	return res
}

// n*(n+1)/2
func calculateFuel(pos int, positions []int) int {

	distance := 0
	for _, num := range positions {
		step := abs(num, pos)
		distance += (step) * (step + 1) / 2
	}
	return distance
}
