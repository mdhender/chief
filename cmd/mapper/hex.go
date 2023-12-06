// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
)

type Hex struct {
	X, Y int
}

func NewHex(x, y int) Hex {
	return Hex{X: x, Y: y}
}

// Move uses the EvenQ logic - even columns are shoved 1/2 of a row
func (h Hex) Move(md MoveDirection) Hex {
	var dx, dy int
	even := (h.X & 1) == 0
	if even {
		switch md {
		case N:
			dx, dy = 0, -1
		case NE:
			dx, dy = 1, 0
		case SE:
			dx, dy = 1, 1
		case S:
			dx, dy = 0, +1
		case SW:
			dx, dy = -1, 1
		case NW:
			dx, dy = -1, 0
		}
	} else {
		switch md {
		case N:
			dx, dy = 0, -1
		case NE:
			dx, dy = 1, -1
		case SE:
			dx, dy = 1, 0
		case S:
			dx, dy = 0, +1
		case SW:
			dx, dy = -1, 0
		case NW:
			dx, dy = -1, -1
		}
	}
	return Hex{X: h.X + dx, Y: h.Y + dy}
}

// String implements the string.Stringer interface.
func (h Hex) String() string {
	return fmt.Sprintf("%02d%02d", h.X, h.Y)
}
