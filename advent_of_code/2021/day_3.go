package adventofcode

import (
	"math"
)

func day3(input []string) int {
	length := len(input[0])
	metaArray := make([]int, length)

	for _, report := range input {
		for index := 0; index < length; index += 1 {
			if report[index] == '1' {
				metaArray[index] += 1
			} else {
				metaArray[index] -= 1
			}
		}
	}

	gamma, epsilon := 0, 0

	for i := 0; i < length; i++ {
		ch := metaArray[length-1-i]
		if ch > 0 {
			gamma += int(math.Pow(2, float64(i)))
		} else {
			epsilon += int(math.Pow(2, float64(i)))
		}
	}

	return gamma * epsilon
}
