// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

// Coordinates are x, y with
//
//	x increasing from left right
//	y increasing from top to bottom
type Coordinates struct {
	q int // x is col is q
	r int // y is row is r
}

func NewCoordinates(x, y int) Coordinates {
	return Coordinates{q: x, r: y}
}

func Neighbor(x, y, direction int) (int, int) {
	for direction < 0 {
		direction += 6
	}
	for direction > 6 {
		direction -= 6
	}
	var dy, dx int
	switch x & 1 {
	case 0: // col is even
		switch direction {
		case 0:
			dx, dy = 0, -1
		case 1:
			dx, dy = 1, -1
		case 2:
			dx, dy = 1, 0
		case 3:
			dx, dy = 0, 1
		case 4:
			dx, dy = -1, 0
		case 5:
			dx, dy = -1, -1
		}
	case 1: // col is odd
		switch direction {
		case 0:
			dx, dy = 0, -1
		case 1:
			dx, dy = 1, 0
		case 2:
			dx, dy = 1, 1
		case 3:
			dx, dy = 0, 1
		case 4:
			dx, dy = -1, 1
		case 5:
			dx, dy = -1, 0
		}
	}
	return x + dx, y + dy
}

// XYToHex assumes
func XYToHex(x, y int) Hex {
	return QOffsetToCube(x, y, ODD)
}

type HEXORIENTATION int

const (
	RADIUS = 30 // radius of a single hex on the board
)

// Axial coordinates are like Cube, but they don't store s.
// We can recover s since s = -q - r.
type Axial struct {
	Q int `json:"q"`
	R int `json:"r"`
}

// Cube coordinates have three axes, q, r, and s
// and the constraint that q + r + s = 0.
type Cube struct {
	Q int `json:"q"`
	R int `json:"r"`
	S int `json:"s"`
}

func (c Cube) ToAxial() Axial {
	var q = c.Q
	var r = c.R
	return Axial{Q: q, R: r}
}

// ToEvenQ shoves even rows to the right
func (c Cube) ToEvenQ() Offset {
	var col = c.Q
	var row = c.R + (c.Q+(c.Q&1))/2
	return Offset{Q: col, R: row}
}

// ToEvenR shoves even columns down
func (c Cube) ToEvenR() Offset {
	var col = c.Q + (c.R+(c.R&1))/2
	var row = c.R
	return Offset{Q: col, R: row}
}

// ToOddQ shoves odd rows to the right
func (c Cube) ToOddQ() Offset {
	var col = c.Q
	var row = c.R + (c.Q-(c.Q&1))/2
	return Offset{Q: col, R: row}
}

// ToOddR shoves odd columns down
func (c Cube) ToOddR() Offset {
	var col = c.Q + (c.R-(c.R&1))/2
	var row = c.R
	return Offset{Q: col, R: row}
}

// Offset coordinates use col and row.
// In an "odd" layout, odd rows are shoved to the right or down.
// In an "even" layout, even rows are shoved to the right or down.
// Which depends on if the layout is horizontal or vertical.
// The "horizontal" layout shoves rows to the right,
// while the "vertical" layout shoves them down.
type Offset struct {
	Q int `json:"q"` // colum
	R int `json:"r"` // row
}
