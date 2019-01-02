/**
In C programming, a character variable holds ASCII value (an integer number between
0 and 127) rather than that character itself.

The ASCII value of lowercase alphabets are from 97 to 122. And, the ASCII value of
 uppercase alphabets are from 65 to 90.

**/
func toLowerCase(str string) string {
	if str == "" {
		return str
	}

	b := []byte(str)
	for i := 0; i < len(b); i++ {
		if b[i] >= 65 && b[i] <= 90 {
			b[i] = b[i] + 32
		}

	}

	return string(b)

	// Or you can just use oneline : return strings.ToLower(str)
}