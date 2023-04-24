package adventofcode

import (
	"regexp"
	"strconv"
	"strings"
)

func day5Part1(crates, steps []string) string {
	// step 1 : get crates count
	var lastLine = crates[len(crates)-1]
	var lastCharacter = lastLine[len(lastLine)-2]
	cratesCount, _ := strconv.Atoi(string(lastCharacter))

	// step 2 : mock crates
	stacks := [][]byte{}
	for i := 0; i < cratesCount; i++ {
		stacks = append(stacks, []byte{})
		// x-axis index
		for j := len(crates) - 2; j >= 0; j-- {
			// y-axis index,for `[D] `, it will take 4 spaces
			k := 4*i + 1
			target := crates[j][k]
			if target == ' ' {
				break
			}
			stacks[i] = append(stacks[i], target)
		}
	}

	// step 3 : follow steps
	r, _ := regexp.Compile(`\D+(\d+)\D+(\d+)\D+(\d+)`)
	for _, step := range steps {
		matchedResult := r.FindStringSubmatch(step)

		moveCount, _ := strconv.Atoi(matchedResult[1])
		from, _ := strconv.Atoi(matchedResult[2])
		to, _ := strconv.Atoi(matchedResult[3])

		from -= 1
		to -= 1

		for i := 0; i < moveCount; i++ {
			if len(stacks[from]) == 0 {
				break
			}
			// I miss Python... stacks[from].push(stacks.pop())
			moveChar := stacks[from][len(stacks[from])-1]
			stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to] = append(stacks[to], moveChar)
		}
	}

	var res strings.Builder
	for i := 0; i < cratesCount; i++ {
		res.WriteByte(stacks[i][len(stacks[i])-1])
	}
	return res.String()
}
