package adventofcode

import (
	"fmt"
)

func day4(draws []int, boards [][][]int) int {

	maps := buildMap(boards)
	counter := make(map[string]int)

	row, column := len(boards[0][0]), len(boards[0])

	for _, num := range draws {
		for i := 0; i < len(boards); i++ {
			if _, exist := maps[i][num]; !exist {
				continue
			}
			coordinates := maps[i][num]

			row_key := fmt.Sprintf("%d_row_%d", i, coordinates[0])
			column_key := fmt.Sprintf("%d_column_%d", i, coordinates[1])
			// check row firstly
			counter[row_key] += 1
			if counter[row_key] == row {
				return buildRes(draws, boards, maps, i, num)
			}

			// check column
			counter[column_key] += 1
			if counter[column_key] == column {
				return buildRes(draws, boards, maps, i, num)
			}

		}
	}

	return -1 // if we can't find an answer
}

func buildMap(boards [][][]int) []map[int][]int {
	// classic tradeoff between time & space
	// let me see if it's in python...
	// map = [{key:f'{x}+{y}} for key,(x,y) in enumerate(boards)]
	// 2 lines most
	res := make([]map[int][]int, 0)

	for _, board := range boards {
		m := make(map[int][]int)
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				m[board[i][j]] = []int{i, j}
			}
		}

		res = append(res, m)
	}

	return res
}

func buildRes(draws []int, boards [][][]int, maps []map[int][]int, i int, num int) int {

	res := 0
	board := boards[i]
	// find unvisited
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res += board[i][j]
		}
	}

	for _, val := range draws {
		if _, exist := maps[i][val]; exist {
			if val != num {
				res -= val
			} else {
				break
			}
		}

	}
	res -= num
	return res * num
}
