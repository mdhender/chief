// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
	"strings"
)

func (w *workbook) loadTransfers(version string) func() error {
	switch version {
	case "1.11":
		return w.loadTransfersV1
	}
	panic(fmt.Sprintf("assert(version != %q)", version))
}

func (w *workbook) loadTransfersV1() error {
	sheet := "Transfers"

	// fetch the contents of the worksheet
	rows, err := w.f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("getRows: %w", err)
	}

	columns := []string{"From", "To", "Item", "Quantity", "Transfer_Timing", "Notes", "Processed"}
	if len(rows) == 0 {
		return fmt.Errorf("headers: missing header row")
	} else if err := w.checkHeaders(sheet, rows[0], columns...); err != nil {
		return fmt.Errorf("headers: %w", err)
	}

	expectedCells := 7
	for n, row := range rows {
		if n == 0 {
			// header row was checked above
			continue
		} else if len(row) == 0 || row[0] == "" {
			// worksheet contains trailing junk
			break
		}
		no := n + 1
		if len(row) < 5 {
			// the last two cells (Notes and Processed) are optional
			log.Printf("%s: %s: row %2d: cells %3d\n", w.Name, sheet, no, len(row))
			return fmt.Errorf("missing cells")
		} else if len(row) > expectedCells {
			log.Printf("%s: %s: row %2d: cells %3d\n", w.Name, sheet, no, len(row))
			return fmt.Errorf("unknown cells")
		}
		t := &Transfer{}
		if t.From, err = cellToString(row, 0); err != nil {
			return fmt.Errorf("from: %w", err)
		} else if t.From, err = cellToString(row, 0); err != nil {
			return fmt.Errorf("from: %w", err)
		} else if t.To, err = cellToString(row, 1); err != nil {
			return fmt.Errorf("to: %w", err)
		} else if t.Item, err = cellToString(row, 2); err != nil {
			return fmt.Errorf("item: %w", err)
		} else if t.Quantity, err = cellToInt(row, 3); err != nil {
			return fmt.Errorf("quantity: %w", err)
		} else if t.TransferTiming, err = cellToString(row, 4); err != nil {
			return fmt.Errorf("transfer-timing: %w", err)
		} else if t.Notes, err = cellToString(row, 5); err != nil {
			return fmt.Errorf("notes: %w", err)
		} else if t.Processed, err = cellToBool(row, 6); err != nil {
			return fmt.Errorf("processed: %w", err)
		}

		if _, ok := w.Units[t.From]; !ok {
			log.Printf("%s: %s: row %2d: unknown unit %q\n", w.Name, sheet, no, t.From)
			w.Units[t.From] = NewUnit(t.From)
		}
		if _, ok := w.Units[t.To]; !ok {
			log.Printf("%s: %s: row %2d: unknown unit %q\n", w.Name, sheet, no, t.To)
			w.Units[t.To] = NewUnit(t.To)
		}
		switch strings.ToUpper(t.TransferTiming) {
		case "AM":
			if w.Transfers == nil {
				w.Transfers = &Transfers{}
			}
			w.Transfers.AfterMovement = append(w.Transfers.AfterMovement, t)
		case "BM":
			if w.Transfers == nil {
				w.Transfers = &Transfers{}
			}
			w.Transfers.BeforeMovement = append(w.Transfers.BeforeMovement, t)
		default:
			return fmt.Errorf("transfer-timing: unknown value %q", t.TransferTiming)
		}
	}

	return nil
}
