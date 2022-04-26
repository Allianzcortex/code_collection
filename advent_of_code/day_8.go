package adventofcode

import "strings"

/**

0:abcefg
9:abcdfg
6:abdefg

2:acdeg
3:acdfg
5:abdfg

1:cf
7:acf
4:bcdf
8:abcdefg

**/

func day8(digits []string) int {

	// not intuitive ehough, should use swithch/if-else..
	var match = map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}

	res := 0

	for _, digit := range digits {
		line := strings.Split(digit, "|")
		output := line[1]
		for _, val := range strings.Fields(output) {
			if _, exist := match[len(val)]; exist {
				res += 1
			}
		}
	}

	return res
}
