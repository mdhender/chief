// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"bytes"
	"fmt"
)

func nextToken(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) == 0 {
		return nil, input
	}
	var pos int
	for pos = 0; pos < len(input) && !isspace(input[pos]); pos++ {
		//
	}
	return input[:pos], input[pos:]
}

func lexToken(input []byte) string {
	_, input = lexSpaces(input)
	var val string
	for len(input) > 0 && len(val) < 17 {
		if isspace(input[0]) {
			val += " "
			_, input = lexSpaces(input)
		} else {
			val += string(input[0])
			input = input[1:]
		}
	}
	return val
}

func lexAmount(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) == 0 || input[0] != '$' {
		return nil, input
	}
	saved := input
	_, input = lexSpaces(input[1:])
	if len(input) == 0 || !isdigit(input[0]) {
		return nil, saved
	}
	var dollars, cents []byte
	pos := 0
	for pos < len(input) && isdigit(input[pos]) {
		pos++
	}
	dollars, input = append(dollars, input[:pos]...), input[pos:]
	if len(input) > 0 && input[0] == '.' {
		input = input[1:]
		pos = 0
		for pos < len(input) && isdigit(input[pos]) {
			pos++
		}
		cents, input = input[:pos], input[pos:]
	}
	if len(cents) == 0 {
		dollars = append(dollars, '.', '0', '0')
	} else if len(cents) == 1 {
		cents = []byte{cents[0], '0'}
		dollars = append(dollars, '.', cents[0], '0')
	} else {
		dollars = append(dollars, '.', cents[0], cents[1])
	}
	return dollars, input
}

func lexClanId(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	pfx := []byte("Tribe ")
	if !bytes.HasPrefix(input, pfx) {
		return nil, input
	}
	token, rest = lexTribeNo(input[len(pfx):])
	if token == nil || token[0] != '0' {
		return nil, input
	}
	return token, rest
}

func lexCost(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	pfx := []byte("Cost: ")
	if !bytes.HasPrefix(input, pfx) {
		return nil, input
	}
	token, rest = lexAmount(input[len(pfx):])
	if token == nil {
		return nil, input
	}
	return token, rest
}

func lexCredit(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	pfx := []byte("Credit: ")
	if !bytes.HasPrefix(input, pfx) {
		return nil, input
	}
	token, rest = lexAmount(input[len(pfx):])
	if token == nil {
		return nil, input
	}
	return token, rest
}

func lexCurrentHex(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	pfx := []byte("Current Hex = ")
	if !bytes.HasPrefix(input, pfx) {
		return nil, input
	}
	token, rest = lexHexNo(input[len(pfx):])
	if token == nil {
		return nil, input
	}
	return token, rest
}

func lexCurrentTurn(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	pfx := []byte("Current Turn ")
	if !bytes.HasPrefix(input, pfx) {
		return nil, input
	}
	token, rest = lexTurnNo(input[len(pfx):])
	if token == nil {
		return nil, input
	}
	return token, rest
}

func lexDelimiter(input []byte, ch byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) == 0 || input[0] != ch {
		return nil, input
	}
	return input[:1], input[1:]
}

func lexDeposit(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	// upper-cased for scouting
	for _, kw := range [][]byte{{'C', 'O', 'A', 'L'}, {'I', 'R', 'O', 'N'}, {'Z', 'I', 'N', 'C'}} {
		if bytes.HasPrefix(input, kw) {
			if len(input) == len(kw) || isdelim(input[len(kw)]) {
				return input[:len(kw)], input[len(kw):]
			}
		}
	}
	return nil, input
}

func lexDirection(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	// dashed variants
	for _, kw := range [][]byte{{'N', '-'}, {'N', 'E', '-'}, {'S', 'E', '-'}, {'S', '-'}, {'S', 'W', '-'}, {'N', 'W', '-'}} {
		if bytes.HasPrefix(input, kw) {
			return input[:len(kw)-1], input[len(kw):]
		}
	}
	// spaced variants
	for _, kw := range [][]byte{{'N', ' '}, {'N', 'E', ' '}, {'S', 'E', ' '}, {'S', ' '}, {'S', 'W', ' '}, {'N', 'W', ' '}} {
		if bytes.HasPrefix(input, kw) {
			return input[:len(kw)-1], input[len(kw):]
		}
	}
	return nil, input
}

// lexFirstLiteral returns the prefix to the first literal, the
// remainder of the input, and the index of the literal (if found).
func lexFirstLiteral(input []byte, literals ...string) (token, rest []byte) {
	_, input = lexSpaces(input)
	for _, literal := range literals {
		if bytes.HasPrefix(input, []byte(literal)) {
			return input[:len(literal)], input[len(literal):]
		}
	}
	return nil, input
}

func lexGoods(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	for _, literal := range [][]byte{[]byte("Coffee"), []byte("Frankincense")} {
		if !bytes.HasPrefix(input, literal) {
			continue
		}
		return input[:len(literal)], input[len(literal):]
	}
	return nil, input
}

func lexHexNo(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) < 7 {
		return nil, input
	} else if !((input[0] == '#' && input[1] == '#') || (isupper(input[0]) && isupper(input[1]))) {
		return nil, input
	} else if !isspace(input[2]) {
		return nil, input
	} else if !isdigit(input[3]) || !isdigit(input[4]) || !isdigit(input[5]) || !isdigit(input[6]) {
		return nil, input
	}
	return input[:7], input[7:]
}

func lexLiteral(input []byte, literal []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) == 0 || !bytes.HasPrefix(input, literal) {
		return nil, input
	}
	return input[:len(literal)], input[len(literal):]
}

func lexMonth(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) < 3 || input[0] != '#' || !isdigit(input[1]) {
		return nil, input
	} else if input[2] == ')' {
		return []byte{'0', input[1]}, input[2:]
	} else if len(input) > 3 && isdigit(input[2]) && input[3] == ')' {
		return input[1:3], input[3:]
	}
	return nil, input
}

func lexNextTurn(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	pfx := []byte("Next Turn ")
	if !bytes.HasPrefix(input, pfx) {
		return nil, input
	}
	token, rest = lexTurnNo(input[len(pfx):])
	if token == nil {
		return nil, input
	}
	return token, rest
}

func lexPreviousHex(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	pfx := []byte("Previous Hex = ")
	if !bytes.HasPrefix(input, pfx) {
		return nil, input
	}
	token, rest = lexHexNo(input[len(pfx):])
	if token == nil {
		return nil, input
	}
	return token, rest
}

func lexReceived(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	pfx := []byte("Received: ")
	if !bytes.HasPrefix(input, pfx) {
		return nil, input
	}
	token, rest = lexAmount(input[len(pfx):])
	if token == nil {
		return nil, input
	}
	return token, rest
}

func lexSeason(input []byte) (run, rest []byte) {
	_, input = lexSpaces(input)
	for _, literal := range [][]byte{[]byte("Winter"), []byte("Spring")} {
		if !bytes.HasPrefix(input, literal) {
			continue
		}
		return input[:len(literal)], input[len(literal):]
	}
	return nil, input
}

func lexSpaces(input []byte) (token, rest []byte) {
	if len(input) == 0 || !isspace(input[0]) {
		return nil, input
	}
	pos := 0
	for pos < len(input) && isspace(input[pos]) {
		pos++
	}
	return input[:pos], input[pos:]
}

func lexTerrain(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	for _, code := range []string{
		"ALPS", "AR", "BH", "BR", "CH", "DE", "DF", "DH",
		"FORDS", "GH", "HSM", "JG", "JH", "L", "LCM", "LJM",
		"LSM", "O", "PI", "PR", "R", "RH", "SH", "SW", "TU",
	} {
		lc := len(code)
		if bytes.HasPrefix(input, []byte(code)) {
			if len(input) == lc || input[lc] == ',' || isspace(input[lc]) {
				return input[:lc], input[lc:]
			}
		}
	}
	// aliases
	for _, code := range []string{
		"CONIFER HILLS",
		"GRASSY HILLS",
		"OCEAN",
		"PRAIRIE",
		"RIVER",
		"ROCKY HILLS",
	} {
		lc := len(code)
		if bytes.HasPrefix(input, []byte(code)) {
			if len(input) == lc || input[lc] == ',' || isspace(input[lc]) {
				switch code {
				case "CONIFER HILLS":
					return []byte{'C', 'H'}, input[lc:]
				case "GRASSY HILLS":
					return []byte{'G', 'H'}, input[lc:]
				case "OCEAN":
					return []byte{'O'}, input[lc:]
				case "PRAIRIE":
					return []byte{'P', 'R'}, input[lc:]
				case "RIVER":
					return []byte{'R'}, input[lc:]
				case "ROCKY HILLS":
					return []byte{'R', 'H'}, input[lc:]
				default:
					panic(fmt.Sprintf("assert(code != %q)", code))
				}
			}
		}
	}
	return nil, input
}

func lexTribeElement(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) < 6 {
		return nil, input
	} else if !(isdigit(input[0]) && isdigit(input[1]) && isdigit(input[2]) && isdigit(input[3])) {
		return nil, input
	} else if !islower(input[4]) {
		return nil, input
	} else if !isdigit(input[5]) {
		return nil, input
	} else if !(len(input) == 6 || isdelim(input[6])) {
		return nil, input
	}
	return input[:6], input[6:]
}

func lexTribeNo(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) < 4 {
		return nil, input
	} else if !(isdigit(input[0]) && isdigit(input[1]) && isdigit(input[2]) && isdigit(input[3])) {
		return nil, input
	} else if !(len(input) == 4 || isdelim(input[4])) {
		return nil, input
	}
	return input[:4], input[4:]
}

func lexTurnDate(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) < 10 {
		return nil, input
	} else if !(isdigit(input[0]) && isdigit(input[1])) {
		return nil, input
	} else if input[2] != '/' {
		return nil, input
	} else if !(isdigit(input[3]) && isdigit(input[4])) {
		return nil, input
	} else if input[5] != '/' {
		return nil, input
	} else if !(isdigit(input[6]) && isdigit(input[7]) && isdigit(input[8]) && isdigit(input[9])) {
		return nil, input
	}
	return []byte{input[6], input[7], input[8], input[9], input[5], input[3], input[4], input[2], input[0], input[1]}, input[10:]
}

func lexTurnNo(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	if len(input) < 6 {
		return nil, input
	} else if !(isdigit(input[0]) && isdigit(input[1]) && isdigit(input[2])) {
		return nil, input
	} else if input[3] != '-' {
		return nil, input
	} else if !(isdigit(input[4]) && isdigit(input[5])) {
		return nil, input
	}
	return input[:6], input[6:]
}

func lexUnit(input []byte) (token, rest []byte) {
	token, rest = lexTribeNo(input)
	if token == nil {
		token, rest = lexTribeElement(input)
	}
	return token, rest
}

func lexWeather(input []byte) (token, rest []byte) {
	_, input = lexSpaces(input)
	for _, literal := range [][]byte{[]byte("FINE")} {
		if !bytes.HasPrefix(input, literal) {
			continue
		}
		return input[:len(literal)], input[len(literal):]
	}
	return nil, input
}

func isdelim(ch byte) bool {
	return ch == ' ' || ch == ',' || ch == '\\'
}
func isdigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func islower(ch byte) bool {
	return 'a' <= ch && ch <= 'z'
}

func isspace(ch byte) bool {
	return ch == ' ' || ch == '\f' || ch == '\n' || ch == '\r' || ch == '\t'
}

func isupper(ch byte) bool {
	return 'A' <= ch && ch <= 'Z'
}

func runToEOL(input []byte) (token, rest []byte) {
	if len(input) == 0 {
		return nil, input
	}
	pos := 0
	for pos < len(input) && input[pos] != '\n' {
		pos++
	}
	if pos < len(input) && input[pos] == '\n' {
		pos++
	}
	return bytes.TrimSpace(input[:pos]), input[pos:]
}

// runToFirstLiteral returns the prefix to the first literal, the
// remainder of the input, and the index of the literal (if found).
// "First" means the literal with the least offset into the input.
// Warning: this can be an expensive search.
func runToFirstLiteral(input []byte, literals ...string) (token, rest []byte, index int) {
	_, input = lexSpaces(input)
	offset := -1
	for i, literal := range literals {
		if pos := bytes.Index(input, []byte(literal)); pos != -1 {
			if offset == -1 || pos < offset {
				index, offset = i, pos
			}
		}
	}
	if offset == -1 {
		return nil, input, -1
	}
	return input[:offset], input[offset:], index
}
