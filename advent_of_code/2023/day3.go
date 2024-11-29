package adventofcode

import (
	"strconv"

	"golang.org/x/exp/slices"
)

type location struct {
	x int
	y int
}

var NonSymbolCharacter = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.'}

func day3Part1(input []string) int {

	res := 0
	uniqueStartLocation := map[location]int{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if slices.Contains(NonSymbolCharacter, input[i][j]) {
				continue
			}
			// so we find a valid character, we should search from 8 directions
			traverse(input, i, j, uniqueStartLocation)
		}

	}

	// iterate through all found values
	for _, val := range uniqueStartLocation {
		res += val
	}

	return res
}

func traverse(input []string, i int, j int, uniqueStartLocation map[location]int) {
	xIndex, yIndex := []int{-1, 0, 1}, []int{-1, 0, 1}
	for _, x := range xIndex {
		for _, y := range yIndex {
			if x == 0 && y == 0 {
				continue
			}
			newX, newY := i+x, j+y
			if newX < 0 || newX > len(input)-1 || newY < 0 || newY > len(input[0])-1 {
				continue
			}
			fillAdjacentNumber(input, newX, newY, uniqueStartLocation)
		}
	}

}

func fillAdjacentNumber(input []string, i int, j int, uniqueStartLocation map[location]int) {

	// if the first letter is not a number, no need to try
	if isNumber, _ := isNumber(input, i, j); !isNumber {
		return
	}

	// next find the leftMost character that is the digit
	for j >= 0 {
		if isNumber, _ := isNumber(input, i, j); isNumber {
			j -= 1
		} else {
			// well forget to J+1 here and spend a long time debugging...
			j += 1
			break
		}
	}

	j = max(0, j)
	tempRes := 0
	startPoint := location{i, j}

	// If we find this number's start point is in record alredy, then it means
	// we have calculated in the past
	if _, isFound := uniqueStartLocation[startPoint]; isFound {
		return
	}
	for j < len(input[0]) {
		isNum, val := isNumber(input, i, j)
		if isNum {
			tempRes = tempRes*10 + val
			j += 1
		} else {
			break
		}
	}

	uniqueStartLocation[startPoint] = tempRes

}

func isNumber(input []string, i int, j int) (bool, int) {

	var (
		val int
		err error
	)

	if val, err = strconv.Atoi(string(input[i][j])); err != nil {
		return false, 0
	}

	return true, val
}

// initially want to process input and fill the unique map in advance
// but it turns out to be too complex and involves too much calculatin logic

/*
func processInput(input []string) map[location]int {
	uniqueStartLocation := map[location]int{}
	for i := 0; i < len(input); {
		for j := 0; j < len(input[0]); {
			tempRes := 0
			startPoint := location{}
			if isNum, _ := isNumber(input, i, j); isNum {
				startPoint = location{i, j}
				for j <= len(input[0])-1 {
					isNum, val := isNumber(input, i, j)
					if isNum {
						tempRes = tempRes*10 + val
					} else {
						break
					}
					j += 1
				}
			}
			if tempRes != 0 {
				uniqueStartLocation[startPoint] = tempRes
			}
			if j > len(input[0])-1 {
				// should switch to next line
				i += 1
				j = 0
			} else {
				// only switch to right
				j += 1
			}
		}
	}

	return uniqueStartLocation
}
*/

func day3Part2(input []string) int {
	res := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '*' {
				continue
			}
			// so we find a star, and need to count how many are in 8 directions
			uniqueStartLocation := map[location]int{}
			xIndex, yIndex := []int{-1, 0, 1}, []int{-1, 0, 1}
			for _, x := range xIndex {
				for _, y := range yIndex {
					if x == 0 && y == 0 {
						continue
					}
					newX, newY := i+x, j+y
					if newX < 0 || newX > len(input)-1 || newY < 0 || newY > len(input[0])-1 {
						continue
					}
					fillAdjacentNumber(input, newX, newY, uniqueStartLocation)
				}
			}
			// after iterating all indexes, process input
			res += processEngineNumber(uniqueStartLocation)
		}
	}

	return res
}

// A gear is any * symbol that is adjacent to exactly two part numbers
func processEngineNumber(uniqueStartLocation map[location]int) int {

	if len(uniqueStartLocation) != 2 {
		return 0
	}

	res := 1
	for _, val := range uniqueStartLocation {
		res *= val
	}
	return res

}
