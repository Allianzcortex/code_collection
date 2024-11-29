package adventofcode

import (
	"strings"
)

func day2Part1(puzzle []string) int {
	points := 0

	for _, round := range puzzle {
		input := strings.Split(round, " ")
		gameResult := runStrategry(input[0], input[1])
		points += calculatePoints(gameResult, input[1])
	}

	return points
}

func calculatePoints(result, action string) int {
	score := 0

	// step 1 based on game result
	switch result {
	case "win":
		score += 6
	case "draw":
		score += 3
	case "lose":
		score += 0
	}

	// step 2 based on choice
	switch action {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}
	return score
}

// return result that should be one of "win,draw,lose"
func runStrategry(action1, action2 string) string {
	var realAction1, realAction2 string

	switch action1 {
	case "A":
		realAction1 = "Rock"
	case "B":
		realAction1 = "Paper"
	case "C":
		realAction1 = "Scissors"
	}

	switch action2 {
	case "X":
		realAction2 = "Rock"
	case "Y":
		realAction2 = "Paper"
	case "Z":
		realAction2 = "Scissors"
	}

	if realAction1 == realAction2 {
		return "draw"
	}
	if (realAction1 == "Rock" && realAction2 == "Paper") || (realAction1 == "Paper" && realAction2 == "Scissors") ||
		(realAction1 == "Scissors" && realAction2 == "Rock") {
		return "win"
	}

	return "lose"
}

func day2Part2(puzzle []string) int {
	points := 0

	for _, round := range puzzle {
		input := strings.Split(round, " ")
		points += runStrategry2(input[0], input[1])
	}
	return points
}

func runStrategry2(action, result string) int {
	score := 0
	var expectedShape string
	// calculate points based on result
	switch result {
	case "X":
		score += 0
	case "Y":
		score += 3
	case "Z":
		score += 6
	}

	// calculate points based on shape
	switch action {
	case "A":
		switch result {
		case "X":
			expectedShape = "S"
		case "Y":
			expectedShape = "R"
		case "Z":
			expectedShape = "P"
		}
	case "B":
		switch result {
		case "X":
			expectedShape = "R"
		case "Y":
			expectedShape = "P"
		case "Z":
			expectedShape = "S"
		}
	case "C":
		switch result {
		case "X":
			expectedShape = "P"
		case "Y":
			expectedShape = "S"
		case "Z":
			expectedShape = "R"
		}
	}

	relationship := map[string]int{
		"R": 1,
		"P": 2,
		"S": 3,
	}

	return score + relationship[expectedShape]
}
