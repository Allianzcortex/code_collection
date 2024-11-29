package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

func day9Part1(input []string) int {

	var matrix = [2000][2000]int{}
	row, column := len(matrix), len(matrix[0])
	hx, hy := 1000, 1000
	tx, ty := 1000, 1000
	matrix[hx][hy] = 1 // 1 means it has been visited by tail

	for _, val := range input {
		order := strings.Fields(val)
		action := order[0]
		step, _ := strconv.Atoi(order[1])

		switch action {
		case "R":
			for i := 1; i <= step; i++ {
				hx += 1
				move(&matrix, hx, hy, &tx, &ty)
			}
		case "L":
			for i := step; i >= 1; i-- {
				hx -= 1
				move(&matrix, hx, hy, &tx, &ty)
			}
		case "U":
			for i := 1; i <= step; i++ {
				hy += 1
				move(&matrix, hx, hy, &tx, &ty)
			}
		case "D":
			for i := 1; i <= step; i++ {
				hy -= 1
				move(&matrix, hx, hy, &tx, &ty)
			}
		}

	}

	cnt := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if (matrix[i][j]) == 1 {
				cnt += 1
			}
		}
	}

	return cnt
}

func move(matrix *[2000][2000]int, hx, hy int, tx, ty *int) {
	fmt.Printf("%d %d\n", hx, hy)
	fmt.Printf("%d %d\n", *tx, *ty)
	xDis, yDis := hx-*tx, hy-*ty
	if abs(xDis)+abs(yDis) == 1 {
		return
	}
	if abs(xDis) == 1 && abs(yDis) == 1 {
		return
	}

	// move up,left,right
	if abs(xDis) == 2 && abs(yDis) == 0 {
		*tx += (hx - *tx) / 2
	}
	if abs(yDis) == 2 && abs(xDis) == 0 {
		*ty += (hy - *ty) / 2
	}
	// next will be for diagonals
	// maybe there are simpler ways, but I will just list all 4 conditions
	// 1. head is right-top of tail
	if hx > *tx && hy > *ty {
		*tx, *ty = *tx+1, *ty+1
	}
	// 2. head is right-bottom of tail
	if hx > *tx && hy < *ty {
		*tx, *ty = *tx+1, *ty-1
	}
	// 3. head is left-top of tail
	if hx < *tx && hy > *ty {
		*tx, *ty = *tx-1, *ty+1
	}
	// 4. head is left-bottom of tail
	if hx < *tx && hy < *ty {
		*tx, *ty = *tx-1, *ty-1
	}
	(*matrix)[*tx][*ty] = 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
