package main

func checkMatrix(matrix [][]int, target int) []int {
	x, y := len(matrix)-1, len(matrix[0])-1
	flag := false
	for {

		if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0])-1 {
			break
		}

		if matrix[x][y] == target {
			flag = true
			break
		}
		if matrix[x][y] > target {

		} else {

		}
	}

	if flag == true {
		return []int{x, y}
	} else {
		return []int{-1, -1}
	}

}

func main() {
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	checkMatrix(matrix, 7)
}
