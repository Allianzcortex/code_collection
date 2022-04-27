package adventofcode

import "fmt"

func day11(matrix [][]int) int {

	flashCnt := 0
	hasFlashed := make(map[string]bool)
	for step := 0; step < 100; step++ {
		// 1. increase 1 for every Octo
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] += 1
			}
		}
		// traverse through matrix
		// since there is no append operation so no need to use pointer

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[i][j] > 9 {
					flash(matrix, i, j, hasFlashed, step, &flashCnt)
				}
			}
		}
	}

	return flashCnt

}

func flash(matrix [][]int, i int, j int, hasFlashed map[string]bool, step int, flashCnt *int) {

	hasFlashed[key(i, j, step)] = true
	*flashCnt += 1
	matrix[i][j] = 0

	// all adjacent engery + 1
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			// don't recalculate itself
			if x == 0 && y == 0 {
				continue
			}
			newX, newY := x+i, y+j
			// check whether it's out of bound
			if newX < 0 || newX > len(matrix)-1 || newY < 0 || newY > len(matrix[0])-1 {
				continue
			}
			// if it has flashed previously, do nothing
			if _, exist := hasFlashed[key(newX, newY, step)]; exist {
				continue
			}
			// finally + 1
			matrix[newX][newY] += 1

			// if it exceeds ,reflash again
			if (matrix[newX][newY]) > 9 {
				flash(matrix, newX, newY, hasFlashed, step, flashCnt)
			}
		}
	}

}

// a pitfall here , key must be related with `step`
// i.e. it means : the Octo just can't flash in this round
// but it can still flash in other rounds
func key(x, y, step int) string {
	return fmt.Sprintf("%d:%d:%d", x, y, step)
}
