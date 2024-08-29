package functions

func MakeSlice(str string) [][]int {
	res := [][]int{{0, 0}}

	for _, char := range str {
		if char == '>' {
			res = append(res, []int{1, 0})
		}

		if char == '<' {
			res = append(res, []int{-1, 0})
		}

		if char == '^' {
			res = append(res, []int{0, 1})
		}

		if char == 'v' {
			res = append(res, []int{0, -1})
		}
	}

	return res
}
