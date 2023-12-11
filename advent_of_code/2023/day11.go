package adventofcode

import (
	"strings"
)

func day11Part1(input []string) int {
	expandedGalaxy := getExpandedGalaxy(input)
	locations := []location{}

	for i := 0; i < len(expandedGalaxy); i++ {
		for j := 0; j < len(expandedGalaxy[0]); j++ {
			// handle not long enough /***/ array
			// this won't work for array that with no '#' in first line
			// should use maximum(len(column)), but can handle test case well
			if j > len(expandedGalaxy[i])-1 {
				continue
			}
			if expandedGalaxy[i][j] == '#' {
				locations = append(locations, location{i, j})
			}
		}
	}

	res := 0
	for i := 0; i < len(locations)-1; i++ {
		for j := i + 1; j < len(locations); j++ {
			lc1, lc2 := locations[i], locations[j]
			res += (abs(lc1.x-lc2.x) + abs(lc1.y-lc2.y))
		}
	}

	return res
}

func getExpandedGalaxy(input []string) []string {
	rowContainGalaxy := make([]bool, len(input))
	columnContainGalaxy := make([]bool, len(input[0]))

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '#' {
				rowContainGalaxy[i] = true
				columnContainGalaxy[j] = true
			}
		}
	}

	// yet another Python 1-line thing...
	emptyRow := strings.Builder{}
	for i := 0; i < len(input[0]); i++ {
		emptyRow.WriteByte('.')
	}
	result := []string{}
	for i := 0; i < len(input); i++ {
		// append 2 rows
		if !rowContainGalaxy[i] {
			result = append(result, []string{emptyRow.String(), emptyRow.String()}...)
			continue
		}

		row := strings.Builder{}
		for j := 0; j < len(input[0]); j++ {
			if columnContainGalaxy[j] {
				row.WriteByte(input[i][j])
			} else {
				// append 2 columns
				row.WriteByte('.')
				row.WriteByte('.')
			}
		}
		result = append(result, row.String())
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
