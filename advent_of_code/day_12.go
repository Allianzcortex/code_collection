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
	output := "start "
	for _, location := range maps["start"] {
		travel(output, location, maps, map[string]bool{"start": true}, &cnt)
	}

	return cnt
}

func travel(output string, location string, maps map[string][]string, visited map[string]bool, cnt *int) {
	if location == "end" {
		*cnt += 1
		fmt.Println(output + " end")
		return
	}

	if isSmallCave(location) {
		visited[location] = true
	}
	output += location + " "

	for _, next := range maps[location] {
		if _, hasVisited := visited[next]; hasVisited {
			continue
		}
		travel(output, next, maps, visited, cnt)
		output = output[:len(output)-1]
		delete(visited, next)
	}
}

func isSmallCave(location string) bool {
	for _, ch := range location {
		if ch == unicode.ToLower(ch) {
			return true
		}
	}

	return false
}
