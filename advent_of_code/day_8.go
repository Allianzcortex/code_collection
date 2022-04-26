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

// func day8Part2(digits []string) int {

// 	var match = map[int]string{
// 		0: "abcefg",
// 		9: "abcdfg",
// 		6: "abdefg",
// 		2: "acdeg",
// 		3: "acdfg",
// 		5: "abdfg",
// 		1: "cf",
// 		7: "acf",
// 		4: "bcdf",
// 		8: "abcdefg",
// 	}

// 	var numMatch = map[int]int{
// 		2: 1,
// 		3: 7,
// 		4: 4,
// 		7: 8,
// 	}

// 	for _, digit := range digits {
// 		line := strings.Split(digit, "|")
// 		input, output := line[0], line[1]
// 		m := make(map[byte]byte)

// 		for _, val := range strings.Fields(input) {
// 			switch len(val) {
// 			case 1, 4, 7, 8:
// 				// being able to find fixed numbers
// 				for i := 0; i < len(val); i++ {
// 					target := numMatch[len(val)]
// 					m[val[i]] = match[target][i]
// 				}
// 			}
// 		}
// 	}
// }
