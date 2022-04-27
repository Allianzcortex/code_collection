package adventofcode

import (
	"sort"
)

var match1 = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func day10Part1(inputs []string) int {

	points := 0

	for _, line := range inputs {
		stack := make([]rune, 0)
		for _, ch := range line {
			if counter, exist := match1[ch]; !exist {
				// if it doesn't exist,push it to stack directly
				stack = append(stack, ch)
			} else {
				// if we find it exists,check whether they're matched
				if counter == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					points += getPoint1(ch)
					break
				}
			}
		}
	}

	return points

}

func getPoint1(ch rune) int {
	switch ch {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return -1
	}
}

var match2 = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func day10Part2(input []string) int {
	points := make([]int, 0)

	for _, line := range input {
		stack := make([]rune, 0)
		for _, ch := range line {
			stack = append(stack, ch)
			if _, exist := match2[ch]; !exist {
				// for corrupted lines
				if stack[len(stack)-1-1] != match1[ch] {
					stack = nil
					break
				} else {
					// for incomplete lines
					stack = stack[:len(stack)-2]
				}

			}
		}
		// find from right->left,
		if stack != nil {
			res := 0
			for j := len(stack) - 1; j >= 0; j-- {
				res = res*5 + getPoint2(match2[stack[j]])
			}
			points = append(points, res)
		}
	}

	sort.Ints(points)
	return points[len(points)/2]
}

func getPoint2(ch rune) int {
	switch ch {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	default:
		return -1
	}
}
