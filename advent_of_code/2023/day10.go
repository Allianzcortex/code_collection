package adventofcode

import (
	"math"
	"strconv"
	"strings"
)

func processInput10(input []string) (map[location][]string, location) {

	match := make(map[location][]string, len(input)*len(input[0]))
	startLocation := location{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			lc := location{x: i, y: j}
			// can use another map here to assign value directly
			// but haven't used switch for a long time
			switch input[i][j] {
			case '|':
				match[lc] = []string{"N->N", "S->S"} // N and S
			case '-':
				match[lc] = []string{"E->E", "W->W"} // E and W
			case 'L':
				match[lc] = []string{"S->E", "W->N"} // North and East
			case 'J':
				match[lc] = []string{"S->W", "E->N"} // North and West
			case '7':
				match[lc] = []string{"E->S", "N->W"}
			case 'F':
				match[lc] = []string{"N->E", "W->S"}
			case '.':
				match[lc] = []string{}
			case 'S':
				startLocation.x = i
				startLocation.y = j
				match[lc] = []string{}

			}
		}
	}

	return match, startLocation
}

func day10Part1(input []string) int {

	match, startLocation := processInput10(input)
	visited := make(map[location]int, len(input)*len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			visited[location{i, j}] = math.MaxInt32
		}
	}

	for _, action := range []string{"N", "S", "W", "E"} {
		traverse10(input, match, visited, startLocation.x, startLocation.y, action, 0)
	}

	res := math.MinInt32
	for _, val := range visited {
		if val != math.MaxInt32 {
			res = max(res, val)
		}
	}
	return res
}

func traverse10(input []string, match map[location][]string, visited map[location]int, x, y int, dir string, curRes int) {

	// for '.'
	newX, newY := x, y
	switch dir {
	case "N":
		newX -= 1
	case "S":
		newX += 1
	case "W":
		newY -= 1
	case "E":
		newY += 1
	}

	if newX < 0 || newX > len(input)-1 || newY < 0 || newY > len(input[0])-1 {
		return
	}

	nextDir := match[location{newX, newY}]
	if len(nextDir) == 0 {
		return
	}
	// if it has been visited
	if visited[location{x, y}] != math.MaxInt32 && visited[location{x, y}] > curRes {
		return
	}

	action := ""
	for _, ac := range nextDir {
		if string(ac[0]) == dir {
			action = string(ac[3])
			break
		}
	}
	// dir not match
	if action == "" {
		return
	}

	visited[location{newX, newY}] = min(visited[location{newX, newY}], curRes+1)

	traverse10(input, match, visited, newX, newY, action, curRes+1)
}

func day10Part2(input []string) int {

	res := 0

	for _, line := range input {

		// with Python list comprehension,it only needs 1-line...
		numberStr := strings.Fields(line)
		nums := []int{}
		for _, numStr := range numberStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		// next store all last numbers
		firstNumbers := []int{nums[0]}
		for {
			nextStepNumbers := make([]int, len(nums)-1)
			for i := 0; i < len(nums)-1; i++ {
				nextStepNumbers[i] = nums[i+1] - nums[i]
			}
			if isAllZero(nextStepNumbers) {
				break
			} else {
				firstNumbers = append(firstNumbers, nextStepNumbers[0])
				nums = nextStepNumbers
			}
		}

		// now do the calculation
		// assume previous val = 0
		previousVal := 0
		for i := len(firstNumbers) - 1; i >= 0; i-- {
			previousVal = firstNumbers[i] - previousVal
		}
		res += previousVal
	}

	return res
}
