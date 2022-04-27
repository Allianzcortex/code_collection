package adventofcode

import "testing"

var octos = [][]int{
	{4, 6, 5, 8, 1, 3, 7, 6, 3, 7},
	{3, 2, 7, 7, 8, 7, 4, 3, 5, 5},
	{4, 5, 2, 5, 6, 1, 1, 1, 8, 3},
	{3, 1, 2, 8, 1, 2, 5, 8, 8, 8},
	{8, 7, 3, 4, 8, 3, 2, 8, 3, 8},
	{4, 1, 7, 5, 4, 6, 3, 2, 5, 7},
	{8, 3, 2, 1, 4, 2, 3, 5, 5, 2},
	{4, 8, 3, 2, 1, 4, 5, 2, 5, 3},
	{8, 2, 8, 6, 8, 3, 4, 8, 5, 1},
	{4, 8, 8, 5, 3, 2, 3, 1, 3, 8},
}

func TestDay11(t *testing.T) {
	expectedRound := 1
	if flashZeroRound := day11(octos); flashZeroRound != expectedRound {
		t.Errorf("expected flash zero round is %v,what we get is %v", flashZeroRound, expectedRound)
	}
}
