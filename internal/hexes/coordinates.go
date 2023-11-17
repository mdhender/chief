// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

// Doubled coordinates use col and row.
// They have the constraint that (col + row) % 2 = 0.
type Doubled struct {
	Q int `json:"q"` // colum
	R int `json:"r"` // row
}

func axial_to_cube(a Axial) Cube {
	var q = a.Q
	var r = a.R
	var s = -q - r
	return Cube{Q: q, R: r, S: s}
}

func axial_to_oddr(a Axial) Offset {
	var col = a.Q + (a.R-(a.R&1))/2
	var row = a.R
	return Offset{Q: col, R: row}
}

func oddr_to_axial(o Offset) Axial {
	var q = o.Q - (o.R-(o.R&1))/2
	var r = o.R
	return Axial{Q: q, R: r}
}

func axial_to_evenr(a Axial) Offset {
	var col = a.Q + (a.R+(a.R&1))/2
	var row = a.R
	return Offset{Q: col, R: row}
}

func evenr_to_axial(o Offset) Axial {
	var q = o.Q - (o.R+(o.R&1))/2
	var r = o.R
	return Axial{Q: q, R: r}
}

func axial_to_oddq(a Axial) Offset {
	var col = a.Q
	var row = a.R + (a.Q-(a.Q&1))/2
	return Offset{Q: col, R: row}
}

func oddq_to_axial(o Offset) Axial {
	var q = o.Q
	var r = o.R - (o.Q-(o.Q&1))/2
	return Axial{Q: q, R: r}
}

func axial_to_evenq(a Axial) Offset {
	var col = a.Q
	var row = a.R + (a.Q+(a.Q&1))/2
	return Offset{Q: col, R: row}
}

func evenq_to_axial(o Offset) Axial {
	var q = o.Q
	var r = o.R - (o.Q+(o.Q&1))/2
	return Axial{Q: q, R: r}
}

func oddr_to_cube(o Offset) Cube {
	var q = o.Q - (o.R-(o.R&1))/2
	var r = o.R
	return Cube{Q: q, R: r, S: -q - r}
}

func evenr_to_cube(o Offset) Cube {
	var q = o.Q - (o.R+(o.R&1))/2
	var r = o.R
	return Cube{Q: q, R: r, S: -q - r}
}

func OddQToCube(col, row int) Cube {
	var q = col
	var r = row
	return oddq_to_cube(Offset{Q: q, R: r})
}

func oddq_to_cube(o Offset) Cube {
	var q = o.Q
	var r = o.R - (o.Q-(o.Q&1))/2
	return Cube{Q: q, R: r, S: -q - r}
}

func evenq_to_cube(o Offset) Cube {
	var q = o.Q
	var r = o.R - (o.Q+(o.Q&1))/2
	return Cube{Q: q, R: r, S: -q - r}
}

func doubleheight_to_axial(d Doubled) Axial {
	var q = d.Q
	var r = (d.R - d.Q) / 2
	return Axial{Q: q, R: r}
}

func axial_to_doubleheight(a Axial) Doubled {
	var col = a.Q
	var row = 2*a.R + a.Q
	return Doubled{Q: col, R: row}
}

func doublewidth_to_axial(d Doubled) Axial {
	var q = (d.Q - d.R) / 2
	var r = d.R
	return Axial{Q: q, R: r}
}

func axial_to_doublewidth(a Axial) Doubled {
	var col = 2*a.Q + a.R
	var row = a.R
	return Doubled{Q: col, R: row}
}

func doubleheight_to_cube(d Doubled) Cube {
	var q = d.Q
	var r = (d.R - d.Q) / 2
	return Cube{Q: q, R: r, S: -q - r}
}

func cube_to_doubleheight(c Cube) Doubled {
	var col = c.Q
	var row = 2*c.R + c.Q
	return Doubled{Q: col, R: row}
}

func doublewidth_to_cube(d Doubled) Cube {
	var q = (d.Q - d.R) / 2
	var r = d.R
	return Cube{Q: q, R: r, S: -q - r}
}

func cube_to_doublewidth(c Cube) Doubled {
	var col = 2*c.Q + c.R
	var row = c.R
	return Doubled{Q: col, R: row}
}
