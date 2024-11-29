package adventofcode

import (
	"math"
	"slices"
	"strings"
)

func day4Part1(input []string) int {

	res := 0
	// we can use regex here, but keep splitting^
	for _, line := range input {
		cardContent := strings.Split(line, ":")[1]
		numbers := strings.Split(cardContent, "|")
		winningNumbers, currentNumbers := numbers[0], numbers[1]

		cnt := 0
		cleanWinningNumbers := strings.Fields(winningNumbers)
		for _, num := range strings.Fields(currentNumbers) {
			if slices.Contains(cleanWinningNumbers, num) {
				cnt += 1
			}
		}
		if cnt != 0 {
			res += int(math.Exp2(float64(cnt - 1)))
		}
	}

	return res
}

func day4Part2(input []string) int {

	res := 0

	// pre-fill match relationship
	match := make(map[int]int, len(input))
	for i := 1; i <= len(input); i++ {
		match[i] = 1
	}

	for index, line := range input {
		cardNumber := index + 1

		// can extract the following content into a share func
		cardContent := strings.Split(line, ":")[1]
		numbers := strings.Split(cardContent, "|")
		winningNumbers, currentNumbers := numbers[0], numbers[1]

		cnt := 0
		cleanWinningNumbers := strings.Fields(winningNumbers)
		for _, num := range strings.Fields(currentNumbers) {
			if slices.Contains(cleanWinningNumbers, num) {
				cnt += 1
			}
		}

		// next generate number of new cards
		originalCnt := match[cardNumber]
		for i := 1; i <= cnt; i++ {
			match[cardNumber+i] += 1 * originalCnt
		}
	}

	for _, scratchCardNum := range match {
		res += scratchCardNum
	}
	return res
}
