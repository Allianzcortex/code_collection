package adventofcode

import (
	"fmt"
	"regexp"
	"strconv"
)

var pattern *regexp.Regexp = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func day5(lines []string) int {

	cnt := 0
	counter := make(map[string]int)
	for _, line := range lines {
		match := pattern.FindStringSubmatch(line)
		startX, _ := strconv.Atoi(match[1])
		startY, _ := strconv.Atoi(match[2])
		endX, _ := strconv.Atoi(match[3])
		endY, _ := strconv.Atoi(match[4])

		if startX == endX && startY == endY {
			AddCoordinate(startX, startY, counter, &cnt)
			continue
		}

		// part 2 : add diagonal line
		if abs(startX, endX) == abs(startY, endY) {
			var stepX, stepY, steps int
			if endX-startX > 0 {
				stepX = 1
			} else {
				stepX = -1
			}

			if endY-startY > 0 {
				stepY = 1
			} else {
				stepY = -1
			}
			steps = abs(startX, endX)

			// we don't need to care about whether it should be startX<=endX or startX>=endX
			// in different situations. we only need to know how many tims we need to iterate
			for i := 0; i <= steps; i++ {
				AddCoordinate(startX, startY, counter, &cnt)
				startX += stepX
				startY += stepY
			}
			continue
		}

		// a trap here,there is no guarantee that startY is always smaller than endY
		// should find min/max value
		if startX == endX {
			for i := min(startY, endY); i <= max(startY, endY); i++ {
				AddCoordinate(startX, i, counter, &cnt)
			}
		}

		if startY == endY {
			for i := min(startX, endX); i <= max(startX, endX); i++ {
				AddCoordinate(i, startY, counter, &cnt)
			}
		}
	}

	return cnt

}

func AddCoordinate(x int, y int, counter map[string]int, cnt *int) {
	key := fmt.Sprintf("%d:%d", x, y)
	if val, exist := counter[key]; !exist {
		counter[key] += 1
	} else {
		if val == 1 {
			*cnt += 1
		}
		counter[key] += 1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
