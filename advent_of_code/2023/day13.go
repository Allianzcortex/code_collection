package adventofcode

import (
	"fmt"
	"strings"
)

func day13Part1(input []string) int {

	parsedInput := processInput13(input)

	res := 0
	for _, matrix := range parsedInput {
		// find how many rows are in-mirror
		for i := 0; i < len(matrix)-1; i++ {
			if isArrayMirrorInRow(matrix, i) {
				res += 100 * (i + 1)
			}
		}

		// transpose and check how many rows are in-mirror,
		// this will be equal to how many columns are in-mirror

		transposedMatrix := transposeMatrix(matrix)
		fmt.Println("==transposed matrix is ===")
		for i := 0; i < len(transposedMatrix); i++ {
			for j := 0; j < len(transposedMatrix[0]); j++ {
				fmt.Print(string(transposedMatrix[i][j]))
			}
			fmt.Println()
		}
		for i := 0; i < len(transposedMatrix)-1; i++ {
			if isArrayMirrorInRow(transposedMatrix, i) {
				res += (i + 1)
			}
		}
	}

	return res
}

func processInput13(input []string) [][]string {
	output := [][]string{}

	tempInput := []string{}
	for i := 0; i < len(input); i++ {
		if len(input[i]) == 0 {
			output = append(output, tempInput)
			tempInput = []string{}
			continue
		} else {
			tempInput = append(tempInput, input[i])
		}
	}
	output = append(output, tempInput)
	return output
}

func isArrayMirrorInRow(input []string, index int) bool {
	prevIndex, nextIndex := index, index+1
	steps := min(prevIndex, len(input)-nextIndex-1)
	for i := 0; i <= steps; i++ {
		line1, line2 := input[prevIndex-i], input[nextIndex+i]
		for j := 0; j < len(line1); j++ {
			if line1[j] != line2[j] {
				return false
			}
		}
	}

	return true
}

/*
*

1 2 3 4
5 6 7 8
9 10 11 12

1 5 9
2 6 10
3 7 11
4 8 12

*
*/
func transposeMatrix(matrix []string) []string {
	row, column := len(matrix), len(matrix[0])
	res := make([]string, column)

	for j := 0; j < column; j++ {
		temp := strings.Builder{}
		for i := 0; i < row; i++ {
			temp.WriteByte(matrix[i][j])
		}
		res[j] = temp.String()
	}
	return res
}
