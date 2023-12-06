// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
	"strings"
)

func (w *workbook) loadScoutMovement(version string) func() error {
	switch version {
	case "1.11":
		return w.loadScoutMovementV1
	}
	panic(fmt.Sprintf("assert(version != %q)", version))
}

func (w *workbook) loadScoutMovementV1() error {
	sheet := "Scout_Movement"

	columns := []string{"TRIBE", "No_of_Scouts", "No_of_Horses", "No_of_Elephants", "No_of_Camels", "Mission"}
	for i := 1; i <= 9; i++ {
		columns = append(columns, fmt.Sprintf("Movement%d", i))
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
		if len(row) == 1 {
			// the setup workbook usually contains an extra row with just the tribe
			log.Printf("%s: %s: row %2d: cells %3d\n", w.Name, sheet, no, len(row))
			break
		} else if len(row) < expectedCells {
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

		scout := &Scout{}
		if scout.Scouts, err = cellToInt(row, 1); err != nil {
			return fmt.Errorf("scouts: %w", err)
		}
		if scout.Horses, err = cellToInt(row, 2); err != nil {
			return fmt.Errorf("horses: %w", err)
		}
		if scout.Elephants, err = cellToInt(row, 3); err != nil {
			return fmt.Errorf("elephants: %w", err)
		}
		if scout.Camels, err = cellToInt(row, 4); err != nil {
			return fmt.Errorf("camels: %w", err)
		}
		if scout.Mission, err = cellToString(row, 5); err != nil {
			return fmt.Errorf("mission: %w", err)
		}

		toLimit := false
		for n := 6; n < len(columns)-1; n++ {
			dir, err := cellToString(row, n)
			if err != nil {
				return fmt.Errorf("movement%d: %w", n-2, err)
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
			case "N", "NE", "SE", "S", "SW", "NW":
				scout.Moves = append(scout.Moves, &ScoutMove{Direction: dir})
			case "NL", "NEL", "SEL", "SL", "SWL", "NWL":
				scout.Moves = append(scout.Moves, &ScoutMove{Direction: dir, ToLimit: true})
				toLimit = true
			case "FOR", "FOL":
				scout.Moves = append(scout.Moves, &ScoutMove{Direction: dir})
				toLimit = true
			case "FRR", "FRL":
				scout.Moves = append(scout.Moves, &ScoutMove{Direction: dir})
				toLimit = true
			default:
				log.Printf("%s: %s: row %2d: invalid move: %d: %q\n", w.Name, sheet, no, n, dir)
				return fmt.Errorf("movement%d: %w", n-2, fmt.Errorf("invalid"))
			}
		}

		if scout.Processed, err = cellToBool(row, len(columns)-1); err != nil {
			return fmt.Errorf("processed: %w", err)
		}

		unit.Scouts = append(unit.Scouts, scout)
	}

	return nil
}
