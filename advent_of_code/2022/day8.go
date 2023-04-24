package adventofcode

import (
	"strconv"
)

func day8Part1(input []string) int {
	row, column := len(input), len(input[0])
	topMax := make([][]int, row)
	downMax := make([][]int, row)
	leftMax := make([][]int, row)
	rightMax := make([][]int, row)

	for i := 0; i < row; i++ {
		topMax[i] = make([]int, column)
		downMax[row-i-1] = make([]int, column)
	}

	for j := 0; j < column; j++ {
		for i := 0; i < row; i++ {
			if i == 0 {
				continue
			}

			topVal, _ := strconv.Atoi(string(input[i-1][j]))
			topMax[i][j] = max(topMax[i-1][j], topVal)

			downVal, _ := strconv.Atoi(string(input[row-i][j]))
			// need to take a look at it later
			downMax[row-i-1][j] = max(downMax[row-i][j], downVal)

		}
	}

	for i := 0; i < row; i++ {
		leftMax[i] = make([]int, column)
		rightMax[i] = make([]int, column)
		for j := 0; j < column; j++ {
			if j == 0 {
				continue
			}
			leftVal, _ := strconv.Atoi(string(input[i][j-1]))
			leftMax[i][j] = max(leftMax[i][j-1], leftVal)

			rightVal, _ := strconv.Atoi(string(input[i][column-j]))
			rightMax[i][column-j-1] = max(rightMax[i][column-j], rightVal)
		}
	}

	cnt := 0
	for i := 1; i < row-1; i++ {
		for j := 1; j < column-1; j++ {
			val, _ := strconv.Atoi(string(input[i][j]))
			if val > topMax[i][j] || val > downMax[i][j] || val > leftMax[i][j] || val > rightMax[i][j] {
				cnt += 1
			}
		}
	}

	return cnt + (row+column)*2 - 4
}
