// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

type DoubledCoord struct {
	col, row int
}

func NewDoubledCoord(col, row int) DoubledCoord {
	return DoubledCoord{col: col, row: row}
}

func (a DoubledCoord) Equals(b DoubledCoord) bool {
	return a.col == b.col && a.row == b.row
}

func qdoubled_from_cube(h Hex) DoubledCoord {
	col := h.q
	row := 2*h.r + h.q

	return NewDoubledCoord(col, row)
}

func qdoubled_to_cube(h DoubledCoord) Hex {
	q := h.col
	r := (h.row - h.col) / 2
	s := -q - r

	return NewHex(q, r, s)
}

func rdoubled_from_cube(h Hex) DoubledCoord {
	col := 2*h.q + h.r
	row := h.r

	return NewDoubledCoord(col, row)
}

func rdoubled_to_cube(h DoubledCoord) Hex {
	q := (h.col - h.row) / 2
	r := h.row
	s := -q - r

	return NewHex(q, r, s)
}
