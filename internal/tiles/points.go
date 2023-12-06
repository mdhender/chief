// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import "fmt"

// point can be a pixel?
type point struct {
	x, y float64
}

func (p point) Bytes() []byte {
	return []byte(fmt.Sprintf("%f,%f", p.x, p.y))
}

func (p point) Coords() (x, y float64) {
	return p.x, p.y
}
