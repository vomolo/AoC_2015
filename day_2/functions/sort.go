package functions

func SortInts(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	copy(result, nums)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				// Swap result[j] and result[j+1]
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}
