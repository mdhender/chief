// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package svg

import (
	"bytes"
	"fmt"
	"github.com/mdhender/chief/internal/hexes"
	"github.com/mdhender/chief/internal/terrain"
	"math"
)

const (
	EDGE   = 20
	EDGES  = EDGE * 2
	RADIUS = 10
)

type SVG struct {
	id       string
	width    float64
	height   float64
	viewBox  viewBox
	layout   hexes.Layout
	polygons []*polygon
}

type Hex struct {
	cube    hexes.Cube
	terrain terrain.Terrain
}

func New(cols, rows int, ac bool) *SVG {
	addCoordinates = ac

	//offset := (math.Sqrt(3) * RADIUS) / 2
	//// cols -> width  -> x
	//maxX := EDGES + offset*float64(cols*2)
	//// rows -> height -> y
	//maxY := EDGES + offset*float64(rows)*math.Sqrt(3)

	s := &SVG{
		id:     "s",
		width:  2.0 * RADIUS,
		height: math.Sqrt(3.0) * RADIUS,
		viewBox: viewBox{
			minX: 0,
			minY: 0,
			//width:  int(maxX),
			//height: int(maxY),
		},
		//polygons: make(map[int]*polygon),
	}

	// assumes flat with even-q layout
	s.layout = hexes.NewFlatLayout(hexes.NewPoint(RADIUS, RADIUS), hexes.NewPoint(s.height, s.width))

	return s
}

func (s *SVG) AddHex(x, y int, t terrain.Terrain) {
	poly := &polygon{
		x:       x,
		y:       y,
		radius:  s.height / 2.0,
		terrain: t,
	}
	h := hexes.XYToHex(x, y)
	poly.cx, poly.cy = s.layout.CenterPoint(h).Coords()
	//if int(poly.cx) > s.viewBox.width {
	//	s.viewBox.width = int(poly.cx + poly.radius)
	//}
	//if int(poly.cy) > s.viewBox.height {
	//	s.viewBox.height = int(poly.cy + poly.radius)
	//}

	poly.style.stroke = "Grey"
	poly.style.fill = t.ToFill()
	if poly.style.fill == poly.style.stroke {
		poly.style.stroke = "Black"
		if poly.style.fill == poly.style.stroke {
			poly.style.stroke = "White"
		}
	}
	poly.style.strokeWidth = "2px"
	poly.style.strokeWidth = "1px"

	for _, p := range s.layout.PolygonCorners(h) {
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
		terrain.Clear,
		terrain.Delta,
		terrain.Desert,
		terrain.Forest,
		terrain.Gravel,
		terrain.Ice,
		terrain.Mountain,
		terrain.Ocean,
		terrain.Plain,
		terrain.Rock,
		terrain.Rough,
		terrain.SacredMountain,
		terrain.SaltMarsh,
		terrain.Sea,
		terrain.Steppe,
		terrain.Swamp,
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
				buf.Write(poly.Bytes(id))
			} else {
				buf.Write(poly.Use(ref, id))
			}
		}
	}

	buf.Write([]byte("</svg>"))

	return buf.Bytes()
}

type point struct {
	x, y float64
}

func (p point) Bytes() []byte {
	return []byte(fmt.Sprintf("%f,%f", p.x, p.y))
}

type viewBox struct {
	minX, minY    int
	width, height int
}

func (v viewBox) Bytes() []byte {
	return []byte(fmt.Sprintf(` viewBox="%d %d %d %d"`, v.minX, v.minY, v.width+EDGE/2, v.height+EDGE/2))
}
