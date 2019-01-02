func uniqueMorseRepresentations(words []string) int {
	if len(words) == 0 {
		return 0
	}
	table := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	result := make(map[string]bool)

	for _, word := range words {
		str := ""
		for i := 0; i < len(word); i++ {
			str += table[word[i]-97]
		}
		result[str] = true
	}

	return len(result)
}