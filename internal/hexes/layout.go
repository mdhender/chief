// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

import "math"

// Layout is used to convert between hex and screen coordinates.
// See https://www.redblobgames.com/grids/hexagons/#basics for details
// on what size is used for.
type Layout struct {
	orientation  Orientation
	size, origin Point
}

// NewLayout returns an initialized Layout.
func NewLayout(orientation Orientation, size, origin Point) Layout {
	return Layout{orientation: orientation, size: size, origin: origin}
}

// NewFlatLayout returns an initialized layout using flat hexes.
func NewFlatLayout(size, origin Point) Layout {
	return Layout{orientation: layout_flat, size: size, origin: origin}
}

// NewPointyLayout returns an initialized layout using pointy hexes
func NewPointyLayout(size, origin Point) Layout {
	return Layout{orientation: layout_pointy, size: size, origin: origin}
}

// CenterPoint returns the center point of the hex on the screen.
func (l Layout) CenterPoint(h Hex) Point {
	M := l.orientation
	size := l.size
	origin := l.origin

	x := (M.f0*float64(h.q) + M.f1*float64(h.r)) * size.x
	y := (M.f2*float64(h.q) + M.f3*float64(h.r)) * size.y

	return NewPoint(x+origin.x, y+origin.y)
}

// CoordToHex converts an x, y coordinate to a fractional hex on the map.
func (l Layout) CoordToHex(x, y int) Hex {
	p := NewPoint(float64(x), float64(y))
	M := l.orientation
	size := NewPoint(1.0, 1.0)
	origin := l.origin

	pt := NewPoint((p.x-origin.x)/size.x, (p.y-origin.y)/size.y)

	q := M.b0*pt.x + M.b1*pt.y
	r := M.b2*pt.x + M.b3*pt.y

	return NewFractionalHex(q, r, -q-r).Round()
}

// HexCornerOffset returns the offset of a hex corner from the center of the hex.
// The offset accounts for the size of the hex and the orientation.
// Corner ranges from 0..5.
func (l Layout) HexCornerOffset(corner int) Point {
	M := l.orientation
	size := l.size

	angle := 2.0 * math.Pi * (M.start_angle - float64(corner)) / 6.0

	return NewPoint(size.x*math.Cos(angle), size.y*math.Sin(angle))
}

// PixelToHex converts a point on the screen to a fractional hex on the map.
func (l Layout) PixelToHex(p Point) FractionalHex {
	M := l.orientation
	size := l.size
	origin := l.origin

	pt := NewPoint((p.x-origin.x)/size.x, (p.y-origin.y)/size.y)

	q := M.b0*pt.x + M.b1*pt.y
	r := M.b2*pt.x + M.b3*pt.y

	return NewFractionalHex(q, r, -q-r)
}

// PolygonCorners returns the screen coordinates for all the corners of the hex.
// It uses the layout to determine the orientation of the hex and the center point
// of it on the screen.
func (l Layout) PolygonCorners(h Hex) (corners []Point) {
	center := l.CenterPoint(h)
	for i := 0; i < 6; i++ {
		offset := l.HexCornerOffset(i)
		corners = append(corners, NewPoint(center.x+offset.x, center.y+offset.y))
	}

	return corners
}

var layout_flat = NewOrientation(3.0/2.0, 0.0, math.Sqrt(3.0)/2.0, math.Sqrt(3.0), 2.0/3.0, 0.0, -1.0/3.0, math.Sqrt(3.0)/3.0, FLATHEX)

var layout_pointy = NewOrientation(math.Sqrt(3.0), math.Sqrt(3.0)/2.0, 0.0, 3.0/2.0, math.Sqrt(3.0)/3.0, -1.0/3.0, 0.0, 2.0/3.0, POINTYHEX)
