package functions

func RemoveDuplicates(coords [][]int) [][]int {
	seen := make(map[[2]int]bool)
	uniqueCoords := [][]int{}

	for _, coord := range coords {
		key := [2]int{coord[0], coord[1]}
		if !seen[key] {
			uniqueCoords = append(uniqueCoords, coord)
			seen[key] = true
		}
	}

	return uniqueCoords
}
