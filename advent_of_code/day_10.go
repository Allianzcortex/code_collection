package adventofcode

var match = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func day10(inputs []string) int {

	points := 0

	for _, line := range inputs {
		stack := make([]rune, 0)
		for _, ch := range line {
			if counter, exist := match[ch]; !exist {
				// if it doesn't exist,push it to stack directly
				stack = append(stack, ch)
			} else {
				// if we find it exists,check whether they're matched
				if counter == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					points += getPoint(ch)
					break
				}

			}
		}
	}

	return points

}

func getPoint(ch rune) int {
	switch ch {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return -1
	}
}
