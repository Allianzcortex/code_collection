package adventofcode

import (
	"strconv"
)

func day1Part1(puzzle []string) int {
	res := 0
	for _, line := range puzzle {
		var (
			start, end       = 0, len(line) - 1
			startVal, endVal = 0, 0
			err              error
		)

		for {
			// check start value and keep going right until finding a digit
			if startVal, err = strconv.Atoi(string(line[start])); err != nil {
				start += 1
			} else {
				break
			}
		}

		for {
			// check end value and keep going left until finding a digit
			if endVal, err = strconv.Atoi(string(line[end])); err != nil {
				end -= 1
			} else {
				break
			}
		}

		res += (startVal*10 + endVal)
	}

	return res
}

func day1Part2(puzzle []string) int {

	res := 0

	for _, line := range puzzle {
		var (
			startVal, endVal = 0, 0
		)

		for i := 0; i < len(line); i++ {
			if isFoundDigit, val := isValidDigitString(line, i, true); isFoundDigit {
				startVal = val
				break
			}
		}

		for j := len(line) - 1; j >= 0; j-- {
			if isFoundDigit, val := isValidDigitString(line, j, false); isFoundDigit {
				endVal = val
				break
			}
		}
		res += (startVal * 10) + endVal
	}
	return res
}

func isValidDigitString(input string, index int, isFromLeft bool) (bool, int) {
	validDigitString := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	if val, err := strconv.Atoi(string(input[index])); err == nil {
		return true, val
	}

	for digit, val := range validDigitString {

		length := len(digit)
		if isFromLeft {
			rightIndex := index + length - 1
			if rightIndex > len(input)-1 {
				continue
			}
			if input[index:rightIndex+1] == digit {
				return true, val
			}
		} else {
			leftIndex := index - length + 1
			if leftIndex <= 0 {
				continue
			}
			if input[leftIndex:index+1] == digit {
				return true, val
			}
		}
	}
	return false, 0
}
