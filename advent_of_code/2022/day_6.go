package adventofcode

func day6Part1(input string) int {
	m := map[byte]int{}

	for index := 0; index < len(input); index++ {
		ch := input[index]
		if index >= 4 {
			m[input[index-4]] -= 1
			if m[input[index-4]] == 0 {
				delete(m, input[index-4])
			}
		}

		if _, exist := m[ch]; !exist {
			if len(m) == 3 {
				return index + 1
			}
		}
		m[ch] += 1
	}

	return -1
}
