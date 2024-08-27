package functions

func Sum(num []int) int {

	var sum int

	for _, char := range num {
		sum += char
	}

	return sum

}
