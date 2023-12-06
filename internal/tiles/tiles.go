// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import (
	"github.com/mdhender/chief/internal/edge"
	"github.com/mdhender/chief/internal/terrain"
)

type Tile struct {
	Hex
	Terrain terrain.Terrain
	// N, NE, SE, S, SW, NW
	Edges [6]edge.Edge
	id    string
}

// Id is the unique identifier for the Tile.
func (t *Tile) Id() string {
	return t.id
}
