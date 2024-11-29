package adventofcode

import (
	"strings"
)

func day15Part1(puzzle []string) int {

	res := 0
	for _, input := range strings.Split(strings.Join(puzzle, ""), ",") {
		prevRes := 0
		for _, ch := range input {
			prevRes += int(ch)
			prevRes *= 17
			prevRes = prevRes % 256
		}
		res += prevRes

	}

	return res
}
