package main

import (
	"testing"
)

func TestProblem26(t *testing.T) {
	input1 := []int{5, 7, 6, 9, 11, 10, 8}
	input2 := []int{7, 4, 6, 5}
	if RuntProblem26(input1) != true && RuntProblem26(input2) != false {
		t.Error("wrong")
	}
}

func TestProblem11(t *testing.T) {
	if !(RunProblem11(3, 2) == 9 && RunProblem11(0, 0) == 0) {
		t.Error("Wrong when test problem11")
	}
}
