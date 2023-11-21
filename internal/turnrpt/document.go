// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"fmt"
	str "github.com/mdhender/chief/internal/scanners/turn_report"
	"log"
)

type parser struct {
	scanner *str.Scanner
}

func ParseDocument(filename string) ([]*TribeSection, error) {
	p := &parser{}

	log.Printf("parse: filename %s\n", filename)

	if s, err := str.NewTurnReportScanner(filename); err != nil {
		log.Fatal(err)
	} else {
		p.scanner = s
	}

	if err := p.document(); err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}

	return nil, fmt.Errorf("!implemented")
}

func (p *parser) document() error {
	// status section includes next-turn section
	log.Printf("[parse] unit status section\n")
	state, priorTokens := str.StStatus, [2]str.Type{str.Unknown, str.Unknown}
	for {
		tok := p.scanner.Next(state)
		log.Printf("tok %+v\n", tok)
		if tok.Type == str.Error {
			return fmt.Errorf("unit status: unexpected %q", string(tok.Value))
		} else if tok.Type == str.Unknown {
			return fmt.Errorf("unit status: unexpected %q", string(tok.Value))
		}

		if tok.Is(str.TribeNo) {
			if priorTokens[0] == str.Tribe && priorTokens[1] == str.Unknown {
				p.scanner.ClanStatusHeading = []byte(fmt.Sprintf("%s Status:", string(tok.Value)))
			}
		}

		priorTokens[0], priorTokens[1] = tok.Type, priorTokens[0]

		if tok.Type == str.EOF || tok.Type == str.Error || tok.Type == str.Unknown {
			break
		}
		if tok.Type == str.EndOfSection {
			log.Printf("[parse] %d: found end of section\n", state)
			break
		}
	}
	if p.scanner.ClanStatusHeading == nil {
		return fmt.Errorf("could not find clan in status section")
	}
	//switch state {
	//case str.StStatus:
	//	switch tok.Type {
	//	case str.Currency:
	//		if priorTokens[1] == str.Cost {
	//			state = str.StNextTurn
	//			log.Printf("state status::cost>currency -> nextTurn\n")
	//		}
	//	case str.Weather:
	//		state = str.StNextTurn
	//		log.Printf("state status::weather -> nextTurn\n")
	//	}
	//case str.StNextTurn:
	//	switch tok.Type {
	//	case str.Currency:
	//		state = str.StStatus
	//		log.Printf("state nextTurn::currency -> status\n")
	//	case str.Date:
	//		state = str.StStatus
	//		log.Printf("state nextTurn::weather -> status\n")
	//	}
	//}

	log.Printf("[parse] scouting results section\n")
	state, priorTokens = str.StScouting, [2]str.Type{str.Unknown, str.Unknown}
	for {
		tok := p.scanner.Next(state)
		log.Printf("tok %+v\n", tok)
		if tok.Type == str.Error {
			return fmt.Errorf("unit status: unexpected %q", string(tok.Value))
		} else if tok.Type == str.Unknown {
			return fmt.Errorf("unit status: unexpected %q", string(tok.Value))
		}

		priorTokens[0], priorTokens[1] = tok.Type, priorTokens[0]

		if tok.Type == str.EndOfSection {
			log.Printf("[parse] %d: found end of section\n", state)
			break
		}
		if tok.Type == str.EOF || tok.Type == str.Error || tok.Type == str.Unknown {
			break
		}
	}

	if !p.scanner.IsEof() {
		return fmt.Errorf("extra tokens in input")
	}

	return fmt.Errorf("!implemented")
}
