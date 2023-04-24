package adventofcode

import (
	"regexp"
	"strconv"
	"strings"
)

type Op struct {
	val1      string
	val2      string
	operation string
}

func day21Part1(input []string) int {
	var (
		opMap  = map[string]Op{}
		numMap = map[string]int{}
	)

	for _, line := range input {
		arr := strings.Split(line, ":")
		left, right := arr[0], arr[1]
		right = strings.TrimSpace(right)
		if numVal, err := strconv.Atoi(right); err == nil {
			// it matches a number
			numMap[left] = numVal
		} else {
			// it matches a calculation
			opMap[left] = generateOp(right)
		}
	}

	return getCalculateResult(&opMap, &numMap, "root")
}

var (
	r, _ = regexp.Compile(`([a-z]{4})\s(.?)\s([a-z]{4})`)
)

func generateOp(input string) Op {
	matchedResult := r.FindStringSubmatch(input)
	return Op{
		matchedResult[1],
		matchedResult[3],
		matchedResult[2],
	}
}

func getCalculateResult(opMap *map[string]Op, numMap *map[string]int, input string) int {
	if val, ok := (*numMap)[input]; ok {
		return val
	}
	matchedOp := (*opMap)[input]
	res1 := getCalculateResult(opMap, numMap, matchedOp.val1)
	res2 := getCalculateResult(opMap, numMap, matchedOp.val2)

	res := 0
	switch matchedOp.operation {
	case "+":
		res = res1 + res2
	case "-":
		res = res1 - res2
	case "*":
		res = res1 * res2
	case "/":
		res = res1 / res2
	}

	(*numMap)[input] = res
	return res
}
