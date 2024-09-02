package functions

func DoubleLetters(str string) string {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return str
		}
	}
	return ""
}
