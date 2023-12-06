// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import "math"

// Layout is a collection of hexes.
// We are going to use flat top hex in an "even-q" vertical layout.
// This shoves even columns down.
type Layout struct {
	// Size is measured from the center of a tile to a vertex
	// (equivalent to radius of smallest circle that encloses the tile).
	Size float64

	// Width for flat top hexes is 2 * Size
	Width float64

	// Height for flat top hexes is sqrt(3) * Size
	Height float64

	// horizontal distance between two hexes is size * 3 / 2
	HorizontalOffset float64

	// vertical distance between two hexes is simply size * sqrt(3)
	VerticalOffset float64

	Origin      point
	Orientation orientation
}

func NewLayout(radius float64) Layout {
	return Layout{
		Size:             radius,
		Width:            2 * radius,
		Height:           math.Sqrt(3) * radius,
		HorizontalOffset: radius * 3 / 2,
		VerticalOffset:   math.Sqrt(3) / 2,
		Orientation:      flatOrientation(3.0/2.0, 0.0, math.Sqrt(3.0)/2.0, math.Sqrt(3.0), 2.0/3.0, 0.0, -1.0/3.0, math.Sqrt(3.0)/3.0),
	}
}

// orientation stores the forward matrix (the fN variables) and backward matrix
// (the bN variables), plus the start angle. The start angle determines if we
// have a "flat top" (which is 0°) or "pointy top" (which is 60°) hex.
type orientation struct {
	f0, f1, f2, f3 float64
	b0, b1, b2, b3 float64
	// The starting angle should be 0.0 for 0° (flat top) or 0.5 for 60° (pointy top).
	startAngle float64 // in multiples of 60°
}

// flatOrientation returns an initialized Orientation.
func flatOrientation(f0, f1, f2, f3, b0, b1, b2, b3 float64) orientation {
	return orientation{
		f0: f0, f1: f1, f2: f2, f3: f3,
		b0: b0, b1: b1, b2: b2, b3: b3,
		startAngle: 0.0,
	}
}

var (
	qBasisVector = struct {
		x, y float64
	}{
		x: 3 / 2,
		y: math.Sqrt(3) / 2,
	}
	rBasisVector = struct {
		x, y float64
	}{
		x: 0,
		y: math.Sqrt(3),
	}
)

// centerPoint returns the center point of the hex on the screen.
func (l Layout) centerPoint(hex Hex) point {
	M := l.Orientation

	x := (M.f0*float64(hex.q) + M.f1*float64(hex.r)) * l.Width
	y := (M.f2*float64(hex.q) + M.f3*float64(hex.r)) * l.Height

	return point{l.Origin.x + x, l.Origin.y + y}
}

// hexCornerOffset returns the offset of a hex corner from the center of the hex.
// The offset accounts for the size of the hex and the orientation.
// Corner ranges from 0..5.
func (l Layout) hexCornerOffset(corner int) point {
	M := l.Orientation
	angle := 2.0 * math.Pi * (M.startAngle - float64(corner)) / 6.0
	return point{x: l.Width * math.Cos(angle), y: l.Height * math.Sin(angle)}
}

func (l Layout) hexToPoint(hex Hex) point {
	return point{
		x: l.Size * (float64(hex.q) * qBasisVector.x /* + float64(hex.r) * rBasisVector.x */),
		y: l.Size * (float64(hex.q)*qBasisVector.y + float64(hex.r)*rBasisVector.y),
	}
}

// polygonCorners returns the screen coordinates for all the corners of the hex.
// It uses the layout to determine the orientation of the hex and the center point
// of it on the screen.
func (l Layout) polygonCorners(h Hex) (corners []point) {
	center := l.centerPoint(h)
	for i := 0; i < 6; i++ {
		offset := l.hexCornerOffset(i)
		corners = append(corners, point{x: center.x + offset.x, y: center.y + offset.y})
	}

	return corners
}
