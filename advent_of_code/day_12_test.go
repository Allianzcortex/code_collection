package adventofcode

import "testing"

var paths = []string{
	"KF-sr",
	"OO-vy",
	"start-FP",
	"FP-end",
	"vy-mi",
	"vy-KF",
	"vy-na",
	"start-sr",
	"FP-lh",
	"sr-FP",
	"na-FP",
	"end-KF",
	"na-mi",
	"lh-KF",
	"end-lh",
	"na-start",
	"wp-KF",
	"mi-KF",
	"vy-sr",
	"vy-lh",
	"sr-mi",
}

func TestDay12(t *testing.T) {
	expectedRes := 4885
	if res := day12(paths); res != expectedRes {
		t.Errorf("expected res is %v,what we get is %v", expectedRes, res)
	}
}
