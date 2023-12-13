// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

//// Layout is used to convert between hex and screen coordinates.
//// See https://www.redblobgames.com/grids/hexagons/#basics for details
//// on what size is used for.
//type Layout struct {
//	orientation  Orientation
//	size, origin Point
//}
//
//// CoordToHex converts an x, y coordinate to a fractional hex on the map.
//func (l Layout) CoordToHex(x, y int) Hex {
//	p := NewPoint(float64(x), float64(y))
//	M := l.orientation
//	size := NewPoint(1.0, 1.0)
//	origin := l.origin
//
//	pt := NewPoint((p.x-origin.x)/size.x, (p.y-origin.y)/size.y)
//
//	q := M.b0*pt.x + M.b1*pt.y
//	r := M.b2*pt.x + M.b3*pt.y
//
//	return NewFractionalHex(q, r, -q-r).Round()
//}
//
//// PixelToHex converts a point on the screen to a fractional hex on the map.
//func (l Layout) PixelToHex(p Point) FractionalHex {
//	M := l.orientation
//	size := l.size
//	origin := l.origin
//
//	pt := NewPoint((p.x-origin.x)/size.x, (p.y-origin.y)/size.y)
//
//	q := M.b0*pt.x + M.b1*pt.y
//	r := M.b2*pt.x + M.b3*pt.y
//
//	return NewFractionalHex(q, r, -q-r)
//}
