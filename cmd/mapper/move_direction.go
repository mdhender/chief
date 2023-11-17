// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"github.com/mdhender/chief/internal/hexes"
)

type MoveDirection int

const (
	N MoveDirection = iota
	NE
	SE
	S
	SW
	NW
)

// MarshalJSON implements the json.Marshaler interface.
func (md MoveDirection) MarshalJSON() ([]byte, error) {
	switch md {
	case N:
		return []byte{'N'}, nil
	case NE:
		return []byte{'N', 'E'}, nil
	case SE:
		return []byte{'S', 'E'}, nil
	case S:
		return []byte{'S'}, nil
	case SW:
		return []byte{'S', 'W'}, nil
	case NW:
		return []byte{'N', 'W'}, nil
	}
	panic(fmt.Sprintf("assert(md != %d)", md))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (md *MoveDirection) UnmarshalJSON(b []byte) error {
	switch len(b) {
	case 1:
		switch b[0] {
		case 'N':
			*md = N
			return nil
		case 'S':
			*md = S
			return nil
		}
	case 2:
		switch b[0] {
		case 'N':
			switch b[1] {
			case 'E':
				*md = NE
				return nil
			case 'W':
				*md = NW
				return nil
			}
		case 'S':
			switch b[1] {
			case 'E':
				*md = SE
				return nil
			case 'W':
				*md = SW
				return nil
			}
		}
	}
	return fmt.Errorf("unknown direction")
}

// String implements the string.Stringer interface.
func (md MoveDirection) String() string {
	switch md {
	case N:
		return "N"
	case NE:
		return "NE"
	case SE:
		return "SE"
	case S:
		return "S"
	case SW:
		return "SW"
	case NW:
		return "NW"
	}
	panic(fmt.Sprintf("assert(md != %d)", md))
}

// ToHexDirection is a helper for calculating moves.
func (md MoveDirection) ToHexDirection() hexes.Hex {
	var q, r, s int
	switch md {
	case N:
		q, r, s = 0, -1, +1
	case NE:
		q, r, s = +1, -1, 0
	case SE:
		q, r, s = +1, 0, -1
	case S:
		q, r, s = 0, +1, -1
	case SW:
		q, r, s = -1, +1, 0
	case NW:
		q, r, s = -1, 0, +1
	}
	return hexes.NewHex(q, r, s)
}
