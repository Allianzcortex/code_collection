import (
	"fmt"
	"strconv"
)

/**
You are given an array A of strings.

Two strings S and T are special-equivalent if after any number of moves, S == T.

A move consists of choosing two indices i and j with i % 2 == j % 2, and swapping S[i] with S[j].

Now, a group of special-equivalent strings from A is a non-empty subset S of A such that any string not in S is not special-equivalent with any string in S.

Return the number of groups of special-equivalent strings from A.



Example 1:

Input: ["a","b","c","a","c","c"]
Output: 3
Explanation: 3 groups ["a","a"], ["b"], ["c","c","c"]
Example 2:

Input: ["aa","bb","ab","ba"]
Output: 4
Explanation: 4 groups ["aa"], ["bb"], ["ab"], ["ba"]
Example 3:

Input: ["abc","acb","bac","bca","cab","cba"]
Output: 3
Explanation: 3 groups ["abc","cba"], ["acb","bca"], ["bac","cab"]
Example 4:

Input: ["abcd","cdab","adcb","cbad"]
Output: 1
Explanation: 1 group ["abcd","cdab","adcb","cbad"]


Note:

1 <= A.length <= 1000
1 <= A[i].length <= 20
All A[i] have the same length.
All A[i] consist of only lowercase letters.


**/
func int2String(arr [26]int) string {
	s := ""
	for i := 0; i < 26; i++ {
		s += strconv.Itoa(arr[i])
	}

	return s
}

func numSpecialEquivGroups(A []string) int {
	if len(A) == 0 {
		return 0
	}

	result := make(map[string]bool)
	for _, value := range A {
		odd := [26]int{}
		even := [26]int{}
		for i := 0; i < len(value); i++ {
			if i%2 == 0 {
				//fmt.Println(value[i]-'a')
				odd[value[i]-'a']++
			} else {
				even[value[i]-'a']++
			}
		}
		fmt.Println(int2String(odd) + int2String(even))
		result[int2String(odd)+int2String(even)] = true

	}

	return len(result)

}