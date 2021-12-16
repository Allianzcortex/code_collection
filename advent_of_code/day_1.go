package adventofcode

func day1(input []int) int {
	if input == nil {
		return 0
	}

	count := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count += 1
		}
	}

	return count
}
