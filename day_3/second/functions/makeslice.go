package functions

func MakeSlice(str string) [][]int {
	res := [][]int{{0, 0}}

	directionMap := map[rune][]int{
		'>': {1, 0},
		'<': {-1, 0},
		'^': {0, 1},
		'v': {0, -1},
	}

	for _, char := range str {
		if direction, ok := directionMap[char]; ok {
			res = append(res, direction)
		}
	}
	return res
}
