package functions

// func main() {
//     // Example slice of strings
//     strSlice := []string{"42", "123", "0", "-7", "abc", "3.14"}

//     // Convert slice of strings to slice of ints
//     intSlice, err := ConvertToInt(strSlice)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println("Slice of integers:", intSlice)
// }

func ConvertToInt(strSlice []string) []int {
	intSlice := make([]int, len(strSlice))

	for i, str := range strSlice {
		num := parseInt(str)

		intSlice[i] = num
	}

	return intSlice
}

func parseInt(s string) int {
	var result int
	var sign int = 1
	startIndex := 0

	// Check for negative sign
	if s[0] == '-' {
		sign = -1
		startIndex = 1
	} else if s[0] == '+' {
		startIndex = 1
	}

	// Convert characters to integer
	for i := startIndex; i < len(s); i++ {
		char := s[i]

		result = result*10 + int(char-'0')
	}

	return result * sign
}
