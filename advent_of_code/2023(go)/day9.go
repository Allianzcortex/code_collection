package adventofcode

import (
	"strconv"
	"strings"
)

func day9Part1(input []string) int {

	res := 0

	for _, line := range input {

		// with Python list comprehension,it only needs 1-line...
		numberStr := strings.Fields(line)
		nums := []int{}
		for _, numStr := range numberStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		// next store all last numbers
		lastNumbers := []int{nums[len(nums)-1]}
		for {
			nextStepNumbers := make([]int, len(nums)-1)
			for i := 0; i < len(nums)-1; i++ {
				nextStepNumbers[i] = nums[i+1] - nums[i]
			}
			if isAllZero(nextStepNumbers) {
				break
			} else {
				lastNumbers = append(lastNumbers, nextStepNumbers[len(nextStepNumbers)-1])
				nums = nextStepNumbers
			}
		}

		for _, lastNumbe := range lastNumbers {
			res += lastNumbe
		}
	}

	return res
}

// with Python also one-line...
func isAllZero(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			return false
		}
	}
	return true
}

func day9Part2(input []string) int {

	res := 0

	for _, line := range input {

		// with Python list comprehension,it only needs 1-line...
		numberStr := strings.Fields(line)
		nums := []int{}
		for _, numStr := range numberStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		// next store all last numbers
		firstNumbers := []int{nums[0]}
		for {
			nextStepNumbers := make([]int, len(nums)-1)
			for i := 0; i < len(nums)-1; i++ {
				nextStepNumbers[i] = nums[i+1] - nums[i]
			}
			if isAllZero(nextStepNumbers) {
				break
			} else {
				firstNumbers = append(firstNumbers, nextStepNumbers[0])
				nums = nextStepNumbers
			}
		}

		// now do the calculation
		// assume previous val = 0
		previousVal := 0
		for i := len(firstNumbers) - 1; i >= 0; i-- {
			previousVal = firstNumbers[i] - previousVal
		}
		res += previousVal
	}

	return res
}
