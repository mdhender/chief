// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"bytes"
	"fmt"
)

// Scouting is complex.
//
//	action :== result (BackSlash result)*
//
//	result :== unitId
//	  | OceanBlocked Direction OfHex (Comma finding)*
//	  | RiverBlocked Direction OfHex (Comma finding)*
//	  | MPExhausted Direction Terrain (Comma finding)*
//	  | Direction Dash Terrain (Comma Terrain Direction (Comma Direction)*)* (Comma finding)*
//
//	finding :== NothingOfInterest
//	  | Find ore
//	  | Patrolled unitId+

type lexer struct{}
type item struct{}

// Scout 1:Scout SE-RH, River SE S SW\NE-PR, River S\SE-PR, River SE S SW\ 1190\NE-PR, River S\ not enough M.P's to move to SE into PRAIRIE, nothing of interest found
// Scout 2:Scout,  can't Move on Ocean to SW of HEX, Patrolled and found 1138
// Scout 3:Scout SE-GH\ S-PR, River SE S\ No Ford on River to S of HEX, Nothing of interest found
// Scout 6:Scout N-PR,  O SW,  NW\N-PR,  O NE,  SW,  NW,  N\ Can't Move on Ocean to NE of HEX,  Nothing of interest found
// Scout 7:Scout N-GH\ NE-PR\ NW-PR\ SE-PR\  Not enough M.P's to move to N into ROCKY HILLS, Nothing of interest found
// Scout 8:Scout SW-PR,  O N\NW-CH,  O NE\NW-RH,  O NW,  N, Find Iron Ore\ Can't Move on Ocean to NW of HEX,  Nothing of interest found
func newScoutingLexer(input []byte) *lexer {
	var items []item
	_, input = lexSpaces(input)
	for n := 1; n <= 8; n++ {
		pfx := []byte(fmt.Sprintf("SCOUT %d:", n))
		if bytes.HasPrefix(input, pfx) {
			items = append(items, item{})
			input = input[len(pfx):]
			continue
		}
	}
	for _, input = lexSpaces(input); len(input) != 0; _, input = lexSpaces(input) {
		var token, rest []byte
		if input[0] == ',' {
			items = append(items, item{})
			input = input[1:]
		} else if input[0] == '\\' {
			items = append(items, item{})
			input = input[1:]
		} else if token, rest = lexDeposit(input); token != nil {
			items = append(items, item{})
			input = rest
		} else if token, rest = lexDirection(input); token != nil {
			items = append(items, item{})
			input = rest
		} else if token, rest = lexTerrain(input); token != nil {
			items = append(items, item{})
			input = rest
		} else if token, rest = lexUnit(input); token != nil {
			items = append(items, item{})
			input = rest
		} else if token = []byte("CAN'T MOVE ON OCEAN TO "); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
		} else if token = []byte("FIND "); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
		} else if token = []byte("INTO "); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
			continue
		} else if token = []byte("NO FORD ON RIVER TO "); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
			continue
		} else if token = []byte("NOT ENOUGH M.P'S TO MOVE TO "); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
			continue
		} else if token = []byte("NOTHING OF INTEREST FOUND"); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
			continue
		} else if token = []byte("OF HEX "); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
			continue
		} else if token = []byte("PATROLLED AND FOUND "); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
			continue
		} else if token = []byte("SCOUT,"); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
			continue
		} else if token = []byte("SCOUT "); bytes.HasPrefix(input, token) {
			items = append(items, item{})
			input = input[len(token)-1:]
			continue
		} else if token, rest = nextToken(input); token != nil {
			items = append(items, item{})
			input = rest
		}
	}

	return &lexer{}
}
