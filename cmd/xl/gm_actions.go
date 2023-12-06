// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
	"strings"
)

func (w *workbook) loadGMActions(version string) func() error {
	switch version {
	case "1.11":
		return w.loadGMActionsV1
	}
	panic(fmt.Sprintf("assert(version != %q)", version))
}

func (w *workbook) loadGMActionsV1() error {
	sheet := "GM Actions"

	// fetch the contents of the worksheet
	rows, err := w.f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("getRows: %w", err)
	}

	columns := []string{"Unit", "What does the GM need to do?"}
	if len(rows) == 0 {
		return fmt.Errorf("headers: missing header row")
	} else if err := w.checkHeaders(sheet, rows[0], columns...); err != nil {
		return fmt.Errorf("headers: %w", err)
	}

	expectedCells := 2
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
			// ignore blank comments
			continue
		} else if len(row) > expectedCells {
			log.Printf("%s: %s: row %2d: cells %3d\n", w.Name, sheet, no, len(row))
			return fmt.Errorf("unknown cells")
		}
		id, err := cellToString(row, 0)
		if err != nil {
			return fmt.Errorf("unit: %w", err)
		}
		unit, ok := w.Units[id]
		if !ok {
			log.Printf("%s: %s: row %2d: unknown unit %q\n", w.Name, sheet, no, id)
			unit = NewUnit(id)
			w.Units[unit.Unit] = unit
		}
		request, err := cellToString(row, 1)
		if err != nil {
			return fmt.Errorf("request: %w", err)
		}
		request = strings.TrimSpace(request)
		if request == "" {
			// ignore blank rows
			continue
		}
		unit.GMRequests = append(unit.GMRequests, request)
	}

	return nil
}
