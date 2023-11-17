// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

import "math"

type FractionalHex struct {
	q, r, s float64
}

// NewFractionalHex returns an initialized FractionalHex
func NewFractionalHex(q, r, s float64) FractionalHex {
	if math.Round(q+r+s) != 0 {
		panic("assert(q + r + s == 0)")
	}
	return FractionalHex{q: q, r: r, s: s}
}

// Lerp does a linear interpolation of
func (fh FractionalHex) Lerp(b FractionalHex, t float64) FractionalHex {
	return NewFractionalHex(fh.q*(1.0-t)+b.q*t, fh.r*(1.0-t)+b.r*t, fh.s*(1.0-t)+b.s*t)
}

// Round returns the hex that the fractional hex is located in.
func (fh FractionalHex) Round() Hex {
	qi := int(math.Round(fh.q))
	q_diff := math.Abs(float64(qi) - fh.q)

	ri := int(math.Round(fh.r))
	r_diff := math.Abs(float64(ri) - fh.r)

	si := int(math.Round(fh.s))
	s_diff := math.Abs(float64(si) - fh.s)

	if q_diff > r_diff && q_diff > s_diff {
		qi = -ri - si
	} else if r_diff > s_diff {
		ri = -qi - si
	} else {
		si = -qi - ri
	}

	return NewHex(qi, ri, si)
}
