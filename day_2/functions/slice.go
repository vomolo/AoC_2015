package functions

func Slice(str string) []string {
	res := []string{}
	num := ""

	for _, char := range str {
		if char == 'x' {
			if num != "" {
				res = append(res, num)
				num = ""
			}
		} else if char != ' ' {
			num += string(char)
		}
	}
	if num != "" {
		res = append(res, num)
	}
	return res
}
