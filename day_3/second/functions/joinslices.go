package functions

func CombineSlices(slices [][]int, slicer [][]int) [][]int {
	combined := make([][]int, 0, len(slices)+len(slicer))
	combined = append(combined, slices...)
	combined = append(combined, slicer...)
	return combined
}
