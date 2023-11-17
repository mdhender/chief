// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package board

import (
	"github.com/mdhender/chief/internal/hexes"
	"github.com/mdhender/chief/internal/svg"
	"github.com/mdhender/chief/internal/terrain"
	"log"
)

const (
	RADIUS = 30 // radius of a single hex on the board
)

type Board struct {
	Cols int // x is col
	Rows int // y is row

	// hexes are indexed by y, x
	hexes [][]*hex
}

type hex struct {
	coords  hexes.Coordinates
	terrain terrain.Terrain
}

func New(cols, rows int) *Board {
	//log.Printf("board: cols %d rows %d\n", cols, rows)
	b := &Board{
		Cols:  cols,
		Rows:  rows,
		hexes: make([][]*hex, rows, rows),
	}
	for y := 0; y < rows; y++ {
		b.hexes[y] = make([]*hex, cols, cols)
	}
	return b
}

func (b *Board) AsSVG(addCoordinates bool) []byte {
	s := svg.New(b.Cols, b.Rows, addCoordinates)
	for y := 0; y < len(b.hexes); y++ {
		for x := 0; x < len(b.hexes[y]); x++ {
			if b.IsSet(x, y) {
				s.AddHex(x, y, b.hexes[y][x].terrain)
			}
		}
	}
	return s.Bytes()
}

// Bounds returns the minimum and maximum value for rows and columns on the board
func (b *Board) Bounds() (minCol, minRow, maxCol, maxRow int) {
	return 0, 0, b.Cols, b.Rows
}

func (b *Board) IsSet(x, y int) bool {
	return b.hexes[y][x] != nil
}

func (b *Board) GetTerrain(x, y int) terrain.Terrain {
	return b.hexes[y][x].terrain
}

func (b *Board) SetTerrain(x, y int, t terrain.Terrain) {
	if !(0 <= x && x < b.Cols) {
		log.Printf("bad: %s: x: 0 <= %d < %d\n", t.String(), x, b.Cols)
		return
	} else if !(0 <= y && y < b.Rows) {
		log.Printf("bad: %s: y: 0 <= %d < %d\n", t.String(), y, b.Rows)
		return
	}
	if b.hexes[y][x] == nil {
		b.hexes[y][x] = &hex{coords: hexes.NewCoordinates(x, y)}
	}
	b.hexes[y][x].terrain = t
}
