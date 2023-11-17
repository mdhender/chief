// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

type ScoutingParty struct {
	Id       string   `json:"id"`
	Action   string   `json:"action,omitempty"`
	Start    *Hex     `json:"start,omitempty"`
	Finish   *Hex     `json:"finish,omitempty"`
	RawText  string   `json:"raw_text,omitempty"`
	RawMoves []string `json:"raw_moves,omitempty"`
	Moves    []*Move  `json:"moves,omitempty"`
}
