package adventofcode

import "testing"

var input6 = []string{
	"Time:        48     93     85     95",
	"Distance:   296   1928   1236   1391",
}

func TestDay6Part1(t *testing.T) {
	expectedRes := 2756160
	if res := day6Part1(input6); res != expectedRes {
		t.Errorf("expected res is %v,what we get is %v", expectedRes, res)
	}
}

func TestDay6Part2(t *testing.T) {
	expectedRes := 34788142
	if res := day6Part2(input6); res != expectedRes {
		t.Errorf("expected res is %v,what we get is %v", expectedRes, res)
	}
}
