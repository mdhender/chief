// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

// Direction is a direction out of the hex.
type Direction int

// enums for Direction
const (
	N Direction = iota
	NE
	SE
	S
	SW
	NW
)

// Add returns a new direction.
// When n is positive, it is clockwise.
// Negative is counter-clockwise.
func (d Direction) Add(n int) Direction {
	return Direction(modulo(int(d)+n, 6))
}

// Subtract returns a new direction.
// When n is positive, it is counter-clockwise.
// Negative is clockwise.
func (d Direction) Subtract(n int) Direction {
	return Direction(modulo(int(d)-n, 6))
}
