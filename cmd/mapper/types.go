// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"bytes"
	"fmt"
	"github.com/mdhender/chief/internal/terrain"
)

type Consolidated struct {
	// Turns is a map of turn number (e.g. 800-12) to Turn data.
	Turns map[string]*Turn
}

// Turn contains all the movement and scouting results for the turn.
type Turn struct {
	Id       string `json:"id,omitempty"`
	FileName string `json:"filename,omitempty"`
	// Units is a map of unit id (e.g. 1138c1) to Unit data
	Units map[string]*Unit `json:"units,omitempty"`
}

// Unit is data for the unit moving or scouting.
type Unit struct {
	Id string `json:"id,omitempty"`
	// Kind is tribe, clan, element, or courier
	Kind string `json:"kind,omitempty"`
	// Notes are GM or player notes for the unit
	Notes []*Note `json:"notes,omitempty"`
	// Location is where the unit started and finished the turn
	Location *UnitLocation `json:"location,omitempty"`
	// Movement is a slice of the unit's movements.
	Movement []*Movement
	// Scouts is a slice of the tribe's scouts.
	Scouts map[string]*Scout
	// Check is a sanity check from translating the report to json.
	Check *Check
}

// UnitLocation tracks the starting and ending hexes for a unit.
type UnitLocation struct {
	// Current is the Hex the unit ends the turn in
	Current string `json:"current,omitempty"`
	// StartedIn is the Hex the unit starts the turn in
	StartedIn string `json:"previous,omitempty"`
}

// Note is free-text for a hex.
type Note struct {
	Hex  string `json:"hex,omitempty"`
	Text string `json:"text,omitempty"`
}

// Movement is a unit's movement, not its scouts.
type Movement struct {
	Direction string          `json:"direction,omitempty"`
	Result    *MovementResult `json:"result,omitempty"`
}

// MovementResult is the result of an attempted move.
// It may fail or succeed.
type MovementResult struct {
	// From is the hex the movement started in.
	From string `json:"from,omitempty"`
	// Failed is set only if the movement failed.
	Failed *MovementFailure `json:"failed,omitempty"`
	// To is the hex the movement ended in.
	// If the movement failed, then From and To will be the same value.
	To string `json:"to,omitempty"`
	// Terrain is set only if the move succeeded.
	// It is the terrain of the hex entered.
	Terrain terrain.Terrain `json:"terrain,omitempty"`
	// Edges defines the terrain or feature that can be seen in a given direction.
	// It is always relative to the hex the movement ended in.
	Edges map[string]*HexEdge `json:"edges,omitempty"`
}

// MovementFailure represents why the movement attempt failed.
type MovementFailure struct {
	NoFord      bool `json:"no-ford,omitempty"`
	NotEnoughMp bool `json:"not-enough-mp"`
	OceanCoast  bool `json:"ocean-coast,omitempty"`
}

type Hexagon struct {
	Grid    string `json:"grid,omitempty"`
	Row     int    `json:"row,omitempty"`
	Column  int    `json:"column,omitempty"`
	Id      string `json:"id,omitempty"`
	Terrain string `json:"terrain,omitempty"`
	Edges   [6]struct {
		Kind  string   `json:"kind,omitempty"`
		To    *Hexagon `json:"-"`
		Notes []string `json:"notes,omitempty"`
	}
	Notes []string `json:"notes,omitempty"`
}

// HexEdge describes the crossing from one hex to another.
// MovementPoints are related to the type of the edge.
type HexEdge struct {
	Direction    string `json:"direction,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ConiferHills bool   `json:"conifer-hills,omitempty"`
	Ford         bool   `json:"ford,omitempty"`
	GrassyHills  bool   `json:"grassy-hills,omitempty"`
	Ocean        bool   `json:"ocean,omitempty"`
	Prairie      bool   `json:"prairie,omitempty"`
	River        bool   `json:"river,omitempty"`
	RockyHills   bool   `json:"rocky-hills,omitempty"`
	Swamp        bool   `json:"swamp,omitempty"`
}

func (e HexEdge) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", e.Kind)), nil
}

func (e *HexEdge) UnmarshalJSON(buf []byte) error {
	if len(buf) == 0 || bytes.Equal(buf, []byte{'n', 'u', 'l', 'l'}) {
		return nil
	} else if buf[0] != '"' || buf[len(buf)-1] != '"' {
		return fmt.Errorf("invalid edge %q", string(buf))
	}
	buf = buf[1 : len(buf)-1]
	switch string(bytes.ToUpper(buf)) {
	case "CH", "CONIFER HILLS":
		e.ConiferHills = true
		e.Kind = "CH"
	case "FORD":
		e.Ford = true
		e.Kind = "FO"
	case "GH", "GRASSY HILLS":
		e.GrassyHills = true
		e.Kind = "GH"
	case "O", "OCEAN":
		e.Ocean = true
		e.Kind = "O"
	case "PR", "PRAIRIE":
		e.Prairie = true
		e.Kind = "PR"
	case "RIVER":
		e.River = true
		e.Kind = "R"
	case "RH", "ROCKY HILLS":
		e.RockyHills = true
		e.Kind = "RH"
	case "SWAMP":
		e.Swamp = true
		e.Kind = "SW"
	default:
		return fmt.Errorf("unknown edge %q", string(buf))
	}
	return nil
}

// Scout is a single scouting party.
type Scout struct {
	Id string `json:"-"`
	// Scout is the "scout" action
	Scout []*ScoutAction `json:"scout,omitempty"`
}

// ScoutAction is the "scout" action
type ScoutAction struct {
	Direction string          `json:"direction"`
	Result    *MovementResult `json:"result"`
}

// Check is a sanity check.
type Check struct {
	Hex     string              `json:"hex,omitempty"`
	Terrain string              `json:"terrain,omitempty"`
	Edges   map[string]*HexEdge `json:"edges,omitempty"`
}
