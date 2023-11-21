// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package turnrpt implements a parser for turn reports.
package turnrpt

import "encoding/json"

type Report struct {
	FileName           string      `json:"fileName,omitempty"`
	ClanId             string      `json:"clanId,omitempty"`
	Current            *Turn       `json:"current,omitempty"`
	Next               *Turn       `json:"next,omitempty"`
	Accounting         *Accounting `json:"accounting:omitempty"`
	DesiredCommodities []string    `json:"desiredCommodities,omitempty"`

	GMNotes string `json:"gmNotes,omitempty"`

	Tribes []*Tribe

	input []byte
	tribe *Tribe
	scout *Scout
}

func (r *Report) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

type Tribe struct {
	Id          string   `json:"id,omitempty"`
	Hex         string   `json:"hex,omitempty"`
	StartingHex string   `json:"startingHex,omitempty"`
	GoodsTribe  string   `json:"goodsTribe,omitempty"`
	Scouts      []*Scout `json:"scouts,omitempty"`
}

type Turn struct {
	Turn    string `json:"turn,omitempty"`
	Month   string `json:"month,omitempty"`
	Season  string `json:"season,omitempty"`
	Weather string `json:"weather,omitempty"`
	Due     string `json:"due,omitempty"`
}

type Accounting struct {
	Received string `json:"received,omitempty"`
	Cost     string `json:"cost,omitempty"`
	Credit   string `json:"credit,omitempty"`
}

type Scout struct {
	Id              string
	MovementResults []*MovementResult
	mr              *MovementResult
}

type MovementResult struct {
	Succeeded bool
	Direction string
	BlockedBy string
	Terrain   string
	Found     []string
	Notes     []string
}
