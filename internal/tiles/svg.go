// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import (
	"bytes"
	"fmt"
	"github.com/mdhender/chief/internal/terrain"
	"math"
)

const (
	EDGE   = 20
	EDGES  = EDGE * 2
	RADIUS = 10
)

type SVG struct {
	id             string
	width          float64
	height         float64
	viewBox        viewBox
	layout         Layout
	polygons       []*polygon
	addCoordinates bool
}

type SHex struct {
	cube    Hex
	terrain terrain.Terrain
}

func NewSVG(ac bool) *SVG {
	s := &SVG{
		id:     "s",
		width:  2.0 * RADIUS,
		height: math.Sqrt(3.0) * RADIUS,
		viewBox: viewBox{
			minX: 0,
			minY: 0,
		},
		addCoordinates: ac,
	}

	s.layout = NewLayout(RADIUS)

	return s
}

func (s *SVG) AddTile(tile *Tile) {
	x, y := tile.ToXY()
	poly := &polygon{
		x:       x,
		y:       y,
		radius:  s.height / 2.0,
		terrain: tile.Terrain,
	}
	h := tile.Hex
	poly.cx, poly.cy = s.layout.centerPoint(h).Coords()

	poly.style.stroke = "Grey"
	poly.style.fill = tile.Terrain.ToFill()
	if poly.style.fill == poly.style.stroke {
		poly.style.stroke = "Black"
		if poly.style.fill == poly.style.stroke {
			poly.style.stroke = "White"
		}
	}
	poly.style.strokeWidth = "2px"
	poly.style.strokeWidth = "1px"

	for _, p := range s.layout.polygonCorners(h) {
		px, py := p.Coords()
		poly.points = append(poly.points, point{x: px, y: py})
		if int(px) > s.viewBox.width {
			s.viewBox.width = int(px)
		}
		if int(py) > s.viewBox.height {
			s.viewBox.height = int(py)
		}
	}

	s.polygons = append(s.polygons, poly)
}

func (s *SVG) Bytes() []byte {
	buf := bytes.Buffer{}

	buf.WriteString("<svg")
	if s.id != "" {
		buf.WriteString(fmt.Sprintf(" id=%q", s.id))
	}
	//buf.WriteString(fmt.Sprintf(` width="%dpx" height="%dpx"`, s.viewBox.width, s.viewBox.height))
	buf.Write(s.viewBox.Bytes())
	buf.Write([]byte(` xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`))
	buf.WriteByte('\n')

	buf.WriteString("<style>@import url(medoly.css);</style>\n")

	for i, t := range []terrain.Terrain{
		terrain.Unknown,
		terrain.ALPS,
		terrain.AR,
		terrain.BH,
		terrain.BR,
		terrain.CH,
		terrain.DE,
		terrain.DF,
		terrain.DH,
		terrain.FORDS,
		terrain.GH,
		terrain.HSM,
		terrain.JG,
		terrain.JH,
		terrain.L,
		terrain.LCM,
		terrain.LJM,
		terrain.LSM,
		terrain.O,
		terrain.PI,
		terrain.PR,
		terrain.R,
		terrain.RH,
		terrain.SH,
		terrain.SW,
		terrain.TU,
	} {
		if i > 0 {
			buf.WriteByte('\n')
		}
		var ref *polygon
		id := t.String()
		for _, poly := range s.polygons {
			if poly.terrain != t {
				continue
			}
			if ref == nil {
				ref = poly
				buf.Write(poly.Bytes(id, s.addCoordinates))
			} else {
				buf.Write(poly.Use(ref, id, s.addCoordinates))
			}
		}
	}

	buf.Write([]byte("</svg>"))

	return buf.Bytes()
}

type viewBox struct {
	minX, minY    int
	width, height int
}

func (v viewBox) Bytes() []byte {
	return []byte(fmt.Sprintf(` viewBox="%d %d %d %d"`, v.minX, v.minY, v.width+EDGE/2, v.height+EDGE/2))
}
