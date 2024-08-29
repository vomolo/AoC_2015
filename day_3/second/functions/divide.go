package functions

func Divide(str string) (string, string) {
	var odd, even string
	for i := range str {
		if i%2 == 0 {
			even += string(str[i])
		} else {
			odd += string(str[i])
		}
	}
	return even, odd
}
