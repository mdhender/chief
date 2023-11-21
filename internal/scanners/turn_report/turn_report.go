// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package scanner

import (
	"fmt"
	"github.com/mdhender/chief/internal/docconv"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

// loosely inspired by https://tribenet.wiki/turns/report

// Scanner is a scanner that is specific to turn reports.
// This code assumes that we're scanning from the Word document directly.
type Scanner struct {
	filename             string
	state                State
	input                []byte // the body of the Word document
	pos, start           int
	meta                 map[string]string // metadata from the Word document
	ClanStatusHeading    []byte
	UnitHexStatusHeading []byte
}

// State maps to the report section
type State int

const (
	StStatus State = iota
	StActivity
	StScouting
	StUnitHex
	StPeople
	StGoods
	StSkills
	StWeight
	StEndOfInput
)

func NewTurnReportScanner(filename string) (*Scanner, error) {
	s := &Scanner{
		filename: filepath.Clean(filename),
		state:    StStatus,
	}
	log.Printf("parse: filename %s\n", s.filename)
	fp, err := os.Open(filename)
	if err != nil {
		log.Printf("[ptr] open %v\n", err)
		return nil, err
	}
	if body, meta, err := docconv.ConvertDocx(fp); err != nil {
		log.Printf("[ptr] convert %v\n", err)
		return nil, err
	} else {
		s.input, s.meta = []byte(body), meta
	}
	log.Printf("tr.scanner: input %d bytes\n", len(s.input))
	return s, nil
}

func (s *Scanner) IsEof() bool {
	for !s.iseof() && isspace(s.peekch()) {
		s.getch()
	}
	return s.iseof()
}

// Next returns the next token in the input.
func (s *Scanner) Next(state State) Token {
	for !s.iseof() && isspace(s.peekch()) {
		s.getch()
	}
	s.start = s.pos
	if s.iseof() {
		return Token{Type: EOF}
	}
	switch state {
	case StStatus:
		return s.nextStatus()
	case StActivity:
	case StScouting:
		return s.nextScout()
	case StUnitHex:
	case StPeople:
	case StGoods:
	case StSkills:
	case StWeight:
	}
	panic(fmt.Sprintf("assert(state != %d)", state))
}

func (s *Scanner) accept() []byte {
	b := s.input[s.start:s.pos]
	s.start = s.pos
	return b
}

func (s *Scanner) error() Token {
	for !s.iseof() && !isspace(s.peekch()) {
		s.getch()
	}
	return Token{Type: Error, Value: s.accept()}
}

func (s *Scanner) getch() byte {
	if s.iseof() {
		return 0
	}
	ch := s.input[s.pos]
	s.pos++
	return ch
}

func (s *Scanner) ignore() {
	s.start = s.pos
}

func (s *Scanner) iseof() bool {
	return !(s.pos < len(s.input))
}

func (s *Scanner) peekch() byte {
	if s.iseof() {
		return 0
	}
	return s.input[s.pos]
}

func (s *Scanner) reject() {
	s.pos = s.start
}

func (s *Scanner) ungetch() {
	if s.pos > 0 {
		s.pos--
	}
}

func (s *Scanner) unknown() Token {
	for !s.iseof() {
		ch := s.getch()
		if isspace(ch) || ch == ',' {
			s.ungetch()
			break
		}
	}
	return Token{Type: Unknown, Value: s.accept()}
}

func isalpha(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func isdigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

func isspace(ch byte) bool {
	return unicode.IsSpace(rune(ch))
}
