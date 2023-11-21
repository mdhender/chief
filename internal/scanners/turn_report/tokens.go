// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package scanner

import "fmt"

type Token struct {
	Type  Type
	Value []byte
}

func (t Token) String() string {
	return fmt.Sprintf("{%d, %q}", t.Type, string(t.Value))
}

func (t Token) Is(typ Type) bool {
	return t.Type == typ
}

func (t Token) IsNot(typ Type) bool {
	return t.Type != typ
}

type Type int

func (t Type) String() string {
	return fmt.Sprintf("(%d)", t)
}

const (
	Unknown Type = iota
	EOF
	EndOfSection
	Literal
	BulletNumber
	CantMoveOnOcean
	Commodity
	Cost
	Credit
	Currency
	CurrentHex
	CurrentTurn
	Date
	DesiredCommodities
	Direction
	GoodsTribe
	HexNo
	Into
	NextTurn
	NoDesiredCommodities
	NoFordOnRiver
	NotEnoughMPs
	NothingOfInterest
	OfHex
	PreviousHex
	Received
	Season
	Scout
	ScoutNo
	Terrain
	Tribe
	TribeNo
	TurnMonth
	TurnNo
	Weather
	Delimiter
	BackSlash
	Colon
	Comma
	Dash
	EqualSign
	LeftParen
	RightParen
	Error
)
