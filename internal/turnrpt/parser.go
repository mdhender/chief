// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"bytes"
	"fmt"
	"github.com/mdhender/chief/internal/docconv"
	"github.com/mdhender/chief/internal/terrain"
	"log"
	"os"
	"strconv"
)

type parser struct {
	scanner *scanner
	input   []byte
	pos     int
	meta    map[string]string
	report  []*Report
	err     error
}

func ParseDocument(filename string) ([]*TribeSection, error) {
	log.Printf("parse: filename %s\n", filename)
	fp, err := os.Open(filename)
	if err != nil {
		log.Printf("[ptr] open %v\n", err)
		return nil, err
	}
	body, meta, err := docconv.ConvertDocx(fp)
	if err != nil {
		log.Printf("[ptr] convert %v\n", err)
		return nil, err
	}
	return Parse([]byte(body), meta)
}

func Parse(input []byte, _ map[string]string) ([]*TribeSection, error) {
	var sections []*TribeSection
	loops := 0
	for len(input) != 0 && loops < 8 {
		loops++
		section, rest, err := parseTribeSection(skipws(input))
		if section != nil {
			sections = append(sections, section)
		}
		if err != nil {
			return sections, fmt.Errorf("parse: %w", err)
		}
		input = skipws(rest)
	}

	return sections, nil
}

func shinp(pfx string, input []byte) {
	poof := string(input)
	if len(poof) > 25 {
		poof = poof[:25]
	}
	log.Printf("[%s] input %q\n", pfx, poof)
}

// Received: $0.07, Cost: $ 1.25  Credit: $ 8.25
func parseAccounts(input []byte) (*Account, []byte, error) {
	shinp("accounts", input)
	run, rest := literal([]byte("Received: "), skipws(input))
	if run == nil {
		return nil, input, nil
	}
	account := &Account{}
	if run, rest = currency(skipws(rest)); run == nil {
		return account, input, fmt.Errorf("accounts: received: expected amount")
	}
	account.Received = string(run)
	if run, rest = literal([]byte{','}, skipws(rest)); run == nil {
		return account, rest, fmt.Errorf("accounts: expected ','")
	}
	if run, rest = literal([]byte("Cost: "), skipws(rest)); run == nil {
		return account, input, fmt.Errorf("accounts: expected cost")
	}
	if run, rest = currency(skipws(rest)); run == nil {
		return account, input, fmt.Errorf("accounts: cost: expected amount")
	}
	account.Cost = string(run)
	if run, rest = literal([]byte("Credit: "), skipws(rest)); run == nil {
		return account, input, fmt.Errorf("accounts: expected credit")
	}
	if run, rest = currency(skipws(rest)); run == nil {
		return account, input, fmt.Errorf("accounts: credit: expected amount")
	}
	account.Credit = string(run)
	return account, rest, nil
}

// parses "Current Hex = ## NNNN"
func parseCurrentHex(input []byte) (*Hex, []byte, error) {
	shinp("curr hex", input)
	run, rest := literal([]byte("Current Hex = "), skipws(input))
	if run == nil {
		return nil, input, nil
	}
	hex, rest, err := parseHexNo(skipws(rest))
	if err != nil {
		return nil, rest, fmt.Errorf("current hex: %w", err)
	} else if hex == nil {
		return nil, rest, fmt.Errorf("current hex: missing hex")
	}
	return hex, rest, nil
}

// Current Turn 900-01 (#1), Spring, FINE
func parseCurrentTurn(input []byte) (*Turn, []byte, error) {
	shinp("curr turn", input)
	run, rest := literal([]byte("Current Turn "), skipws(input))
	if run == nil {
		return nil, input, nil
	}
	yearMonth, rest, err := parseYearDashMonth(skipws(rest))
	if err != nil {
		return nil, rest, fmt.Errorf("current turn: %w", err)
	} else if yearMonth == nil {
		return nil, rest, fmt.Errorf("current turn: expected year-month")
	}
	turn := &Turn{Year: yearMonth.Year, Month: yearMonth.Month}
	run, rest, err = parseGameTurn(skipws(rest))
	if err != nil {
		return nil, rest, fmt.Errorf("current turn: %w", err)
	} else if run == nil {
		return turn, rest, fmt.Errorf("current turn: expected game turn")
	}
	turn.No, _ = strconv.Atoi(string(run))
	if run, rest = literal([]byte{','}, rest); run == nil {
		return turn, rest, fmt.Errorf("current turn: expected ','")
	}
	if run, rest = season(skipws(rest)); run == nil {
		return turn, rest, fmt.Errorf("current turn: expected season")
	}
	turn.Season = string(run)
	if run, rest = literal([]byte{','}, rest); run == nil {
		return turn, rest, fmt.Errorf("current turn: expected ','")
	}
	run, rest = weather(skipws(rest))
	if run == nil {
		return turn, rest, fmt.Errorf("current turn: expected weather")
	}
	turn.Weather = string(run)
	return turn, rest, nil
}

// Desired Commodities: No com....
func parseDesiredCommodities(input []byte) (*DesiredCommodities, []byte, error) {
	shinp("desired commodities", input)
	run, rest := literal([]byte("Desired Commodities: "), skipws(input))
	if run == nil {
		log.Printf("dc not found\n")
		return nil, input, nil
	}
	run, rest = literal([]byte("No commodities allocated"), skipws(rest))
	if run == nil {
		return nil, rest, fmt.Errorf("desired commodities: expected commodities")
	}
	return &DesiredCommodities{}, rest, nil
}

// Game turn looks like "(#N)"
func parseGameTurn(input []byte) ([]byte, []byte, error) {
	shinp("game turn", input)
	run, rest := literal([]byte{'(', '#'}, skipws(input))
	if run == nil {
		return nil, input, nil
	}
	run, rest = integer(rest)
	if run == nil {
		return nil, input, nil
	}
	turn := run
	run, rest = literal([]byte{')'}, skipws(rest))
	if run == nil {
		return nil, rest, fmt.Errorf("game turn: expected ')'")
	}
	return turn, rest, nil
}

// Goods Tribe: No GT
func parseGoodsTribe(input []byte) ([]byte, []byte, error) {
	run, rest := literal([]byte("Goods Tribe: "), skipws(input))
	if run == nil {
		return nil, input, nil
	}
	run, rest = toeol(skipws(rest))
	return bytes.TrimSpace(run), rest, nil
}

// parses "## CCRR"
func parseHexNo(input []byte) (*Hex, []byte, error) {
	shinp("hex no", input)
	run, rest := hexNo(skipws(input))
	if run == nil {
		return nil, input, nil
	}
	hex := &Hex{
		GridX: string(run[0]),
		GridY: string(run[1]),
		Col:   10*int(run[3]-'0') + int(run[4]-'0'),
		Row:   10*int(run[5]-'0') + int(run[6]-'0'),
	}
	return hex, rest, nil
}

// Next Turn 900-02 (#2), 12/11/2023
func parseNextTurn(input []byte) (*Turn, []byte, error) {
	shinp("next turn", input)
	run, rest := literal([]byte("Next Turn "), skipws(input))
	if run == nil {
		return nil, input, nil
	}
	yearMonth, rest, err := parseYearDashMonth(skipws(rest))
	if err != nil {
		return nil, rest, fmt.Errorf("next turn: %w", err)
	} else if yearMonth == nil {
		return nil, rest, fmt.Errorf("next turn: expected year-month")
	}
	turn := &Turn{Year: yearMonth.Year, Month: yearMonth.Month}
	run, rest, err = parseGameTurn(skipws(rest))
	if err != nil {
		return nil, rest, fmt.Errorf("next turn: %w", err)
	} else if run == nil {
		return turn, rest, fmt.Errorf("next turn: expected game turn")
	}
	turn.No, _ = strconv.Atoi(string(run))
	if run, rest = literal([]byte{','}, rest); run == nil {
		return turn, rest, fmt.Errorf("next turn: expected ','")
	}
	run, rest = ddMmYyyy(skipws(rest))
	if run == nil {
		return turn, rest, fmt.Errorf("next turn: expected date")
	}
	turn.Date = fmt.Sprintf("%s/%s/%s", string(run[6:10]), string(run[3:5]), string(run[:2]))
	return turn, rest, nil
}

// parses "(Previous Hex = ## NNNN)"
func parsePreviousHex(input []byte) (*Hex, []byte, error) {
	shinp("prev hex", input)
	run, rest := literal([]byte("(Previous Hex = "), skipws(input))
	if run == nil {
		return nil, input, nil
	}
	hex, rest, err := parseHexNo(skipws(rest))
	if err != nil {
		return nil, rest, fmt.Errorf("previous hex: %w", err)
	} else if hex == nil {
		return nil, rest, fmt.Errorf("previous hex: missing hex")
	}
	run, rest = literal([]byte{')'}, skipws(rest))
	if run == nil {
		return hex, rest, fmt.Errorf("previous hex: missing ')'")
	}
	return hex, rest, nil
}

// Examples:
// Scout SW-RH\ S-GH\  Not enough M.P's to move to S into CONIFER HILLS, Nothing of interest found
// Scout S-PR\ S-PR,  O S,River SE\ Can't Move on Ocean to S of HEX, Nothing of interest found
// Scout SE-GH\ S-PR, River SE S\ No Ford on River to S of HEX, Nothing of interest found
// Scout NW-GH\ SW-GH\ S-PR\  Not enough M.P's to move to S into CONIFER HILLS, Nothing of interest found
// Scout NE-GH\ SE-PR, River SE\S-PR, River NE SE S\ No Ford on River to S of HEX, Nothing of interest found
// Scout NE-GH\ SE-PR, River SE\NE-PR, River S\ Not enough M.P's to move to NW into ROCKY HILLS, Nothing of interest found
// Scout N-GH\ NE-PR\ NW-PR\ SE-PR\  Not enough M.P's to move to N into ROCKY HILLS, Nothing of interest found
// Scout NW-GH\ SW-GH\ NW-PR\  Not enough M.P's to move to NE into GRASSY HILLS, Nothing of interest found
func parseScoutingResult(id int, input []byte) (*ScoutingResult, []byte, error) {
	shinp("scout result", input)
	run, rest := literal([]byte("Scout "), skipws(input))
	if run == nil {
		return nil, input, nil
	}
	result := &ScoutingResult{Id: fmt.Sprintf("Scout %d", id)}
	moves := bytes.Split(bytes.ToUpper(bytes.TrimSpace(rest)), []byte{'\\'})
	for no, move := range moves {
		move = bytes.TrimSpace(move)
		log.Printf("Scout %3d: move     %4d: %q\n", id, no+1, string(move))
		fields := bytes.Split(move, []byte{','})
		var mr MovementResult
		for no, field := range fields {
			field = bytes.TrimSpace(field)
			log.Printf("Scout %3d: ...field %4d: %q\n", id, no+1, string(field))
			if no == 0 { // extract direction
				mr.Direction, field = directionSuccess(field)
				if mr.Direction == "" {
					mr.Failed = true
					if bytes.HasPrefix(field, []byte("CAN'T MOVE ON OCEAN TO ")) {
						// can't move on OCEAN to DIRECTION of hex
						mr.Direction, field = directionFailed(field[23:])
					} else if bytes.HasPrefix(field, []byte("NO FORD ON RIVER TO ")) {
						// no ford on RIVER to DIRECTION of hex
						mr.Direction, field = directionFailed(field[20:])
					} else if bytes.HasPrefix(field, []byte("NOT ENOUGH M.P'S TO MOVE TO ")) {
						// not enough M.P.'s to move DIRECTION into TERRAIN
						mr.Direction, field = directionFailed(field[28:])
					} else { // ... movement failed
						panic(fmt.Sprintf("assert(scouting direction != %q)", string(field)))
					}
				}
			} else { // extract findings
				if bytes.HasPrefix(field, []byte("O ")) {
					// OCEAN DIRECTION...
				} else if bytes.HasPrefix(field, []byte("RIVER ")) {
					// RIVER DIRECTION...
				} else if bytes.Equal(field, []byte("NOTHING OF INTEREST FOUND")) {
					mr.Found = append(mr.Found, "NOTHING")
				} else {
					panic(fmt.Sprintf("assert(scouting found != %q)", string(field)))
				}
			}
		}
		result.MovementResults = append(result.MovementResults, mr)
	}
	return result, nil, nil
}

func parseScoutingSection(input []byte) ([]*ScoutingResult, []byte, error) {
	id := 1
	var results []*ScoutingResult
	for {
		shinp(fmt.Sprintf("scouting %d", id), input)
		run, rest := literal([]byte(fmt.Sprintf("Scout %d:", id)), skipws(input))
		if run == nil {
			break
		}
		run, rest = toeol(rest)
		if result, junk, err := parseScoutingResult(id, run); err != nil {
			return results, rest, fmt.Errorf("scout %d: %w", id, err)
		} else if len(junk) != 0 {
			return results, rest, fmt.Errorf("scout %d: unknown input %q", id, string(junk))
		} else if result != nil {
			results = append(results, result)
		}
		input = skipws(rest)
		id = id + 1
	}
	return results, input, nil
}

func parseStatus(statusLiteral string, input []byte) (*Status, []byte, error) {
	run, rest := literal([]byte(statusLiteral), skipws(input))
	if run == nil {
		return nil, input, nil
	}
	status := &Status{}
	var err error
	status.Terrain, rest, err = parseTerrain(skipws(rest))
	if err != nil {
		return status, rest, fmt.Errorf("status: terrain: %w", err)
	} else if status.Terrain == terrain.NONE {
		return status, rest, fmt.Errorf("status: expected terrain")
	}
	return status, rest, fmt.Errorf("unexpected eof")
}

func parseTerrain(input []byte) (terrain.CODE, []byte, error) {
	run, rest := word(skipws(input))
	if run == nil {
		return terrain.NONE, input, nil
	}
	if bytes.Equal(run, []byte("ALPS")) {
		return terrain.ALPS, rest, nil
	} else if bytes.Equal(run, []byte("AR")) {
		return terrain.AR, rest, nil
	} else if bytes.Equal(run, []byte("BH")) {
		return terrain.BH, rest, nil
	} else if bytes.Equal(run, []byte("BR")) {
		return terrain.BR, rest, nil
	} else if bytes.Equal(run, []byte("CH")) || bytes.Equal(run, []byte("CONIFER HILLS")) {
		return terrain.CH, rest, nil
	} else if bytes.Equal(run, []byte("DE")) {
		return terrain.DE, rest, nil
	} else if bytes.Equal(run, []byte("DF")) {
		return terrain.DF, rest, nil
	} else if bytes.Equal(run, []byte("DH")) {
		return terrain.DH, rest, nil
	} else if bytes.Equal(run, []byte("Fords")) {
		return terrain.Fords, rest, nil
	} else if bytes.Equal(run, []byte("GH")) {
		return terrain.GH, rest, nil
	} else if bytes.Equal(run, []byte("HSM")) {
		return terrain.HSM, rest, nil
	} else if bytes.Equal(run, []byte("JG")) {
		return terrain.JG, rest, nil
	} else if bytes.Equal(run, []byte("JH")) {
		return terrain.JH, rest, nil
	} else if bytes.Equal(run, []byte("L")) {
		return terrain.L, rest, nil
	} else if bytes.Equal(run, []byte("LCM")) {
		return terrain.LCM, rest, nil
	} else if bytes.Equal(run, []byte("LJM")) {
		return terrain.LJM, rest, nil
	} else if bytes.Equal(run, []byte("LSM")) {
		return terrain.LSM, rest, nil
	} else if bytes.Equal(run, []byte("PI")) {
		return terrain.PI, rest, nil
	} else if bytes.Equal(run, []byte("PR")) {
		return terrain.PR, rest, nil
	} else if bytes.Equal(run, []byte("R")) {
		return terrain.R, rest, nil
	} else if bytes.Equal(run, []byte("RH")) {
		return terrain.RH, rest, nil
	} else if bytes.Equal(run, []byte("SH")) {
		return terrain.SH, rest, nil
	} else if bytes.Equal(run, []byte("SW")) {
		return terrain.SW, rest, nil
	} else if bytes.Equal(run, []byte("TU")) {
		return terrain.TU, rest, nil
	}
	return terrain.NONE, input, nil
}

func parseTribeId(input []byte) (string, []byte, error) {
	run, rest := keyword([]byte("Tribe"), skipws(input))
	if run == nil {
		return "", input, nil
	}
	run, rest = tribeId(skipws(rest))
	if run == nil {
		return "", rest, fmt.Errorf("missing tribe id")
	}
	return string(run), rest, nil
}

// Tribe 0138, , Current Hex = ## 0102, (Previous Hex = ## 0201)
func parseTribeSection(input []byte) (*TribeSection, []byte, error) {
	shinp("tribe section", input)
	tribeId, rest, err := parseTribeId(skipws(input))
	if err != nil {
		return nil, rest, fmt.Errorf("tribe section: %w", err)
	} else if tribeId == "" {
		return nil, input, nil
	}
	t := &TribeSection{
		Id: tribeId,
	}
	run, rest := literal([]byte{','}, rest)
	if run == nil {
		return t, rest, fmt.Errorf("tribe: tribe: expected ','")
	}

	// unknown field
	run, rest = literal([]byte{','}, skipws(rest))
	if run == nil {
		return t, rest, fmt.Errorf("tribe: unknown field: expected ','")
	}

	// current hex
	t.CurrHex, rest, err = parseCurrentHex(skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe: %w", err)
	}
	run, rest = literal([]byte{','}, skipws(rest))
	if run == nil {
		return t, rest, fmt.Errorf("tribe: current hex: expected ','")
	}

	// previous hex
	t.PrevHex, rest, err = parsePreviousHex(skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe section: %w", err)
	}

	// Current Turn 900-01 (#1), Spring, FINE
	t.Current, rest, err = parseCurrentTurn(skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe section: %w", err)
	} else if t.Current == nil {
		return t, rest, fmt.Errorf("tribe section: expected current turn")
	}

	// Next Turn 900-01 (#1), 29/10/2023
	t.Next, rest, err = parseNextTurn(skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe section: %w", err)
	} else if t.Next == nil {
		return t, rest, fmt.Errorf("tribe section: expected next turn")
	}

	// Received: $0.07, Cost: $ 1.25  Credit: $ 8.25
	t.Accounts, rest, err = parseAccounts(skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe section: %w", err)
	} else if t.Accounts == nil {
		return t, rest, fmt.Errorf("tribe section: expected accounting")
	}

	// Goods Tribe: No GT
	run, rest, err = parseGoodsTribe(skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe section: %w", err)
	} else if len(run) == 0 {
		return t, rest, fmt.Errorf("tribe section: expected goods tribe")
	}
	t.GoodsTribe = string(run)

	// Desired Commodities
	t.DesiredCommodities, rest, err = parseDesiredCommodities(skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe section: desired commodities: %w", err)
	} else if t.DesiredCommodities == nil {
		return t, rest, fmt.Errorf("tribe section: expected desired commodities")
	}

	// Scouting results
	t.Scouting.Results, rest, err = parseScoutingSection(skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe section: scouting: %w", err)
	}

	// int used
	// int paragraph
	// Tribe Activities:
	// Final Activities:
	//   Transfer goods to
	//   Tribe Movement to

	// tribe status
	statusLiteral := fmt.Sprintf("%s Status:", t.Id)
	t.Status, rest, err = parseStatus(statusLiteral, skipws(rest))
	if err != nil {
		return t, rest, fmt.Errorf("tribe section: status: %w", err)
	} else if t.Status == nil {
		return t, rest, fmt.Errorf("tribe section: expected status")
	}

	return t, rest, nil
}

func parseYearDashMonth(input []byte) (*YearMonth, []byte, error) {
	shinp("year-mo", input)
	run, rest := turnNo(skipws(input))
	if run == nil {
		return nil, input, nil
	}
	year, _ := strconv.Atoi(string(run[:3]))
	month, _ := strconv.Atoi(string(run[4:]))
	return &YearMonth{Year: year, Month: month}, rest, nil
}

func directionFailed(input []byte) (string, []byte) {
	if bytes.HasPrefix(input, []byte{'N', ' '}) {
		return "N", input[2:]
	} else if bytes.HasPrefix(input, []byte{'N', 'E', ' '}) {
		return "NE", input[3:]
	} else if bytes.HasPrefix(input, []byte{'S', 'E', ' '}) {
		return "SE", input[3:]
	} else if bytes.HasPrefix(input, []byte{'S', ' '}) {
		return "S", input[2:]
	} else if bytes.HasPrefix(input, []byte{'S', 'W', ' '}) {
		return "SW", input[3:]
	} else if bytes.HasPrefix(input, []byte{'N', 'W', ' '}) {
		return "NW", input[3:]
	}
	return "", input
}

func directionSuccess(input []byte) (string, []byte) {
	if bytes.HasPrefix(input, []byte{'N', '-'}) {
		return "N", input[2:]
	} else if bytes.HasPrefix(input, []byte{'N', 'E', '-'}) {
		return "NE", input[3:]
	} else if bytes.HasPrefix(input, []byte{'S', 'E', '-'}) {
		return "SE", input[3:]
	} else if bytes.HasPrefix(input, []byte{'S', '-'}) {
		return "S", input[2:]
	} else if bytes.HasPrefix(input, []byte{'S', 'W', '-'}) {
		return "SW", input[3:]
	} else if bytes.HasPrefix(input, []byte{'N', 'W', '-'}) {
		return "NW", input[3:]
	}
	return "", input
}
