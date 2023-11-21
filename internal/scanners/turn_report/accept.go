// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package scanner

import "bytes"

func acceptBulletNumber(input []byte) []byte {
	if len(input) < 3 {
		return nil
	} else if !(input[0] == '(' && isdigit(input[1]) && input[2] == ')') {
		return nil
	}
	return input[:3]
}

func acceptCantMoveOnOceanTo(input []byte) []byte {
	if !bytes.HasPrefix(input, []byte("Can't Move on Ocean to ")) {
		return nil
	}
	return input[:23]
}

// currency looks like "$N", "$N.NN", "$ N", or "$ N.NN" (the dot and cents are optional)
func acceptCurrency(input []byte) []byte {
	if len(input) < 2 { // minimum length is two characters (eg, $0)
		return nil
	} else if input[0] != '$' {
		return nil
	}
	pos := 1 // skip the '$'
	// $ may be followed by space
	if pos < len(input) && isspace(input[pos]) {
		pos++
	}
	dollars, cents := 0, 0
	for pos < len(input) && isdigit(input[pos]) {
		pos, dollars = pos+1, dollars+1
	}
	// dollars may be followed by '.' then one or two cents digits
	if pos < len(input) && input[pos] == '.' {
		pos++
		for pos < len(input) && isdigit(input[pos]) {
			pos, cents = pos+1, cents+1
		}
	}
	// check the number of digits and cents
	if !(dollars > 0 && cents <= 2) {
		return nil // not a valid currency amount
	}
	if pos < len(input) && !(isspace(input[pos]) || input[pos] == ',') { // invalid delimiter
		return nil
	}
	return input[:pos]
}

func acceptCommodity(input []byte) []byte {
	if len(input) == 0 {
		return nil
	}
	for _, commodity := range []string{"Coffee", "Frankincense"} {
		lc := len(commodity)
		if bytes.HasPrefix(input, []byte(commodity)) {
			if len(input) == lc {
				return input[:lc]
			} else if input[lc] == ',' || input[lc] == '-' || isspace(input[lc]) {
				return input[:lc]
			}
		}
	}
	return nil
}

// date looks like dd/mm/yyyy
func acceptDate(input []byte) []byte {
	if len(input) < 10 {
		return nil
	} else if !(isdigit(input[0]) && isdigit(input[1])) {
		return nil
	} else if input[2] != '/' {
		return nil
	} else if !(isdigit(input[3]) && isdigit(input[4])) {
		return nil
	} else if input[5] != '/' {
		return nil
	} else if !(isdigit(input[6]) && isdigit(input[7]) && isdigit(input[8]) && isdigit(input[9])) {
		return nil
	}
	return input[:10]
}

func acceptDirection(input []byte) []byte {
	if len(input) == 0 {
		return nil
	}
	for _, code := range []string{"N", "NE", "SE", "S", "SW", "NW"} {
		lc := len(code)
		if bytes.HasPrefix(input, []byte(code)) {
			if len(input) == lc {
				return input[:lc]
			} else if input[lc] == '\\' || input[lc] == ',' || input[lc] == '-' || isspace(input[lc]) {
				return input[:lc]
			}
		}
	}
	return nil
}

// must be "## XXYY" or "NN XXYY" for hex number
func acceptHexNo(input []byte) []byte {
	if len(input) < 7 {
		return nil
	} else if !((isdigit(input[0]) && isdigit(input[1])) || (input[0] == '#' || input[1] == '#')) {
		return nil
	} else if !isspace(input[2]) {
		return nil
	} else if !(isdigit(input[3]) && isdigit(input[4]) && isdigit(input[5]) && isdigit(input[6])) {
		return nil
	} else if !(input[7] == ',' || input[7] == ')') {
		return nil
	}
	return input[:7]
}

func acceptInto(input []byte) []byte {
	if !bytes.HasPrefix(input, []byte("into ")) {
		return nil
	}
	return input[:4]
}

func acceptNoFordOnRiver(input []byte) []byte {
	if !bytes.HasPrefix(input, []byte("No Ford on River to ")) {
		return nil
	}
	return input[:19]
}

func acceptNotEnoughMPs(input []byte) []byte {
	if !bytes.HasPrefix(input, []byte("Not enough M.P's to move to ")) {
		return nil
	}
	return input[:27]
}

func acceptNothingOfInterestFound(input []byte) []byte {
	if !bytes.HasPrefix(input, []byte("Nothing of interest found")) {
		return nil
	}
	return input[:25]
}

func acceptOfHex(input []byte) []byte {
	if bytes.HasPrefix(input, []byte("of HEX ")) {
		return input[:6]
	} else if bytes.HasPrefix(input, []byte("of HEX,")) {
		return input[:6]
	}
	return nil
}

func acceptReceived(input []byte) []byte {
	if !bytes.HasPrefix(input, []byte("Received: ")) {
		return nil
	}
	return input[:8]
}

func acceptScoutSectionStart(input []byte) []byte {
	if !bytes.HasPrefix(input, []byte("Scout 1:")) {
		return nil
	}
	return input[:7]
}

func acceptScoutNo(input []byte) []byte {
	pos := 0
	for pos < len(input) && isdigit(input[pos]) {
		pos++
	}
	if pos == 0 {
		return nil
	} else if pos < len(input) && input[pos] != ':' {
		return nil
	}
	return input[:pos]
}

func acceptTerrain(input []byte) []byte {
	if len(input) == 0 {
		return nil
	}
	for _, code := range []string{
		"ALPS",
		"AR",
		"BH",
		"BR",
		"CH", "CONIFER HILLS",
		"DE",
		"DF",
		"DH",
		"Fords",
		"GH", "GRASSY HILLS",
		"HSM",
		"JG",
		"JH",
		"L",
		"LCM",
		"LJM",
		"LSM",
		"O", "Ocean", "OCEAN",
		"PI",
		"PR", "PRAIRIE",
		"R", "River", "RIVER",
		"RH", "ROCKY HILLS",
		"SH",
		"SW",
		"TU",
	} {
		lc := len(code)
		if bytes.HasPrefix(input, []byte(code)) {
			if len(input) == lc {
				return input[:lc]
			} else if input[lc] == '\\' || input[lc] == ',' || isspace(input[lc]) {
				return input[:lc]
			}
		}
	}
	return nil
}

func acceptTribeActivities(input []byte) []byte {
	if !bytes.HasPrefix(input, []byte("Tribe Activities:")) {
		return nil
	}
	return input[:16]
}

func acceptTribeNo(input []byte) []byte {
	if bytes.HasPrefix(input, []byte("No GT")) {
		return input[:5]
	}
	if len(input) < 4 {
		return nil
	} else if !(isdigit(input[0]) && isdigit(input[1]) && isdigit(input[2]) && isdigit(input[3])) {
		return nil
	} else if !(len(input) == 4 || isspace(input[4]) || input[4] == ',') {
		return nil
	}
	return input[:4]
}

// must be "#N" or "#NN"
func acceptTurnMonth(input []byte) []byte {
	if len(input) > 2 && (input[0] == '#' && isdigit(input[1]) && input[2] == ')') {
		return input[:2]
	} else if len(input) > 3 && (input[0] == '#' && isdigit(input[1]) && isdigit(input[2]) && input[3] == ')') {
		return input[:3]
	}
	return nil
}

func acceptTurnNo(input []byte) []byte {
	if len(input) < 6 {
		return nil
	} else if !(isdigit(input[0]) && isdigit(input[1]) && isdigit(input[2])) {
		return nil
	} else if input[3] != '-' {
		return nil
	} else if !(isdigit(input[4]) && isdigit(input[5])) {
		return nil
	}
	return input[:6]
}
