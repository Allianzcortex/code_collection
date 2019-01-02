import "strings"

/**
Given a string, you need to reverse the order of characters in each word
within a sentence while still preserving whitespace and initial word order.

Example 1:
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Note: In the string, each word is separated by single space and there
 will not be any extra space in the string.

**/
func reverseWord(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, value := range s {
		n--
		runes[n] = value
	}
	return string(runes[n:])

}

func reverseWords(s string) string {
	if s == "" {
		return s
	}
	buffer := strings.Split(s, " ")
	// 下面是一开始自己的写法，效率很低，只比 17% 的人快
	result := ""
	for _, value := range buffer {
		result += reverseWord(value)
		result += " "
	}
	return strings.Trim(result, " ")

	// 但更改方法后用 strings.Join() 来连接效率就高很多，比 87% 的人快
	result := []string{}
	for _, value := range buffer {
		result = append(result, reverseWord(value))
	}
	return strings.Join(result, " ")

}