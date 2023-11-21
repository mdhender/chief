// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package scanner

import (
	"bytes"
)

func (s *Scanner) nextScout() Token {
	if s.iseof() {
		return Token{Type: EOF}
	} else if acceptTribeActivities(s.input[s.pos:]) != nil {
		return Token{Type: EndOfSection}
	} else if bytes.HasPrefix(s.input[s.pos:], s.ClanStatusHeading) {
		return Token{Type: EndOfSection}
	}

	switch ch := s.getch(); ch {
	case '\\':
		return Token{Type: BackSlash, Value: s.accept()}
	case ':':
		return Token{Type: Colon, Value: s.accept()}
	case ',':
		return Token{Type: Comma, Value: s.accept()}
	case '-':
		return Token{Type: Dash, Value: s.accept()}
	default:
		s.ungetch()
	}

	if isdigit(s.peekch()) {
		if val := acceptScoutNo(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			return Token{Type: ScoutNo, Value: s.accept()}
		}
		return s.error()
	}

	if isalpha(s.peekch()) {
		if val := acceptDirection(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			return Token{Type: Direction, Value: s.accept()}
		} else if val = acceptInto(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			s.ignore()
			return Token{Type: Into}
		} else if val = acceptCantMoveOnOceanTo(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			s.ignore()
			return Token{Type: CantMoveOnOcean}
		} else if val = acceptNoFordOnRiver(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			s.ignore()
			return Token{Type: NoFordOnRiver}
		} else if val = acceptNotEnoughMPs(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			s.ignore()
			return Token{Type: NotEnoughMPs}
		} else if val = acceptNothingOfInterestFound(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			s.ignore()
			return Token{Type: NothingOfInterest}
		} else if val = acceptOfHex(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			s.ignore()
			return Token{Type: OfHex}
		} else if val := acceptTerrain(s.input[s.pos:]); val != nil {
			s.pos += len(val)
			return Token{Type: Terrain, Value: s.accept()}
		}
		// return a literal or unknown token
		for !s.iseof() && isalpha(s.peekch()) {
			s.getch()
		}
		val := s.accept()
		switch string(val) {
		case "Scout":
			return Token{Type: Scout}
		}
		return Token{Type: Unknown, Value: val}
	}
	// return an error
	s.getch()
	return s.error()
}

func (s *Scanner) nextStatus() Token {
	if s.iseof() {
		return Token{Type: EOF}
	} else if acceptScoutSectionStart(s.input[s.pos:]) != nil {
		return Token{Type: EndOfSection}
	} else if acceptTribeActivities(s.input[s.pos:]) != nil {
		return Token{Type: EndOfSection}
	}

	ch := s.getch()
	switch ch {
	case ',':
		return Token{Type: Comma, Value: s.accept()}
	case '=':
		return Token{Type: EqualSign, Value: s.accept()}
	case ':':
		return Token{Type: Colon, Value: s.accept()}
	case '(':
		if val := acceptBulletNumber(s.input[s.pos-1:]); val != nil {
			s.pos += len(val) - 1
			return Token{Type: BulletNumber, Value: s.accept()}
		}
		return Token{Type: LeftParen, Value: s.accept()}
	case ')':
		return Token{Type: RightParen, Value: s.accept()}
	case '#': // hex number or month number
		if val := acceptHexNo(s.input[s.pos-1:]); val != nil {
			s.pos += len(val) - 1
			return Token{Type: HexNo, Value: s.accept()}
		} else if val = acceptTurnMonth(s.input[s.pos-1:]); val != nil {
			s.pos += len(val) - 1
			return Token{Type: TurnMonth, Value: s.accept()}
		}
		return s.error()
	case '$': // currency
		if val := acceptCurrency(s.input[s.pos-1:]); val != nil {
			s.pos += len(val) - 1
			return Token{Type: Currency, Value: s.accept()}
		}
		return s.error()
	}
	if isdigit(ch) { // tribe number, turn number
		if val := acceptDate(s.input[s.pos-1:]); val != nil {
			s.pos += len(val) - 1
			return Token{Type: Date, Value: s.accept()}
		} else if val = acceptTribeNo(s.input[s.pos-1:]); val != nil {
			s.pos += len(val) - 1
			return Token{Type: TribeNo, Value: s.accept()}
		} else if val = acceptTurnNo(s.input[s.pos-1:]); val != nil {
			s.pos += len(val) - 1
			return Token{Type: TurnNo, Value: s.accept()}
		}
		return s.error()
	}
	s.ungetch()
	if bytes.HasPrefix(s.input[s.pos:], []byte("Current Hex")) {
		s.pos += 11
		s.accept()
		return Token{Type: CurrentHex}
	} else if bytes.HasPrefix(s.input[s.pos:], []byte("Current Turn")) {
		s.pos += 12
		s.accept()
		return Token{Type: CurrentTurn}
	} else if bytes.HasPrefix(s.input[s.pos:], []byte("Desired Commodities")) {
		s.pos += 19
		s.accept()
		return Token{Type: DesiredCommodities}
	} else if bytes.HasPrefix(s.input[s.pos:], []byte("Goods Tribe")) {
		s.pos += 11
		s.accept()
		return Token{Type: GoodsTribe}
	} else if bytes.HasPrefix(s.input[s.pos:], []byte("Next Turn")) {
		s.pos += 9
		s.accept()
		return Token{Type: NextTurn}
	} else if bytes.HasPrefix(s.input[s.pos:], []byte("No commodities allocated")) {
		s.pos += 24
		s.accept()
		return Token{Type: NoDesiredCommodities}
	} else if bytes.HasPrefix(s.input[s.pos:], []byte("No GT")) {
		s.pos += 5
		return Token{Type: TribeNo, Value: s.accept()}
	} else if bytes.HasPrefix(s.input[s.pos:], []byte("Previous Hex")) {
		s.pos += 12
		s.accept()
		return Token{Type: PreviousHex}
	} else if val := acceptCommodity(s.input[s.pos:]); val != nil {
		s.pos += len(val)
		return Token{Type: Commodity, Value: s.accept()}
	} else if val = acceptReceived(s.input[s.pos:]); val != nil {
		s.pos += len(val)
		s.accept()
		return Token{Type: Received}
	}
	for !s.iseof() && isalpha(s.peekch()) {
		s.getch()
	}
	val := s.accept()
	switch string(val) {
	case "Cost":
		return Token{Type: Cost}
	case "Credit":
		return Token{Type: Credit}
	case "FINE":
		return Token{Type: Weather, Value: val}
	case "Received":
		return Token{Type: Received, Value: val}
	case "Spring":
		return Token{Type: Season, Value: val}
	case "Tribe":
		return Token{Type: Tribe}
	case "Winter":
		return Token{Type: Season, Value: val}
	}
	return Token{Type: Unknown, Value: val}
}
