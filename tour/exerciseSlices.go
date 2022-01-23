package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, 0)

	for i := 0; i < dy; i += 1 {
		row := make([]uint8, 0)
		for j := 0; j < dx; j += 1 {
			row = append(row, uint8(i*j%123))
		}
		p = append(p, row)
	}

	return p
}

func main() {
	pic.Show(Pic)
}
