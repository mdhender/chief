// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import "fmt"

// Hex uses cube coordinates.
// We will need a function to translate to offset coordinates.
type Hex struct {
	q, r, s int
}

// Add returns the hex that is the given vector away from the source.
// Here, a vector is a direction times a scalar.
func (hex Hex) Add(vec Hex) Hex {
	return Hex{q: hex.q + vec.q, r: hex.r + vec.r, s: hex.s + vec.s}
}

// CheckConstraint returns true only if the hex satisfies the constraint q + r + s == 0.
func (hex Hex) CheckConstraint() bool {
	return hex.q+hex.r+hex.s == 0
}

// Distance is the number of hexes that must be entered
// when moving from source to destination hex.
func (hex Hex) Distance(b Hex) int {
	vec := hex.Subtract(b)
	return max(abs(vec.q), abs(vec.r), abs(vec.s))
}

// Neighbor returns the hex that is next to the source in the given direction.
func (hex Hex) Neighbor(d Direction) Hex {
	return hex.Add(AsVector(d))
}

// Subtract returns vector between two hexes.
// The result, when added to the original hex, should return the destination hex.
func (hex Hex) Subtract(b Hex) Hex {
	return Hex{q: hex.q - b.q, r: hex.r - b.r, s: hex.s - b.s}
}

// ToRowCol returns the coordinates of the hex in row, column format.
func (hex Hex) ToRowCol() (row, col int) {
	row, col = hex.r+(hex.q+(hex.q&1))/2, hex.q
	return row, col
}

// ToXY returns the coordinates of the hex in x, y format.
func (hex Hex) ToXY() (x, y int) {
	y, x = hex.ToRowCol()
	return x, y
}

// Wedge is the general direction between two hexes.
func (hex Hex) Wedge(b Hex) Direction {
	vec := hex.Subtract(b)
	wedge := max(abs(vec.q-vec.r), abs(vec.r-vec.s), abs(vec.s-vec.q))
	return Direction(modulo(wedge, 6))
}

var (
	// cc is short for cube coordinates, our internal system
	ccDirectionVectors = []Hex{
		N:  {q: 0, r: -1, s: +1},
		NE: {q: +1, r: -1, s: 0},
		SE: {q: +1, r: 0, s: -1},
		S:  {q: 0, r: +1, s: -1},
		SW: {q: -1, r: +1, s: 0},
		NW: {q: -1, r: 0, s: +1},
	}
)

// AsVector transforms a direction into a vector.
func AsVector(d Direction) Hex {
	return ccDirectionVectors[d]
}

// RowColToHex converts the coordinates to a Hex with the same coordinates.
func RowColToHex(row, col int) Hex {
	q, r := col, row-(col+(col&1))/2
	return Hex{q: q, r: r, s: -q - r}
}

func (hex Hex) XY() string {
	x, y := hex.ToXY()
	return fmt.Sprintf("(%d, %d)", x, y)
}

func (hex Hex) String() string {
	return fmt.Sprintf("{%d, %d, %d}", hex.q, hex.r, hex.s)
}
