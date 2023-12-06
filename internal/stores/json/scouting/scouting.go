// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package scouting implements a JSON store for scouting results.
package scouting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mdhender/chief/internal/edge"
	"github.com/mdhender/chief/internal/terrain"
	"os"
)

func ReadFile(name string) (*Results, error) {
	r := Results{FileName: name}

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("scouting: read: %w", err)
	}

	// return an error if the input has unexpected fields.
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&r); err != nil {
		return nil, err
	}

	// link id fields
	for id, unit := range r.Units {
		unit.Id = id
		for id, scout := range unit.Scouts {
			scout.Id = id
		}
	}

	return &r, nil
}

// Results is the data in the scouting results.
type Results struct {
	FileName string `json:"file-name,omitempty"`
	Turn     string `json:"turn,omitempty"`
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
	// Follows is set only if the unit is following another.
	Follows string `json:"follows,omitempty"`
	// Movement is a slice of the unit's movements.
	Movement []*Movement `json:"movement,omitempty"`
	// Scouts is a slice of the tribe's scouts.
	Scouts map[string]*Scout `json:"scouts,omitempty"`
	// Check is a sanity check from translating the report to json.
	Check *Check `json:"check,omitempty"`
}

// Note is free-text for a hex.
type Note struct {
	Hex  string `json:"hex,omitempty"`
	Text string `json:"text,omitempty"`
}

// UnitLocation tracks the starting and ending hexes for a unit.
type UnitLocation struct {
	// Current is the Hex the unit ends the turn in
	Current string `json:"current,omitempty"`
	// StartedIn is the Hex the unit starts the turn in
	StartedIn string `json:"previous,omitempty"`
}

// Movement is a unit's movement, or a scout's "scout" action.
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
	Edges map[string]*edge.Edge `json:"edges,omitempty"`
	// Found is a list of things the scouting party found.
	// If the movement succeeded, these will be in the To hex.
	// Otherwise, they are in the From hex.
	Found []string `json:"found,omitempty"`
}

// MovementFailure represents why the movement attempt failed.
type MovementFailure struct {
	NoFord      bool `json:"no-ford,omitempty"`
	NotEnoughMp bool `json:"not-enough-mp"`
	OceanCoast  bool `json:"ocean-coast,omitempty"`
}

// Scout is a single scouting party.
type Scout struct {
	Id string `json:"-"`
	// Scout is the "scout" action
	Scout []*Movement `json:"scout,omitempty"`
}

// ScoutAction is the "scout" action
type ScoutAction struct {
	Direction string          `json:"direction"`
	Result    *MovementResult `json:"result"`
}

// Check is a sanity check.
type Check struct {
	Hex     string                `json:"hex,omitempty"`
	Terrain terrain.Terrain       `json:"terrain,omitempty"`
	Edges   map[string]*edge.Edge `json:"edges,omitempty"`
	// Found is a list of things found in the hex.
	Found []string `json:"found,omitempty"`
}
