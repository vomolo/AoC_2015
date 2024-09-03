package functions

func NonOverlappingPair(s string) string {
	pairPositions := make(map[string][]int)

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		pairPositions[pair] = append(pairPositions[pair], i)

	}

	for _, positions := range pairPositions {
		if len(positions) > 1 {
			for j := 0; j < len(positions)-1; j++ {
				for k := j + 1; k < len(positions); k++ {
					if positions[k] >= positions[j]+2 {
						return s
					}
				}
			}
		}
	}

	return ""
}
