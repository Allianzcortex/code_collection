package adventofcode

func day6Part1(fishes []int) int {
	for days := 0; days < 80; days++ {
		res := make([]int, len(fishes))
		copy(res, fishes)

		for index, f := range fishes {
			if f == 0 {
				res[index] = 6
				res = append(res, 8)
			} else {
				res[index] -= 1
			}
		}
		fishes = res

		if days == 79 {
			return len(res)
		}
	}

	return -1
}

func day6Part2(fishes []int) int {
	var arr [9]int

	for _, f := range fishes {
		arr[f] += 1
	}

	for days := 0; days < 256; days++ {
		zeroCount := arr[0]
		for index := 1; index <= 8; index++ {
			arr[index-1] = arr[index]
			if index == 7 {
				arr[index-1] += zeroCount
			}
		}
		arr[8] = zeroCount
	}

	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
