package adventofcode

import (
	"math"
	"strings"
)

func getFormulaMap(formula []string) map[string]string {
	m := make(map[string]string)
	for _, form := range formula {
		temp := strings.Split(form, "->")
		key := strings.TrimSpace(temp[0])
		val := strings.TrimSpace(temp[1])
		m[key] = val
	}

	return m
}

func getFreq(counter map[byte]int) (int, int) {
	mostE, leastE := math.MinInt, math.MaxInt
	for _, freq := range counter {
		if freq == 0 {
			// when freq==0,it doesn't influence mostE,but leaseE will
			// always become zero in such a comparison case
			continue
		}
		mostE = max(mostE, freq)
		leastE = min(leastE, freq)
	}

	return mostE, leastE

}

// In part2, we can't simulate to generate all strings
// Instead, we should count the frequency the pairs
func day14Part2(elements string, formula []string) int {
	firstCH, lastCH := elements[0], elements[len(elements)-1]
	m := getFormulaMap(formula)
	res := make(map[string]int)

	for i := 0; i < len(elements)-1; i++ {
		res[elements[i:i+2]] += 1
	}

	for step := 0; step < 40; step++ {
		tempRes := make(map[string]int)
		for key, val := range res {
			target := m[key]
			tempRes[string(key[0])+target] += val
			tempRes[target+string(key[1])] += val
		}
		res = tempRes
	}

	// next,counter
	counter := make(map[byte]int)
	for key, freq := range res {
		counter[key[0]] += freq
		counter[key[1]] += freq
	}

	for ch, freq := range counter {
		if ch == firstCH || ch == lastCH {
			counter[ch] = (freq-1)/2 + 1
		} else {
			counter[ch] = freq / 2
		}
	}
	maxE, minE := getFreq(counter)
	return maxE - minE
}

func day14Part1(elements string, formula []string) int {

	// fristly, convert formula to map-based structure
	m := getFormulaMap(formula)

	// secondly, iterate
	for step := 0; step < 10; step++ {
		var res strings.Builder
		for i := 0; i < len(elements)-1; i++ {
			// pitfall : when converting []byte/[]rune to stirng,using
			// fmt.Sprintf is not always a good choice
			target := m[string([]byte{elements[i], elements[i+1]})]
			res.WriteByte(elements[i])
			res.WriteString(target)
		}
		res.WriteByte(elements[len(elements)-1])
		elements = res.String()
	}

	// lastly, Get the most common and least common
	// counter := [26]int{}
	counter := make(map[byte]int)
	for index := 0; index < len(elements); index++ {
		counter[elements[index]] += 1
	}
	mostE, leastE := getFreq(counter)
	return mostE - leastE

}
