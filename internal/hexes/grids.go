// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package hexes

import (
	"math"
)

// Grid is used to generate hexes.
type Grid struct {
	o      Orientation
	kind   GridType
	origin Point // used when creating images
	size   Point // width and height of a hexagon
	// offset (in screen pixels) to each corner of a hexagon
	hexCornerOffsets [6]Point
}

type GridType int

const (
	EVEN_Q GridType = iota // flat hexes, push even columns down
	ODD_Q                  // flat hexes, push odd columns down
	EVEN_R                 // pointy hexes, push even rows down
	ODD_R                  // pointy hexes, push odd rows down
)

// NewGrid returns an initialized factory for generating hexagons.
// It is also used to convert between hex and screen coordinates.
//
// If flatHexes is true, the grid is vertical with flat-top hexagons.
// Otherwise, it is horizontal with pointy-top hexagons.
//
// When pushEven is true, hexagons in even rows/columns are pushed
// down when converting from Cube to Offset coordinates. (Offset
// coordinates are also known as row+column, or x,y coordinates.)
//
// Please see https://www.redblobgames.com/grids/hexagons/#coordinates-offset
// for a detailed description.
func NewGrid(flatHexes bool, pushEven bool, size, origin Point) Grid {
	g := Grid{}
	if flatHexes {
		g.o = Orientation{
			3.0 / 2.0, 0.0, math.Sqrt(3.0) / 2.0, math.Sqrt(3.0),
			2.0 / 3.0, 0.0, -1.0 / 3.0, math.Sqrt(3.0) / 3.0,
			0.5,
		}
		if pushEven {
			g.kind = EVEN_Q
		} else {
			g.kind = ODD_Q
		}
	} else {
		g.o = Orientation{
			math.Sqrt(3.0), math.Sqrt(3.0) / 2.0, 0.0, 3.0 / 2.0,
			math.Sqrt(3.0) / 3.0, -1.0 / 3.0, 0.0, 2.0 / 3.0,
			0.0,
		}
		if pushEven {
			g.kind = EVEN_R
		} else {
			g.kind = ODD_R
		}
	}

	// calculate offset (in screen pixels) to every corner of the hex
	for corner := 0; corner < 6; corner++ {
		angle := 2 * math.Pi * (g.o.start_angle - float64(corner)) / 6
		g.hexCornerOffsets[corner] = NewPoint(g.size.x*math.Cos(angle), g.size.y*math.Sin(angle))
	}

	return g
}

// CenterPoint returns the screen coordinates of the center point of the hex.
func (g Grid) CenterPoint(hex Hex) Point {
	cx := (g.o.f0*float64(hex.q) + g.o.f1*float64(hex.r)) * g.size.x
	cy := (g.o.f2*float64(hex.q) + g.o.f3*float64(hex.r)) * g.size.y
	return NewPoint(cx+g.origin.x, cy+g.origin.y)
}

func (g Grid) Orientation() Orientation {
	return g.o
}

// PolygonCorners returns the screen coordinates for all the corners of the hex.
// It uses the layout to determine the orientation of the hex and the center point
// of it on the screen.
func (g Grid) PolygonCorners(hex Hex) (corners []Point) {
	center := g.CenterPoint(hex)
	for corner := 0; corner < 6; corner++ {
		corners = append(corners, NewPoint(center.x+g.hexCornerOffsets[corner].x, center.y+g.hexCornerOffsets[corner].y))
	}
	return corners
}
