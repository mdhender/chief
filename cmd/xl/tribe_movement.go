// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
	"strings"
)

func (w *workbook) loadTribeMovement(version string) func() error {
	switch version {
	case "1.11":
		return w.loadTribeMovementV1
	}
	panic(fmt.Sprintf("assert(version != %q)", version))
}

func (w *workbook) loadTribeMovementV1() error {
	sheet := "Tribe_Movement"

	columns := []string{"TRIBE", "FOLLOW_TRIBE", "Hex"}
	for i := 1; i <= 40; i++ {
		columns = append(columns, fmt.Sprintf("MOVEMENT_%d", i))
	}
	columns = append(columns, "Processed")

	// fetch the contents of the worksheet
	rows, err := w.f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("getRows: %w", err)
	}

	if len(rows) == 0 {
		return fmt.Errorf("headers: missing header row")
	} else if err := w.checkHeaders(sheet, rows[0], columns...); err != nil {
		return fmt.Errorf("headers: %w", err)
	}

	expectedCells := len(columns)
	for n, row := range rows {
		if n == 0 {
			// header row was checked above
			continue
		} else if len(row) == 0 || row[0] == "" {
			// worksheet contains trailing junk
			break
		}
		no := n + 1
		if len(row) < expectedCells {
			log.Printf("%s: %s: row %2d: cells %3d\n", w.Name, sheet, no, len(row))
			return fmt.Errorf("missing cells")
		} else if len(row) > expectedCells {
			log.Printf("%s: %s: row %2d: cells %3d\n", w.Name, sheet, no, len(row))
			return fmt.Errorf("unknown cells")
		}

		id, err := cellToString(row, 0)
		if err != nil {
			return fmt.Errorf("tribe: %w", err)
		}
		unit, ok := w.Units[id]
		if !ok {
			log.Printf("%s: %s: row %2d: unknown unit %q\n", w.Name, sheet, no, id)
			unit = NewUnit(id)
			w.Units[id] = unit
		}

		unit.Movement = &Movement{}
		if unit.Movement.Follow, err = cellToString(row, 1); err != nil {
			return fmt.Errorf("follow-tribe: %w", err)
		} else if unit.Movement.Follow == "" {
			// ignore this field
		} else if _, ok := w.Units[unit.Movement.Follow]; !ok {
			log.Printf("%s: %s: row %2d: unknown follow unit %q\n", w.Name, sheet, no, unit.Movement.Follow)
			w.Units[unit.Movement.Follow] = NewUnit(unit.Movement.Follow)
		}

		if unit.Movement.Hex, err = cellToString(row, 2); err != nil {
			return fmt.Errorf("hex: %w", err)
		}

		toLimit := false
		for n := 3; n < len(columns)-1; n++ {
			dir, err := cellToString(row, n)
			if err != nil {
				return fmt.Errorf("movement_%d: %w", n-2, err)
			}
			dir = strings.ToUpper(dir)
			if dir == "EMPTY" {
				continue
			} else if toLimit {
				log.Printf("%s: %s: row %2d: move follows to-limit: %d: %q\n", w.Name, sheet, no, n, dir)
			}
			switch dir {
			case "EMPTY":
				// ignore
			case "FOLLOW":
				if unit.Movement.Follow == "" {
					log.Printf("%s: %s: row %2d: invalid move: %d: %q\n", w.Name, sheet, no, n, dir)
				}
			case "STILL":
				unit.Movement.Moves = append(unit.Movement.Moves, &Move{Still: true})
			case "N", "NE", "SE", "S", "SW", "NW":
				unit.Movement.Moves = append(unit.Movement.Moves, &Move{Direction: dir})
			case "NL", "NEL", "SEL", "SL", "SWL", "NWL":
				unit.Movement.Moves = append(unit.Movement.Moves, &Move{Direction: dir, ToLimit: true})
				toLimit = true
			case "FOR", "FOL":
				unit.Movement.Moves = append(unit.Movement.Moves, &Move{Direction: dir})
				toLimit = true
			case "FRR", "FRL":
				unit.Movement.Moves = append(unit.Movement.Moves, &Move{Direction: dir})
				toLimit = true
			default:
				log.Printf("%s: %s: row %2d: invalid move: %d: %q\n", w.Name, sheet, no, n, dir)
				return fmt.Errorf("movement_%d: %w", n-2, fmt.Errorf("invalid"))
			}
		}

		if unit.Movement.Processed, err = cellToBool(row, len(columns)-1); err != nil {
			return fmt.Errorf("processed: %w", err)
		}
	}

	return nil
}
