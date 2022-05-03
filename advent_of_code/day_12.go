package adventofcode

import (
	"fmt"
	"strings"
	"unicode"
)

func day12(paths []string) int {
	maps := make(map[string][]string)
	isSmallCaveMaps := make(map[string]bool)

	// build map
	for _, path := range paths {
		temp := strings.Split(path, "-")
		start, end := temp[0], temp[1]
		maps[start] = append(maps[start], end)
		maps[end] = append(maps[end], start)

		isSmallCaveMaps[start] = isSmallCave(start)
		isSmallCaveMaps[end] = isSmallCave(end)
	}

	fmt.Println(maps)

	// begin to traverse
	cnt := 0
	output := []string{"start"}
	isVisited := false
	for _, location := range maps["start"] {
		travel(&output, isVisited, location, maps, map[string]bool{"start": false}, &cnt)
	}

	return cnt
}

func travel(output *[]string, isVisitedTwice bool, location string, maps map[string][]string, visited map[string]bool, cnt *int) {
	if location == "end" {
		*cnt += 1
		fmt.Println(*output)
		return
	}

	if isSmallCave(location) {
		visited[location] = true
	}
	*output = append(*output, location)
	val := isVisitedTwice
	for _, next := range maps[location] {
		// tempVisited := *isVisitedTwice
		if hasVisited, exist := visited[next]; exist {
			if !hasVisited {
				// handle "start" especially
				continue
			}
			if !isVisitedTwice {
				val = true
				travel(output, val, next, maps, visited, cnt)
				// *isVisitedTwice = false
			} else {
				continue
			}
		} else {
			travel(output, val, next, maps, visited, cnt)
		}
	}
	// *isVisitedTwice = val
	*output = (*output)[:len(*output)-1]
	delete(visited, location)

}

func isSmallCave(location string) bool {
	for _, ch := range location {
		if ch == unicode.ToLower(ch) {
			return true
		}
	}

	return false
}
