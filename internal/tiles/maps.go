// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import "sort"

type Map struct {
	tiles map[string]*Tile
	grid  struct {
		hash    string // replacement for "##" grid
		columns int    // number of columns in a grid
		rows    int    // number of rows in a grid
	}
}

func New(hash string) *Map {
	m := &Map{tiles: make(map[string]*Tile)}
	m.grid.hash = hash
	m.grid.columns, m.grid.rows = 30, 21
	return m
}

// MakeTile accepts TribeNet's ## XXYY coordinates.
// It returns a Tile using the internal coordinate system.
func (m *Map) MakeTile(gxy string) *Tile {
	x, y := gxyScale(gxy)
	t := &Tile{
		Hex: RowColToHex(y, x),
		id:  gxy,
	}
	return t
}

func (m *Map) Neighbor(from *Tile, d Direction) *Tile {
	hex := from.Neighbor(d)
	id := xlatToGXY(hex.ToXY())
	tt, ok := m.tiles[id]
	if !ok {
		tt = &Tile{Hex: hex, id: id}
		m.tiles[id] = tt
	}
	return tt
}

func (m *Map) Tiles() []*Tile {
	var t []*Tile
	for _, v := range m.tiles {
		t = append(t, v)
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i].id < t[j].id
	})
	return t
}
