package adventofcode

import (
	"fmt"
	"sort"
)

var lowPoint = &[][]int{}

func day9Part1(matrix [][]int) int {
	row, column := len(matrix), len(matrix[0])

	cnt := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			// up
			if i != 0 && matrix[i][j] >= matrix[i-1][j] {
				continue
			}
			// down
			if i != row-1 && matrix[i][j] >= matrix[i+1][j] {
				continue
			}
			// left
			if j != 0 && matrix[i][j] >= matrix[i][j-1] {
				continue
			}
			// right
			if j != column-1 && matrix[i][j] >= matrix[i][j+1] {
				continue
			}
			*lowPoint = append(*lowPoint, []int{i, j})
			cnt += (1 + matrix[i][j])
		}
	}

	return cnt

}

func day9Part2(matrix [][]int) int {
	day9Part1(matrix)
	// now get all small coordinates & resue function 1
	result := make([]int, len(*lowPoint))
	visited := make(map[string]bool)
	for i := 0; i < len(result); i++ {
		x, y := (*lowPoint)[i][0], (*lowPoint)[i][1]
		area := 0
		traverse(matrix, x, y, visited, &area)
		result[i] = area
	}

	sort.Ints(result)
	output := result[len(result)-3:]
	sum := 1
	for _, val := range output {
		sum *= val
	}
	return sum
}

func traverse(matrix [][]int, x int, y int, visited map[string]bool, area *int) {
	if x < 0 || x > len(matrix)-1 || y < 0 || y > len(matrix[0])-1 {
		return
	}
	key := fmt.Sprintf("%d:%d", x, y)
	if _, hasSeen := visited[key]; hasSeen || matrix[x][y] == 9 {
		return
	}

	visited[key] = true
	*area += 1
	// well , don't expect a pitfall here,how long I haven't practiced DFS
	for _, coordinates := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		i, j := coordinates[0], coordinates[1]
		traverse(matrix, x+i, y+j, visited, area)
	}
}
