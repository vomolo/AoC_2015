package functions

func SurfaceArea(str []int) int {

	W := str[0]
	H := str[1]
	L := str[2]

	res := (2*((W*H)+(W*L)+(H*L)) + (W * H))

	return res
}
