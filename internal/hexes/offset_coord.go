// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

type OFFSET int

const (
	EVEN OFFSET = 1
	ODD  OFFSET = -1
)

type OffsetCoord struct {
	col, row int
}

func NewOffsetCoord(col, row int) OffsetCoord {
	return OffsetCoord{col: col, row: row}
}

func (oc OffsetCoord) Equals(b OffsetCoord) bool {
	return oc.col == b.col && oc.row == b.row
}

func (oc OffsetCoord) Coords() (x, y int) {
	return oc.col, oc.row
}

func QOffsetToCube(col, row int, offset OFFSET) Hex {
	return qoffset_to_cube(offset, OffsetCoord{col: col, row: row})
}

func qoffset_from_cube(offset OFFSET, h Hex) OffsetCoord {
	col := h.q
	row := h.r + (h.q+int(offset)*(h.q&1))/2

	return NewOffsetCoord(col, row)
}

func qoffset_to_cube(offset OFFSET, h OffsetCoord) Hex {
	q := h.col
	r := h.row - (h.col+int(offset)*(h.col&1))/2
	s := -q - r

	return NewHex(q, r, s)
}

func roffset_from_cube(offset OFFSET, h Hex) OffsetCoord {
	col := h.q + (h.r+int(offset)*(h.r&1))/2
	row := h.r

	return NewOffsetCoord(col, row)
}

func roffset_to_cube(offset OFFSET, h OffsetCoord) Hex {
	q := h.col - ((h.row + int(offset)*(h.row&1)) / 2)
	r := h.row
	s := -q - r

	return NewHex(q, r, s)
}
