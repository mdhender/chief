// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"bytes"
	"fmt"
	"github.com/mdhender/chief/internal/docconv"
	"log"
	"os"
	"runtime"
)

func Parse(filename string) (*Report, error) {
	log.Printf("[report] filename %s\n", filename)
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var input []byte
	if body, _, err := docconv.ConvertDocx(fp); err != nil {
		return nil, err
	} else {
		input = []byte(body)
	}
	r := &Report{FileName: filename, input: input}

	if err := r.parseClanStatus(); err != nil {
		return nil, fmt.Errorf("clan status: %w", err)
	} else if err = r.parseTurnDetails(); err != nil {
		return nil, fmt.Errorf("turn details: %w", err)
	} else if err = r.parseAccounting(); err != nil {
		return nil, fmt.Errorf("accounting: %w", err)
	} else if err = r.parseGoodsTribe(); err != nil {
		return nil, fmt.Errorf("goods tribe: %w", err)
	} else if err = r.parseDesiredCommodities(); err != nil {
		return nil, fmt.Errorf("desired commodities: %w", err)
	}

	// the clan section has optional notes from the GM.
	if err := r.parseGMNotes(); err != nil {
		return nil, fmt.Errorf("gm notes: %w", err)
	}

	// the clan section gets a little goofy here because the first
	// turn report skips the activity heading and includes only scouting.
	if err := r.parseOptionalScoutReports(); err != nil {
		return nil, fmt.Errorf("scouting: %w", err)
	}

	return r, nil
}

func (r *Report) parseAccounting() error {
	if err := r.expectReceived(); err != nil {
		return fmt.Errorf("received: %w", err)
	} else if err = r.expectComma(); err != nil {
		return fmt.Errorf("cost: %w", err)
	} else if err = r.expectCost(); err != nil {
		return fmt.Errorf("cost: %w", err)
	} else if err = r.expectCredit(); err != nil {
		return fmt.Errorf("credit: %w", err)
	}
	return nil
}

func (r *Report) parseClanStatus() error {
	if err := r.expectClanId(); err != nil {
		return fmt.Errorf("clan id: %w", err)
	} else if err = r.expectComma(); err != nil {
		return fmt.Errorf("clan status: %w", err)
	} else if err = r.expectComma(); err != nil {
		return fmt.Errorf("clan status: %w", err)
	} else if err = r.parseCurrentHex(); err != nil {
		return fmt.Errorf("current hex: %w", err)
	} else if err = r.expectComma(); err != nil {
		return fmt.Errorf("clan status: %w", err)
	} else if err = r.parsePreviousHex(); err != nil {
		return fmt.Errorf("previous hex: %w", err)
	}
	return nil
}

func (r *Report) parseCurrentHex() error {
	if err := r.expectCurrentHex(); err != nil {
		return err
	}
	return nil
}

func (r *Report) parseCurrentTurn() error {
	if err := r.expectCurrentTurn(); err != nil {
		return fmt.Errorf("turn: %w", err)
	} else if err = r.expectCurrentMonth(); err != nil {
		return fmt.Errorf("month: %w", err)
	} else if err = r.expectComma(); err != nil {
		return fmt.Errorf("turn: %w", err)
	} else if err = r.expectCurrentSeason(); err != nil {
		return fmt.Errorf("season: %w", err)
	} else if err = r.expectComma(); err != nil {
		return fmt.Errorf("turn: %w", err)
	} else if err = r.expectCurrentWeather(); err != nil {
		return fmt.Errorf("weather: %w", err)
	}
	return nil
}

func (r *Report) parseDesiredCommodities() error {
	if err := r.expectDesiredCommodities(); err != nil {
		return fmt.Errorf("desired commodities: %w", err)
	}
	return nil
}

// parseGMNotes attempts to pluck the free-form notes from the report.
func (r *Report) parseGMNotes() error {
	// list of literals that end the notes section.
	literals := []string{"Scout 1:", "Tribe Activities:", "Final Activities:", "Transfer goods to"}

	saved := r.input
	if token, _ := lexFirstLiteral(r.input, literals...); token != nil {
		log.Printf("gm notes: found %q, skipping notes...\n", string(token))
		r.input = saved
		return nil
	}

	// collect the literals with new-lines
	var nlLiterals []string
	for _, literal := range literals {
		nlLiterals = append(nlLiterals, "\n"+literal)
	}

	notes, rest, _ := runToFirstLiteral(r.input, nlLiterals...)
	if notes == nil {
		panic("assert(gmNotes are terminated)")
	}
	r.GMNotes = string(notes)
	r.input = rest
	return nil
}

func (r *Report) parseGoodsTribe() error {
	if err := r.expectGoodsTribe(); err != nil {
		return fmt.Errorf("goods tribe: %w", err)
	}
	return nil
}

func (r *Report) parseMoveInto(input []byte) (direction, terrain, rest []byte, err error) {
	pfx, rest := lexLiteral(input, []byte("NOT ENOUGH M.P'S TO MOVE TO "))
	if pfx == nil {
		return nil, nil, input, nil
	}
	direction, rest = lexDirection(rest)
	if direction == nil {
		return nil, nil, input, fmt.Errorf("move: missing direction: %q", input)
	}
	var into []byte
	into, rest = lexLiteral(rest, []byte("INTO "))
	if into == nil {
		return nil, nil, input, fmt.Errorf("move: missing into: %q", input)
	}
	terrain, rest = lexTerrain(rest)
	if terrain == nil {
		return nil, nil, input, fmt.Errorf("move: missing terrain: %q", input)
	}
	return direction, terrain, rest, nil
}

func (r *Report) parseNextTurn() error {
	if err := r.expectNextTurn(); err != nil {
		return fmt.Errorf("turn: %w", err)
	} else if err = r.expectNextMonth(); err != nil {
		return fmt.Errorf("month: %w", err)
	} else if err = r.expectComma(); err != nil {
		return fmt.Errorf("turn: %w", err)
	} else if err = r.expectTurnDate(); err != nil {
		return fmt.Errorf("due: %w", err)
	}
	return nil
}

func (r *Report) parseOceanBlock(input []byte) (direction, rest []byte, err error) {
	pfx, rest := lexLiteral(input, []byte("CAN'T MOVE ON OCEAN TO "))
	if pfx == nil {
		return nil, input, nil
	}
	direction, rest = lexDirection(rest)
	if direction == nil {
		return nil, input, fmt.Errorf("ocean: missing direction: %q", input)
	}
	var ofHex []byte
	ofHex, rest = lexLiteral(rest, []byte("OF HEX"))
	if ofHex == nil {
		return nil, input, fmt.Errorf("ocean: missing of hex: %q", input)
	}
	return direction, rest, nil
}

func (r *Report) parseOptionalScoutReports() error {
	log.Printf("parseOptionalScoutReports\n")
	if !r.peekLiteral("Scout 1:") {
		return nil
	}
	if err := r.parseScoutReports(); err != nil {
		return fmt.Errorf("scouting: %w", err)
	}
	return nil
}

func (r *Report) parsePreviousHex() error {
	if err := r.expectOpenParen(); err != nil {
		return err
	} else if err = r.expectPreviousHex(); err != nil {
		return err
	} else if err = r.expectCloseParen(); err != nil {
		return err
	}
	return nil
}

func (r *Report) parseRiverBlock(input []byte) (direction, rest []byte, err error) {
	pfx, rest := lexLiteral(input, []byte("NO FORD ON RIVER TO "))
	if pfx == nil {
		return nil, input, nil
	}
	direction, rest = lexDirection(rest)
	if direction == nil {
		return nil, input, fmt.Errorf("river: missing direction: %q", input)
	}
	var ofHex []byte
	ofHex, rest = lexLiteral(rest, []byte("OF HEX"))
	if ofHex == nil {
		return nil, input, fmt.Errorf("river: missing of hex: %q", input)
	}
	return direction, rest, nil
}

// parseScoutAction:
//
//	action :== result (BackSlash result)*
//
//	result :== unitId
//	  | OceanBlocked Direction OfHex (Comma finding)*
//	  | RiverBlocked Direction OfHex (Comma finding)*
//	  | MPExhausted Direction Terrain (Comma finding)*
//	  | Direction Dash Terrain (Comma Terrain Direction (Comma Direction)*)* (Comma finding)*
//
//	finding :== NothingOfInterest
//	  | Find ore
//	  | Patrolled unitId+
func (r *Report) parseScoutAction(no int, input []byte) error {
	log.Printf("parseScoutAction %d: %q\n", no, input)
	if len(input) == 0 {
		panic("assert(scout has action)")
	}
	var err error
	input = bytes.ToUpper(append([]byte{}, input...))
	for _, step := range bytes.Split(input, []byte{'\\'}) {
		var mr *MovementResult
		var direction, terrain, rest []byte
		if direction, rest = lexDirection(step); direction != nil { // movement
			log.Printf("parseScoutAction %d: move %q\n", no, direction)
			// expect terrain
			terrain, rest = lexTerrain(rest)
			if terrain == nil {
				return fmt.Errorf("move: unknown terrain %q", step)
			}
			log.Printf("parseScoutAction %d: move %q into %q\n", no, direction, terrain)
			mr = &MovementResult{Direction: string(direction), Succeeded: true, Terrain: string(terrain)}
		} else if direction, terrain, rest, err = r.parseMoveInto(step); err != nil {
			return fmt.Errorf("action: move: %w", err)
		} else if direction != nil { // failed movement
			log.Printf("parseScoutAction %d: move %q into %q *** failed\n", no, direction, terrain)
			mr = &MovementResult{Direction: string(direction), Terrain: string(terrain)}
		} else if direction, rest, err = r.parseOceanBlock(step); err != nil {
			return fmt.Errorf("action: move: %w", err)
		} else if direction != nil { // failed movement
			mr = &MovementResult{Direction: string(direction), Terrain: string(terrain)}
		} else if direction, rest, err = r.parseRiverBlock(step); err != nil {
			return fmt.Errorf("action: move: %w", err)
		} else if direction != nil { // failed movement
			mr = &MovementResult{Direction: string(direction), BlockedBy: "R"}
		} else {
			return fmt.Errorf("action: step: unknown option %q", step)
		}
		var comma []byte
		if comma, rest = lexDelimiter(rest, ','); comma != nil {
			// something else
			if bytes.HasPrefix(rest, []byte{'O', ' '}) { // ocean
				rest = rest[2:]
			} else if bytes.HasPrefix(rest, []byte{'R', 'i', 'v', 'e', 'r', ' '}) { // river
				rest = rest[6:]
			}
		}

		log.Printf("scout, move %v, rest is now %q\n", mr.Succeeded, rest)
	}
	return nil
}

func (r *Report) parseScoutReport(no int) error {
	log.Printf("parseScoutReport %d\n", no)
	scoutId := fmt.Sprintf("Scout %d:", no)
	if !r.acceptLiteral(scoutId) {
		return fmt.Errorf("expected %q: got %q", scoutId, r.nextToken())
	}
	if r.acceptLiteral("Scout ") { // action is "Scout"
		action, rest := runToEOL(r.input)
		if err := r.parseScoutAction(no, action); err != nil {
			return fmt.Errorf("action: %w", err)
		}
		r.input = rest
		return nil
	}
	return fmt.Errorf("unknown action: %q", r.nextToken())
}

func (r *Report) parseScoutReports() error {
	log.Printf("parseScoutReports\n")
	for scoutNo := 1; scoutNo <= 8; scoutNo++ {
		if !r.peekLiteral(fmt.Sprintf("Scout %d:", scoutNo)) {
			break
		}
		if err := r.parseScoutReport(scoutNo); err != nil {
			return fmt.Errorf("scout %d: %w", scoutNo, err)
		}
	}
	r.stopScouting()
	return nil
}

func (r *Report) parseTurnDetails() error {
	if err := r.parseCurrentTurn(); err != nil {
		return fmt.Errorf("current turn: %w", err)
	} else if err = r.parseNextTurn(); err != nil {
		return fmt.Errorf("next turn: %w", err)
	}
	return nil
}

// recover is the handler that turns panics into returns from the top level of Parse.
func (r *Report) recover(errp *error) {
	if e := recover(); e != nil {
		if _, ok := e.(runtime.Error); ok {
			panic(e)
		}
		if r != nil {
			r.stopParse()
		}
		*errp = e.(error)
	}
}

func (r *Report) stopParse() {
	r.input, r.tribe, r.scout = nil, nil, nil
}

func (r *Report) stopScouting() {
	if r.scout != nil {
		r.scout, r.scout.mr = nil, nil
	}
}

func (r *Report) nextToken() string {
	return lexToken(r.input)
}
