package adventofcode

import (
	"strconv"
	"strings"
)

func day6Part1(input []string) int {
	res := 1

	time := strings.Fields(strings.Split(input[0], ":")[1])
	distance := strings.Fields(strings.Split(input[1], ":")[1])
	for i := 0; i < len(time); i++ {
		curTime, _ := strconv.Atoi(time[i])
		curRecord, _ := strconv.Atoi(distance[i])

		waitTime := 0
		for j := 1; j <= curTime-1; j++ {
			if j*(curTime-j) > curRecord {
				waitTime = j
				break
			}
		}

		if waitTime != 0 {
			res *= (curTime - waitTime - waitTime + 1)
		}
	}

	return res
}

func day6Part2(input []string) int {

	time := strings.Fields(strings.Split(input[0], ":")[1])
	distance := strings.Fields(strings.Split(input[1], ":")[1])

	totalTime, _ := strconv.Atoi(strings.Join(time, ""))
	totalDistance, _ := strconv.Atoi(strings.Join(distance, ""))
	for i := 1; i <= totalTime-1; i++ {
		if i*(totalTime-i) > totalDistance {
			return (totalTime - 2*i + 1)
		}
	}

	return 0
}
