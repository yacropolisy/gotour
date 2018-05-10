package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := [][]uint8{}

	for x := 0; x < dx; x++ {
		res = append(res, []uint8{})

		for y := 0; y < dy; y++ {

			res[x] = append(res[x], uint8(x*y))
		}
	}

	return res
}

func main() {
	pic.Show(Pic)
}
