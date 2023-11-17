// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import "fmt"

// Token represents a single token in the input stream.
// Name: mnemonic name (numeric).
// Val: string value of the token from the original stream.
// Pos: position - offset from beginning of stream.
type Token struct {
	Name TokenName
	Val  string
	Pos  int
}

type TokenName int

const (
	ERROR TokenName = iota
	EOF

	PLUS
	MINUS

	COMMA
	DOT

	LITERAL
	INTEGER
	
	TRIBE
)

var tokenNames = [...]string{
	ERROR: "ERROR",
	EOF:   "EOF",

	PLUS:  "PLUS",
	MINUS: "MINUS",

	COMMA: "COMMA",
	DOT:   "DOT",

	LITERAL: "LITERAL",
	INTEGER: "INTEGER",

	TRIBE: "TRIBE",
}

func (tok Token) String() string {
	return fmt.Sprintf("Token{%s, '%s', %d}", tokenNames[tok.Name], tok.Val, tok.Pos)
}
