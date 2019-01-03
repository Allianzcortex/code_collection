package main

import (
	"math"
)

/**
考虑三种情况：
1. base 为 0(无论 exponent 为正数或负数都定义结果为 0)
2. exponent 为 0
3. exponent 为负数

// TODO 应该把计算 unsigned int 的过程抽出来
// 这样可以实现两种计算无符号参数的方法
**/

func RunProblem11(base float64, exponent int) float64 {
	return double(base, exponent)
}

func double(base float64, exponent int) float64 {
	if isFloat64Equal(base, 0) == true {
		return 0
	}
	if exponent == 0 {
		return 1
	}

	result := base
	for i := 1; i < int(math.Abs(float64(exponent))); i++ {
		result = result * base
	}

	if exponent < 0 {
		result = 1.0 / result
	}

	return result

}

func isFloat64Equal(input1 float64, input2 float64) bool {
	var epsilon float64 = 0.0000001
	if math.Abs(input1-input2) < epsilon {
		return true
	}
	return false
}

// func main() {
// 	fmt.Println(double(0, -2))
// }
