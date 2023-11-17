// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

type Move struct {
	From      *Hex          `json:"from,omitempty"`
	Direction MoveDirection `json:"direction"`
	To        *Hex          `json:"to,omitempty"`
	Result    *MoveResult   `json:"result,omitempty"`
}

type MoveResult struct {
	Failed  bool `json:"failed,omitempty"`
	Blocked struct {
		NoFord    bool `json:"no_ford,omitempty"`
		Coastline bool `json:"coastline,omitempty"`
	}
	Terrain string `json:"terrain,omitempty"`
	See     struct {
		N  string `json:"N,omitempty"`
		NE string `json:"NE,omitempty"`
		SE string `json:"SE"`
		S  string `json:"S,omitempty"`
		SW string `json:"SW,omitempty"`
		NW string `json:"NW,omitempty"`
	} `json:"see,omitempty"`
	Text  string `json:"text,omitempty"`
	Found string `json:"found,omitempty"`
}
