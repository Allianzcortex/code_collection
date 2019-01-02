import "fmt"

/**
A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.

For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Example:

Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
Note:
The order of output does not matter.
The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".

**/
var hours = [4]int{1, 2, 4, 8}
var minutes = [6]int{1, 2, 4, 8, 16, 32}

// 给定 hour 的 n 值，返回所有可能的结果
func traverseHour(res *[]int, num int, index int, result int) {
	if result >= 12 {
		return
	}
	if num == 0 || index > 4 {
		*res = append(*res, result)
		return
	}
	for i := index; i < 4; i++ {
		traverseHour(res, num-1, i+1, result+hours[i])
	}
}

func traverseMinute(res *[]int, num int, index int, result int) {
	if result > 59 {
		return
	}
	// >=6 还是 >6 这里没想清
	if num == 0 || index > 6 {
		*res = append(*res, result)
		return
	}
	for i := index; i < 6; i++ {
		traverseMinute(res, num-1, i+1, result+minutes[i])
	}
}

func readBinaryWatch(num int) []string {

	result := []string{}
	// No need to check this situation
	// if(num==0){
	//     return []string{"0:00"}
	// }
	for i := 0; i <= num; i++ {
		hours_result := []int{}
		traverseHour(&hours_result, i, 0, 0)
		//fmt.Println(hours_result,3,0,0)
		minutes_result := []int{}
		traverseMinute(&minutes_result, num-i, 0, 0)

		for _, hour := range hours_result {
			for _, minute := range minutes_result {
				result = append(result, fmt.Sprintf("%d:%02d", hour, minute))
			}
		}

	}
	// 计算
	return result
}
