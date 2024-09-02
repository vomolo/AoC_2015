package functions

func DoesNotContain(str string) string {
	for i := 0; i < len(str)-1; i++ {
		if (str[i] == 'a' && str[i+1] == 'b') || (str[i] == 'c' && str[i+1] == 'd') || (str[i] == 'p' && str[i+1] == 'q') || (str[i] == 'x' && str[i+1] == 'y') {
			return ""
		}
	}
	return str
}
