package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func day13Part1(inputs, actions []string) int {
	m := make(map[string]bool)

	line, _ := strconv.Atoi(strings.Split(actions[0], "=")[1])

	for _, coordinate := range inputs {
		// calculate where it will be after the folding
		temp := strings.Split(coordinate, ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		if actions[0][0] == 'x' {
			// to fold to the left
			if x < line {
				m[coordinate] = true
			} else {
				x = line - (x - line)
				m[fmt.Sprintf("%d,%d", x, y)] = true
			}
		} else {
			// to fold to the up
			if y < line {
				m[coordinate] = true
			} else {
				y = line - (y - line)
				m[fmt.Sprintf("%d,%d", x, y)] = true
			}
		}
	}

	return len(m)
}
