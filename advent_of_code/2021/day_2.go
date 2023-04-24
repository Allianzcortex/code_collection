package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func day2(input []string) int {

	horizon, depth := 0, 0

	for _, order := range input {
		temp := strings.Split(order, "")
		if len(temp) != 2 {
			fmt.Errorf("Invalid line %v", order)
		}
		dir, val := temp[0], temp[1]
		distance, _ := strconv.Atoi(val)

		switch dir {
		case "forward":
			horizon += distance
		case "down":
			depth += distance
		case "up":
			depth -= distance
		default:
			fmt.Errorf("Invalid command %s", dir)
		}
	}

	return horizon * depth
}
