// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

// Orientation stores the forward matrix (the fN variables) and backward matrix
// (the bN variables), plus the start angle. The start angle determines if we
// have a "flat top" (which is 0°) or "pointy top" (which is 60°) hex.
type Orientation struct {
	f0, f1, f2, f3 float64 // todo: document these
	b0, b1, b2, b3 float64 // todo: document these
	// The starting angle should be 0.0 for 0° (flat top) or 0.5 for 60° (pointy top).
	start_angle float64 // in multiples of 60°
}
