package functions

func AddSlices(str [][]int) [][]int {
	res := [][]int{}
	x, y := 0, 0

	for _, char := range str {
		x += char[0]
		y += char[1]
		res = append(res, []int{x, y})
	}

	return res
}
