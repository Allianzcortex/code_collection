package main

/**
给定一组数字，判断是否为二叉搜索树的后序遍历序列

**/

func verifySequenceOfBST(sequence []int, start int, end int) bool {
	if start == end {
		return true
	}

	var i int
	target := sequence[end]
	for i = start; i < end; i++ {
		if sequence[i] > target {
			break
		}
	}

	for j := i; j < end; j++ {
		if sequence[j] < target {
			return false
		}
	}

	return verifySequenceOfBST(sequence, start, i-1) && verifySequenceOfBST(sequence,
		i, end-1)
}

func RuntProblem26(sequence []int) bool {
	return verifySequenceOfBST(sequence, 0, len(sequence)-1)
}

// uncomment it if you want to run this file
// func main() {
// 	// input := []int{5, 7, 6, 9, 11, 10, 8}
// 	input := []int{7, 4, 6, 5}
// 	fmt.Println(verifySequenceOfBST(input, 0, len(input)-1))

// }
