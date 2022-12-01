package adventofcode

import (
	"fmt"
	"strings"
)

func day3Part1(puzzle []string) int {
	priorityPoints := 0

	for _, rucks := range puzzle {
		length := len(rucks)
		res := getPoint(findCommonType(rucks[0:length/2], rucks[length/2:length]))
		priorityPoints += res

	}

	return priorityPoints
}

func day3Part2(puzzle []string) int {
	priorityPoints := 0

	for i := 0; i < len(puzzle); i += 3 {
		fmt.Println(i)
		commonBadge := findCommonBadge(puzzle[i], puzzle[i+1], puzzle[i+2])
		res := getPoint(commonBadge)
		priorityPoints += res
	}

	return priorityPoints
}

func getPoint(char rune) int {
	// fmt.Printf("%v", string(char))
	switch {
	case 'a' <= char && char <= 'z':
		return 1 + int(char-'a')
	case 'A' <= char && char <= 'Z':
		return 27 + int(char-'A')
	}

	return -1
}

func findCommonType(part1, part2 string) rune {
	m := make(map[rune]bool)
	for _, c := range part1 {
		m[c] = true
	}

	for _, c := range part2 {
		if _, exist := m[c]; exist {
			return c
		}
	}

	return -1
}

func findCommonBadge(part1, part2, part3 string) rune {
	// this we can optimize it by iterating from the one with minimum length
	// but doesn't matter tho

	m := make(map[rune]bool)
	for _, c := range part1 {
		m[c] = true
	}

	// iterate through map
	// delete all items that doesnot exist in part2
	for key := range m {
		if !strings.ContainsRune(part2, key) {
			delete(m, key)
		}
	}

	for _, c := range part3 {
		if _, exist := m[c]; exist {
			return c
		}
	}

	return -1
}
