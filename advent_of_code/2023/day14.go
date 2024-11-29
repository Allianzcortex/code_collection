package adventofcode

func day14Part1(puzzle []string) int {
	input := make([][]byte, len(puzzle))
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			input[i] = append(input[i], puzzle[i][j])
		}
	}
	loadRockToNorth(input)

	// next calculate result
	res := 0
	for i := 0; i < len(input); i++ {
		ratio := len(input) - i
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'O' {
				res += ratio
			}
		}
	}

	return res
}

func loadRockToNorth(input [][]byte) {
	row, column := len(input), len(input[0])
	for j := 0; j < column; j++ {
		for i := 0; i < row; i++ {
			if input[i][j] == '.' || input[i][j] == '#' {
				continue
			}
			// if it's a rock O
			index := i - 1
			for index >= 0 {
				if input[index][j] == '#' || input[index][j] == 'O' {
					break
				}
				if input[index][j] == '.' {
					input[index][j] = 'O'
					input[index+1][j] = '.'
				}
				index -= 1
			}
		}
	}
}
