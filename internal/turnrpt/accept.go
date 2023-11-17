// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

func acceptChar(ch byte, input []byte) (run, rest []byte) {
	if len(input) == 0 || input[0] != ch {
		return nil, input
	}
	return input[:1], input[1:]
}

// acceptDate accepts a date in the format DD/MM/YYYY.
func acceptDate(input []byte) (run, rest []byte) {
	if len(input) < 10 {
		return nil, input
	} else if !(isDigit(input[0]) && isDigit(input[1])) {
		return nil, input
	} else if input[2] != '/' {
		return nil, input
	} else if !(isDigit(input[3]) && isDigit(input[4])) {
		return nil, input
	} else if input[5] != '/' {
		return nil, input
	} else if !(isDigit(input[6]) && isDigit(input[7]) && isDigit(input[8]) && isDigit(input[9])) {
		return nil, input
	}
	run, rest = input[:10], input[10:]
	delim := []byte{' ', '\f', '\n', '\r', '\t'}
	if len(rest) != 0 && bytes.IndexByte(delim, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

// acceptHexNo accepts a hex number in the form "NN XXYY"
// where "NN" ie either "##" or an integer, and "XXYY" is an integer.
func acceptHexNo(input []byte) (run, rest []byte) {
	if len(input) < 7 {
		return nil, input
	} else if !((input[0] == '#' && input[1] == '#') || (isDigit(input[0]) && isDigit(input[1]))) {
		return nil, input
	} else if input[2] != ' ' {
		return nil, input
	} else if !(isDigit(input[3]) && isDigit(input[4]) && isDigit(input[5]) && isDigit(input[6])) {
		return nil, input
	}
	run, rest = input[:7], input[7:]
	delim := []byte{' ', '\f', '\n', '\r', '\t', ',', ')'}
	if len(rest) != 0 && bytes.IndexByte(delim, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

// acceptInteger accepts an integer.
func acceptInteger(input []byte) (run, rest []byte) {
	if len(input) < 4 || !isdigit(input[0]) {
		return nil, input
	}
	pos := 0
	for pos < len(input) && isDigit(input[pos]) {
		pos++
	}
	delim := []byte{' ', '\f', '\n', '\r', '\t', ',', '(', ')'}
	if pos < len(input) && bytes.IndexByte(delim, input[pos]) == -1 {
		return nil, input
	}
	return input[:pos], input[pos:]
}

func acceptLiteral(literal string, input []byte) (run, rest []byte) {
	if !bytes.HasPrefix(input, []byte(literal)) {
		return nil, input
	}
	run, rest = input[:len(literal)], input[len(literal):]
	if ws, _ := acceptWS(input); len(ws) == 0 {
		return nil, input
	}
	return run, rest
}

// acceptSeason accepts Spring or Winter.
func acceptSeason(input []byte) (run, rest []byte) {
	for _, literal := range []string{"Spring", "Winter"} {
		if run, rest = acceptLiteral(literal, input); run != nil {
			return run, rest
		}
	}
	return nil, input
}

// acceptTribeNo accepts a four digit integer in the range 0000...9999.
func acceptTribeNo(input []byte) (run, rest []byte) {
	if len(input) < 4 {
		return nil, input
	} else if !(isDigit(input[0]) && isDigit(input[1]) && isDigit(input[2]) && isDigit(input[3])) {
		return nil, input
	}
	run, rest = input[:4], input[4:]
	delim := []byte{' ', '\f', '\n', '\r', '\t', ','}
	if len(rest) != 0 && bytes.IndexByte(delim, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

// acceptTurnNo accepts a turn number (YYY-MM).
func acceptTurnNo(input []byte) (run, rest []byte) {
	if len(input) < 6 {
		return nil, input
	} else if !(isDigit(input[0]) && isDigit(input[1]) || isDigit(input[2])) {
		return nil, input
	} else if input[3] != '-' {
		return nil, input
	} else if !(isDigit(input[4]) && isDigit(input[5])) {
		return nil, input
	}
	run, rest = input[:6], input[6:]
	delim := []byte{' ', '\f', '\n', '\r', '\t', ','}
	if len(rest) != 0 && bytes.IndexByte(delim, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

// acceptWeather accepts FINE.
func acceptWeather(input []byte) (run, rest []byte) {
	for _, literal := range []string{"FINE"} {
		if run, rest = acceptLiteral(literal, input); run != nil {
			return run, rest
		}
	}
	return nil, input
}

func acceptWS(input []byte) (run, rest []byte) {
	if len(input) == 0 {
		return nil, nil
	} else if !isSpace(input[0]) {
		return nil, input
	}
	pos := 0
	for pos < len(input) && isSpace(input[pos]) {
		pos++
	}
	return input[:pos], input[pos:]
}

func alnum(input []byte) (run, rest []byte) {
	if len(input) == 0 {
		return nil, nil
	} else if !isAlNum(input[0]) {
		return nil, input
	}
	pos := 0
	for pos < len(input) && isAlNum(input[pos]) {
		pos++
	}
	return input[:pos], input[pos:]
}

func isAlNum(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}

func isAlpha(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isDelim(ch byte) bool {
	if unicode.IsSpace(rune(ch)) {
		return true
	}
	ws := []byte{' ', '\f', '\n', '\r', '\t'}
	delim := []byte{'.', ',', ':', ';', '(', ')', '#', '$', '/', '\\'}
	return ch == eof || bytes.IndexByte(ws, ch) != -1 || bytes.IndexByte(delim, ch) != -1
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch byte) bool {
	if unicode.IsSpace(rune(ch)) {
		return true
	}
	ws := []byte{' ', '\f', '\n', '\r', '\t'}
	return bytes.IndexByte(ws, ch) != -1
}

// acceptNumber returns a number if it can.
func acceptNumber(input []byte) (run, rest []byte) {
	if len(input) == 0 {
		return nil, input
	}
	pos, rest := 0, input
	if isdigit(rest[0]) {
		// consume the digit
		pos, rest = 1, rest[1:]
	} else if (input[0] == '+' || input[0] == '-') && isdigit(input[1]) {
		// consume the plus/minus and digit
		pos, rest = 2, rest[2:]
	} else {
		// not a number
		return nil, input
	}
	// consume the run of digits
	for len(rest) != 0 && isdigit(rest[0]) {
		pos, rest = pos+1, rest[1:]
	}
	// number must be delimited with space, dot, comma, or end-of-input
	if len(rest) != 0 {
		delim, _ := utf8.DecodeRune(rest)
		validDelimiter := unicode.IsSpace(delim) || delim == ',' || delim == '.'
		if !validDelimiter {
			return nil, input
		}
	}
	return input[:pos], input[pos:]
}

// acceptNumberWithCommas returns a number with commas if it can.
func acceptNumberWithCommas(input []byte) (run, rest []byte) {
	if len(input) == 0 || !isdigit(input[0]) {
		return nil, input
	}
	pos, rest := 0, input
	// first grouping must be 1, 2, or 3 <digit>s
	for pos < 3 && len(rest) != 0 && isdigit(rest[0]) {
		// consume the digit
		pos, rest = pos+1, rest[1:]
	}
	// all remaining groups must be ',<digit><digit><digit>'
	for len(rest) > 3 && rest[0] == ',' && isdigit(rest[1]) && isdigit(rest[2]) && isdigit(rest[3]) {
		// consume the comma and three digits
		pos, rest = pos+4, rest[4:]
	}
	// number must be delimited with space, dot, or end-of-input
	if len(rest) != 0 {
		delim, _ := utf8.DecodeRune(rest)
		validDelimiter := unicode.IsSpace(delim) || delim == '.'
		if !validDelimiter {
			return nil, input
		}
	}
	return input[:pos], input[pos:]
}

// acceptTurnNumber returns a TURN_NUMBER (YYY-MM) if it can.
func acceptTurnNumber(input []byte) (run, rest []byte) {
	if len(input) == 0 {
		return nil, input
	}
	pos, rest := 0, input
	if len(rest) < 3 || !(isdigit(rest[0]) && isdigit(rest[1]) && isdigit(rest[2])) {
		return nil, input
	}
	pos, rest = pos+3, rest[3:]
	if len(rest) == 0 || rest[0] != '-' {
		return nil, input
	}
	pos, rest = pos+1, rest[1:]
	if len(rest) < 2 || !(isdigit(rest[0]) && isdigit(rest[1])) {
		return nil, input
	}
	pos, rest = pos+2, rest[2:]
	// yyy-mm must be delimited with space or end-of-input
	if len(rest) != 0 {
		delim, _ := utf8.DecodeRune(rest)
		validDelimiter := unicode.IsSpace(delim)
		if !validDelimiter {
			return nil, input
		}
	}
	return input[:pos], input[pos:]
}

func isdigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isspace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func runOfDigits(input []byte) (value, rest []byte) {
	pos := 0
	for pos < len(input) {
		r, w := utf8.DecodeRune(input[pos:])
		if !unicode.IsDigit(r) {
			break
		}
		pos += w
	}
	if pos == 0 {
		return nil, input
	}
	return input[:pos], input[pos:]
}

func runOfLetters(input []byte) (value, rest []byte) {
	pos := 0
	for pos < len(input) {
		r, w := utf8.DecodeRune(input[pos:])
		if !unicode.IsLetter(r) {
			break
		}
		pos += w
	}
	if pos == 0 {
		return nil, input
	}
	return input[:pos], input[pos:]
}

func runOfSpaces(input []byte) (value, rest []byte) {
	pos := 0
	for pos < len(input) {
		r, w := utf8.DecodeRune(input[pos:])
		if !unicode.IsSpace(r) {
			break
		}
		pos += w
	}
	if pos == 0 {
		return nil, input
	}
	return input[:pos], input[pos:]
}
