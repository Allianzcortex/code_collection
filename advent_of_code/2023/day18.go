package adventofcode

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func day18Part1(input []string) int {
	// prepare a grid as input, lets assume maximum length is 500

	var (
		gridRow    = 1000
		gridColumn = 1000
		startX     = 300
		startY     = 300
	)
	matrix := make([][]byte, gridRow)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]byte, gridColumn)
	}

	matrix[startX][startY] = '#'

	for _, line := range input {
		parsedInput := strings.Fields(line)

		// do not need to handle trench color in p1
		dir := parsedInput[0]
		length, _ := strconv.Atoi(parsedInput[1])

		for i := 0; i < length; i++ {
			switch dir {
			case "R":
				startY += 1
			case "D":
				startX += 1
			case "U":
				startX -= 1
			case "L":
				startY -= 1
			}
			matrix[startX][startY] = '#'
		}
	}

	// Next need to dig out the interior area
	index := 1
	indexArray := []byte{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if traverse18(matrix, byte(index), i, j) {
					indexArray = append(indexArray, byte(index))
				}
				index += 1
			}
		}
	}
	// we have filled all the array we want, now recalculate
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '#' || slices.Contains(indexArray, matrix[i][j]) {
				res += 1
			}
		}
	}
	return res
}

// bool will return whether it's inside the trench or outside
func traverse18(grid [][]byte, ch byte, x int, y int) bool {
	if x < 0 || x > len(grid)-1 || y < 0 || y > len(grid[0])-1 {
		return false
	}
	if (grid[x][y]) == '#' {
		return true
	}
	// this has been visited in the past
	if grid[x][y] != 0 {
		return true
	}

	// mark this has been visited
	grid[x][y] = ch
	isInsideTrench := true
	indexes := []location{location{-1, 0}, location{1, 0}, location{0, -1}, location{0, 1}}
	for _, idx := range indexes {
		// well, should traverse firstly and then judge
		// if use isInsideTrench && traverse18(grid, ch, x+idx.x, y+idx.y)
		// then many directions won't be explored, still work but cause a waste of time
		isInsideTrench = traverse18(grid, ch, x+idx.x, y+idx.y) && isInsideTrench

	}
	return isInsideTrench
}
