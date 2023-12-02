package adventofcode

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	gameNumberRegex, _ = regexp.Compile(`Game (\d+):(.*)`)
	redBallRegex, _    = regexp.Compile(`(\d+) red`)
	blueBallRegex, _   = regexp.Compile(`(\d+) blue`)
	greenBallRegex, _  = regexp.Compile(`(\d+) green`)
)

const (
	redBallMax   = 12
	greenBallMax = 13
	blueBallMax  = 14
)

func day2Part1(input []string) int {
	res := 0

	matchRelationship := map[*regexp.Regexp]int{
		redBallRegex:   redBallMax,
		greenBallRegex: greenBallMax,
		blueBallRegex:  blueBallMax,
	}
	for _, line := range input {
		gameResult := gameNumberRegex.FindStringSubmatch(line)
		gameNumber, _ := strconv.Atoi(gameResult[1])

		isValid := true

		for _, draw := range strings.Split(gameResult[2], ";") {
			for regexp, maxNumber := range matchRelationship {
				ballNumber := regexp.FindStringSubmatch(draw)
				if len(ballNumber) > 1 {
					if val, _ := strconv.Atoi(ballNumber[1]); val > maxNumber {
						isValid = false
						break
					}
				}

			}

			if !isValid {
				break
			}
		}

		if isValid {
			res += gameNumber
		}
	}

	return res

}

func day2Part2(input []string) int {
	res := 0

	matchRelationship := map[*regexp.Regexp]int{
		redBallRegex:   redBallMax,
		greenBallRegex: greenBallMax,
		blueBallRegex:  blueBallMax,
	}
	for _, line := range input {
		gameResult := gameNumberRegex.FindStringSubmatch(line)

		cubStoreRes := map[*regexp.Regexp]int{
			redBallRegex:   0,
			greenBallRegex: 0,
			blueBallRegex:  0,
		}

		for _, draw := range strings.Split(gameResult[2], ";") {
			for regexp, _ := range matchRelationship {
				ballNumber := regexp.FindStringSubmatch(draw)
				if len(ballNumber) > 1 {
					if val, _ := strconv.Atoi(ballNumber[1]); val > cubStoreRes[regexp] {
						cubStoreRes[regexp] = val
					}
				}
			}

		}

		tempRes := 1
		for _, maxBallNumber := range cubStoreRes {
			tempRes *= maxBallNumber
		}
		res += tempRes
	}

	return res

}
