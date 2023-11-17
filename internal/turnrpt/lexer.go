// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"bytes"
	"unicode"
)

var (
	delimiters = []byte{' ', '\f', '\n', '\r', '\t', ',', '(', ')'}
)

func currency(input []byte) (run, rest []byte) {
	if len(input) == 0 {
		return nil, nil
	} else if input[0] != '$' {
		return nil, input
	}
	rest = input[1:]
	if len(rest) == 0 {
		return nil, input
	} else if isSpace(rest[0]) {
		rest = rest[1:]
	}
	pos := 0
	for pos < len(rest) && isDigit(rest[pos]) {
		pos = pos + 1
	}
	if pos < len(rest) && rest[pos] == '.' {
		pos = pos + 1
		for pos < len(rest) && isDigit(rest[pos]) {
			pos = pos + 1
		}
	}
	run, rest = rest[:pos], rest[pos:]
	if len(rest) == 0 || bytes.IndexByte([]byte{' ', '\f', '\n', '\r', '\t', ','}, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

func ddMmYyyy(input []byte) (run, rest []byte) {
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
	if len(rest) == 0 || bytes.IndexByte(delimiters, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

// hexNo accepts a hex number in the form "NN XXYY"
// where "NN" ie either "##" or an integer, and "XXYY" is an integer.
func hexNo(input []byte) (run, rest []byte) {
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
	if len(rest) != 0 && bytes.IndexByte(delimiters, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

func integer(input []byte) (run, rest []byte) {
	if len(input) == 0 {
		return nil, nil
	}
	pos := 0
	for pos < len(input) && isDigit(input[pos]) {
		pos++
	}
	run, rest = input[:pos], input[pos:]
	if len(rest) != 0 && bytes.IndexByte(delimiters, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

func literal(literal, input []byte) (run, rest []byte) {
	if !bytes.HasPrefix(input, literal) {
		return nil, input
	}
	return input[:len(literal)], input[len(literal):]
}

func keyword(kw, input []byte) (run, rest []byte) {
	if len(input) < len(kw) {
		return nil, input
	} else if !bytes.HasPrefix(input, kw) {
		return nil, input
	}
	run, rest = input[:len(kw)], input[len(kw):]
	if len(rest) == 0 || bytes.IndexByte(delimiters, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

func season(input []byte) (run, rest []byte) {
	for _, season := range [][]byte{[]byte("Winter"), []byte("Spring")} {
		if run, rest := keyword(season, input); run != nil {
			return run, rest
		}
	}
	return nil, input
}

func skipws(input []byte) []byte {
	for len(input) != 0 && isSpace(input[0]) {
		input = input[1:]
	}
	return input
}

func toeol(input []byte) (run, rest []byte) {
	if len(input) == 0 {
		return nil, nil
	}
	pos := 0
	for pos < len(input) && input[pos] != '\n' {
		pos++
	}
	run, rest = input[:pos], input[pos:]
	return run, rest
}

func tribeId(input []byte) (run, rest []byte) {
	if len(input) < 4 {
		return nil, input
	}
	for n := 0; n < 4; n++ {
		if !isDigit(input[n]) {
			return nil, input
		}
	}
	run, rest = input[:4], input[4:]
	if len(rest) == 0 || bytes.IndexByte(delimiters, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

func turnNo(input []byte) (run, rest []byte) {
	if len(input) < 6 {
		return nil, input
	} else if !(isDigit(input[0]) && isDigit(input[1]) && isDigit(input[2])) {
		return nil, input
	} else if input[3] != '-' {
		return nil, input
	} else if !(isDigit(input[4]) && isDigit(input[5])) {
		return nil, input
	}
	run, rest = input[:6], input[6:]
	if len(rest) == 0 || bytes.IndexByte(delimiters, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}

func weather(input []byte) (run, rest []byte) {
	for _, weather := range [][]byte{[]byte("FINE")} {
		if run, rest := keyword(weather, input); run != nil {
			return run, rest
		}
	}
	return nil, input
}

func whitespace(input []byte) (run, rest []byte) {
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

func word(input []byte) (run, rest []byte) {
	pos := 0
	for pos < len(input) && (unicode.IsLetter(rune(input[pos])) || unicode.IsDigit(rune(input[pos]))) {
		pos++
	}
	if pos == 0 {
		return nil, input
	}
	run, rest = input[:pos], input[pos:]
	if len(rest) != 0 && bytes.IndexByte(delimiters, rest[0]) == -1 {
		return nil, input
	}
	return run, rest
}
