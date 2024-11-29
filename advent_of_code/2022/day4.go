package adventofcode

import (
	"strconv"
	"strings"
)

func day4Part1(puzzle []string) int {
	cnt := 0
	for _, val := range puzzle {
		cnt += getFullRangeCount(processData(val))
	}

	return cnt

}

func day4Part2(puzzle []string) int {
	cnt := 0
	for _, val := range puzzle {
		cnt += getPartialRangeCount(processData(val))
	}

	return cnt
}

func processData(item string) (int, int, int, int) {
	var input = strings.Split(item, ",")
	section1, section2 := input[0], input[1]
	// get range of first section
	section1Input := strings.Split(section1, "-")
	left1, _ := strconv.Atoi(section1Input[0])
	right1, _ := strconv.Atoi(section1Input[1])

	// get range of second section
	section2Input := strings.Split(section2, "-")
	left2, _ := strconv.Atoi(section2Input[0])
	right2, _ := strconv.Atoi(section2Input[1])

	return left1, right1, left2, right2
}

func getFullRangeCount(left1, right1, left2, right2 int) int {
	if (left1 >= left2 && right1 <= right2) || (left1 <= left2 && right1 >= right2) {
		return 1
	}
	return 0
}

func getPartialRangeCount(left1, right1, left2, right2 int) int {
	// considering the following two conditions:
	//            [.......]  left1 should be in range of [left2,right2]
	//        [.......]
	//
	//            [.......]
	//                [......]  right1 should be in range of [left2,right2]
	if (left1 <= right2 && right2 <= right1) || (right1 >= left2 && right1 <= right2) {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
