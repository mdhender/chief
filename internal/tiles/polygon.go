// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import (
	"bytes"
	"fmt"
	"github.com/mdhender/chief/internal/terrain"
)

// polygon is the actual hex on the board
type polygon struct {
	x, y    int
	terrain terrain.Terrain // terrain type of the hex

	cx, cy, radius float64 // center of the hex
	points         []point

	style struct {
		fill        string
		stroke      string
		strokeWidth string
	}

	addCircle bool
	text      []string
}

func (p *polygon) Bytes(id string, addCoords bool) []byte {
	buf := bytes.Buffer{}
	buf.WriteString(`<polygon`)
	if id != "" {
		buf.WriteString(fmt.Sprintf(` id="%s"`, id))
	}
	buf.WriteString(fmt.Sprintf(` fill="%s"`, p.style.fill))
	buf.WriteString(fmt.Sprintf(` stroke="%s"`, p.style.stroke))
	buf.WriteString(fmt.Sprintf(` stroke-width="%s"`, p.style.strokeWidth))
	buf.WriteString(fmt.Sprintf(` points="`))

	for i, pt := range p.points {
		if i > 0 {
			buf.WriteByte(' ')
		}
		buf.Write(pt.Bytes())
	}
	buf.WriteString(`"></polygon>`)
	buf.WriteByte('\n')

	if addCoords {
		fontSize := 8
		buf.WriteString(fmt.Sprintf(`<text x="%f" y="%f" text-anchor="middle" fill="grey" font-size="%d" font-weight="bold">%s</text>`, p.cx, p.cy, fontSize, fmt.Sprintf("%02d %02d", p.x, p.y)))
		buf.WriteByte('\n')
	}

	return buf.Bytes()
}

func (p *polygon) Use(ref *polygon, id string, addCoords bool) []byte {
	buf := bytes.Buffer{}
	dx := p.cx - ref.cx
	dy := p.cy - ref.cy
	buf.WriteString(fmt.Sprintf(`<use href="#%s" x="%f" y="%f" />`, id, dx, dy))
	buf.WriteByte('\n')

	if addCoords {
		fontSize := 8
		buf.WriteString(fmt.Sprintf(`<text x="%f" y="%f" text-anchor="middle" fill="grey" font-size="%d" font-weight="bold">%s</text>`, p.cx, p.cy, fontSize, fmt.Sprintf("%02d %02d", p.x, p.y)))
		buf.WriteByte('\n')
	}

	return buf.Bytes()
}
