import "unicode"

/**
Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.  Return a list of all possible strings we could create.

Examples:
Input: S = "a1b2"
Output: ["a1b2", "a1B2", "A1b2", "A1B2"]

Input: S = "3z4"
Output: ["3z4", "3Z4"]

Input: S = "12345"
Output: ["12345"]
Note:

S will be a string with length between 1 and 12.
S will consist only of letters or digits.

**/
func convert(s string, index int) string {
	runes := []rune(s)
	// 下面是自己一开始实现的逻辑，只比 33% 的人快
	// if runes[index] >= 'a' && runes[index] <= 'z' {
	// 	runes[index] = runes[index] - ('a' - 'A')
	// } else {
	// 	runes[index] = runes[index] + ('a' - 'A')
	// }
	// 如果用给定的 API 来计算，那么就快很多
	if unicode.IsLower(runes[index]) {
		runes[index] = unicode.ToUpper(runes[index])
	} else {
		runes[index] = unicode.ToLower(runes[index])
	}
	return string(runes)
}

func letterCasePermutation(S string) []string {
	result := []string{}
	runes := []rune(S)
	result = append(result, S)
	for i := 0; i < len(S); i++ {
		if unicode.IsLetter(runes[i]) {
			for _, value := range result {
				result = append(result, convert(value, i))
			}
		}
	}

	return result

}