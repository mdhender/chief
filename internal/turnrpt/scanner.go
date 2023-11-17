// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"bytes"
	"fmt"
	"unicode"
)

func scan(input []byte) *scanner {
	return &scanner{input: input}
}

type scanner struct {
	input []byte
	start int
	pos   int
}

const eof = 0

type token struct {
	typ string
	val string
	err error
}

func (t token) String() string {
	if t.err != nil {
		return fmt.Sprintf("{%s %q %v}", t.typ, t.val, t.err)
	}
	return fmt.Sprintf("{%s %q}", t.typ, t.val)
}

// accept accepts the current lexeme and updates the start point
func (s *scanner) accept() (lexeme []byte) {
	lexeme, s.start = s.input[s.start:s.pos], s.pos
	return lexeme
}

// backup moves one byte backwards in the input.
// it will not back up beyond the start of the input.
func (s *scanner) backup() {
	if s.pos <= 0 {
		return
	}
	s.pos--
}

// emit returns a token and updates `start` to prepare for the next token.
func (s *scanner) emit(typ string) token {
	t := token{typ: typ, val: string(s.input[s.start:s.pos])}
	s.start = s.pos
	return t
}

// errorf returns an error token.
// like emit, it updates `start` to prepare for the next token.
func (s *scanner) errorf(format string, args ...any) token {
	t := token{typ: "error", val: string(s.input[s.start:s.pos]), err: fmt.Errorf(format, args...)}
	s.start = s.pos
	return t
}

// get returns the next byte or eof.
func (s *scanner) get() byte {
	if s.isEof() {
		return eof
	}
	ch := s.input[s.pos]
	s.pos++
	return ch
}

// ignore updates `start` to prepare for the next token.
func (s *scanner) ignore() {
	s.start = s.pos
}

// isEof returns true if the input is exhausted.
func (s *scanner) isEof() bool {
	return len(s.input) <= s.pos
}

// peek returns the next byte (or eof) without advancing the input.
func (s *scanner) peek() byte {
	if s.isEof() {
		return eof
	}
	return s.input[s.pos]
}

// next returns the next token in the input after skipping any leading whitespace.
// note: this is the "default state" for scanning.
func (s *scanner) next() token {
	// skip any leading whitespace
	for isSpace(s.peek()) {
		s.get()
	}
	s.ignore()
	ch := s.get()
	switch ch {
	case eof:
		return s.emit("eof")
	case ',':
		return s.emit("comma")
	case '.':
		return s.emit("dot")
	}
	if ch == 'T' {
		if ch = s.get(); ch == 'r' {
			if ch = s.get(); ch == 'i' {
				if ch = s.get(); ch == 'b' {
					if ch = s.get(); ch == 'e' {
						if isSpace(s.peek()) {
							return s.emit("tribe")
						}
					}
				}
			}
		}
		for !isDelim(s.peek()) {
			s.get()
		}
		return s.emit("literal")
	}
	if isDigit(ch) {
		if ch = s.get(); isDigit(ch) {
			if ch = s.get(); isDigit(ch) {
				if ch = s.get(); isDigit(ch) {
					if s.peek() == eof || s.peek() == ',' {
						return s.emit("integer")
					}
				}
			}
		}
	}

	// anything else must be a literal
	for !isDelim(s.peek()) {
		s.get()
	}
	return s.emit("literal")
}

func (s *scanner) restore(save int) {
	s.start, s.pos = save, save
}

func (s *scanner) skipws(accept bool) {
	for s.pos < len(s.input) && unicode.IsSpace(rune(s.input[0])) {
		s.pos++
	}
	if accept {
		s.start = s.pos
	}
}

func (s *scanner) nextLiteral() (lexeme []byte) {
	defer func(save int) {
		if lexeme == nil {
			s.restore(save)
		}
	}(s.start)
	s.skipws(true)
	panic("!")
}

func (s *scanner) nextHexNo() (lexeme []byte) {
	defer func(save int) {
		if lexeme == nil {
			s.restore(save)
		}
	}(s.start)
	s.skipws(true)

	if x, y := s.get(), s.get(); !((x == '#' && y == '#') || (isDigit(x) && isDigit(y))) {
		return nil
	}
	if sp := s.get(); sp != ' ' {
		return nil
	}
	for i := 0; i < 4; i++ {
		if n := s.get(); !isDigit(n) {
			return nil
		}
	}
	delim := []byte{' ', '\f', '\n', '\r', '\t', ',', ')', eof}
	if bytes.IndexByte(delim, s.peek()) == -1 {
		return nil
	}
	return s.accept()
}

type Hex struct {
	GridX, GridY string
	Col, Row     int
}

func (h *Hex) String() string {
	if h == nil {
		return "nil"
	}
	return fmt.Sprintf("{X: %s, Y: %s, Col: %02d, Row: %02d}", h.GridX, h.GridY, h.Col, h.Row)
}

// CurrentHex returns a Hex if it finds "Current Hex = ## 0102"
func (s *scanner) CurrentHex() (hex *Hex, err error) {
	panic("!")
	//defer func(saved []byte) {
	//	if hex == nil {
	//		s.input = saved
	//	}
	//}(s.input)
	//
	//s.skipws()
	//if !bytes.HasPrefix(s.input, []byte("Current Hex = ")) {
	//	return nil, nil
	//}
	//s.input = s.input[14:]
	//h := s.HexNo()
	//if h == nil {
	//	return nil, fmt.Errorf("current hex: missing hex no")
	//}
	//return h, nil
}
