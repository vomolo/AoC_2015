package functions

func ContainsVowels(str string) string {
	count := 0

	for _, char := range str {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
		if count >= 3 {
			return str
		}
	}
	return ""
}
