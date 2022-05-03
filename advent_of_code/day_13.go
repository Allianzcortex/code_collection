package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func day13(inputs, actions []string) int {
	m := make(map[string]bool)

	for _, coordinate := range inputs {
		m[coordinate] = true
	}

	var maxX, maxY int

	for _, action := range actions {
		maxX, maxY = 0, 0
		line, _ := strconv.Atoi(strings.Split(action, "=")[1])
		for coordinate, _ := range m {
			x, y := parseCoordinate(coordinate)
			maxX = max(maxX, x)
			maxY = max(maxY, y)
			if action[0] == 'x' {
				// to fold to the left
				if x > line {
					// good to know it's safe to delete key during the iteration
					delete(m, coordinate)
					x = line - (x - line)
					m[fmt.Sprintf("%d,%d", x, y)] = true
				}
			} else {
				// to fold to the up
				if y > line {
					delete(m, coordinate)
					y = line - (y - line)
					m[fmt.Sprintf("%d,%d", x, y)] = true
				}
			}
		}

	}

	// pitfall here
	// should notice how it defines x,y ...
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			coordinate := fmt.Sprintf("%d,%d", j, i)
			if _, exist := m[coordinate]; exist {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		// **    ** *  *  **  **** *  * *  * *
		// *  *    * *  * *  *    * *  * * *  *
		// *       * **** *  *   *  **** **   *
		// *       * *  * ****  *   *  * * *  *
		// *  * *  * *  * *  * *    *  * * *  *
		//  **   **  *  * *  * **** *  * *  *  **
	}

	return len(m)
}

func parseCoordinate(coordinate string) (int, int) {
	temp := strings.Split(coordinate, ",")
	x, _ := strconv.Atoi(temp[0])
	y, _ := strconv.Atoi(temp[1])
	return x, y
}
