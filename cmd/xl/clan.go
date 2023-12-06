// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
)

func (w *workbook) loadClan(version string) func() error {
	switch version {
	case "1.11":
		return w.loadClanV1
	}
	panic(fmt.Sprintf("assert(version != %q)", version))
}

func (w *workbook) loadClanV1() error {
	sheet := "Clan"

	// fetch the contents of the worksheet
	rows, err := w.f.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("getRows: %w", err)
	}

	columns := []string{"Unit", "GT", "Warrior", "Active", "Inactive", "Slave", "Eaters", "Provs", "Months", "Hirelings", "Mercs", "Locals", "Auxiliaries", "Workers", "Used", "Remains", "Cattle", "Dog", "Elephant", "Goat", "Horse", "Camel", "Herders"}
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
		u := &Unit{}
		if u.Unit, err = cellToString(row, 0); err != nil {
			return fmt.Errorf("unit: %w", err)
		} else if u.GT, err = cellToString(row, 1); err != nil {
			return fmt.Errorf("gt: %w", err)
		} else if u.Warrior, err = cellToInt(row, 2); err != nil {
			return fmt.Errorf("warrior: %w", err)
		} else if u.Active, err = cellToInt(row, 3); err != nil {
			return fmt.Errorf("active: %w", err)
		} else if u.Inactive, err = cellToInt(row, 4); err != nil {
			return fmt.Errorf("inactive: %w", err)
		} else if u.Slave, err = cellToInt(row, 5); err != nil {
			return fmt.Errorf("slave: %w", err)
		} else if u.Eaters, err = cellToInt(row, 6); err != nil {
			return fmt.Errorf("eaters: %w", err)
		} else if u.Provs, err = cellToInt(row, 7); err != nil {
			return fmt.Errorf("provs: %w", err)
		} else if u.Months, err = cellToFloat(row, 8); err != nil {
			return fmt.Errorf("months: %w", err)
		} else if u.Hirelings, err = cellToInt(row, 9); err != nil {
			return fmt.Errorf("hirelings: %w", err)
		} else if u.Mercs, err = cellToInt(row, 10); err != nil {
			return fmt.Errorf("mercs: %w", err)
		} else if u.Locals, err = cellToInt(row, 11); err != nil {
			return fmt.Errorf("locals: %w", err)
		} else if u.Auxiliaries, err = cellToInt(row, 12); err != nil {
			return fmt.Errorf("auxiliaries: %w", err)
		} else if u.Workers, err = cellToInt(row, 13); err != nil {
			return fmt.Errorf("workers: %w", err)
		} else if u.Used, err = cellToInt(row, 14); err != nil {
			return fmt.Errorf("used: %w", err)
		} else if u.Remains, err = cellToInt(row, 15); err != nil {
			return fmt.Errorf("remains: %w", err)
		} else if u.Cattle, err = cellToInt(row, 16); err != nil {
			return fmt.Errorf("cattle: %w", err)
		} else if u.Dog, err = cellToInt(row, 17); err != nil {
			return fmt.Errorf("dog: %w", err)
		} else if u.Elephant, err = cellToInt(row, 18); err != nil {
			return fmt.Errorf("elephant: %w", err)
		} else if u.Goat, err = cellToInt(row, 19); err != nil {
			return fmt.Errorf("goat: %w", err)
		} else if u.Horse, err = cellToInt(row, 20); err != nil {
			return fmt.Errorf("horse: %w", err)
		} else if u.Camel, err = cellToInt(row, 21); err != nil {
			return fmt.Errorf("camel: %w", err)
		} else if u.Herders, err = cellToInt(row, 21); err != nil {
			return fmt.Errorf("herders: %w", err)
		}
		if w.ClanId == "" {
			w.ClanId = u.Unit
		}
		u.DeriveKind()
		if u.GT == u.Unit {
			u.GT = ""
		}
		w.Units[u.Unit] = u
	}

	return nil
}
