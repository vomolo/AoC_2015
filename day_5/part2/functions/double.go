package functions

func RepeatedLetterWithOneInBetween(s string) string {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && isLetter(s[i]) {
			return s
		}
	}
	return ""
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
