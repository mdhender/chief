// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
)

func (w *workbook) checkHeaders(tab string, cells []string, expect ...string) error {
	var errors, idx int
	// invalid cells
	for idx = 0; idx < len(cells) && idx < len(expect); idx++ {
		if cells[idx] == expect[idx] {
			continue
		}
		log.Printf("%s: %s: header: want %q, got %q\n", w.Name, tab, cells[idx], expect[idx])
		errors++
	}
	// extra cells
	for ; idx < len(cells); idx++ {
		log.Printf("%s: %s: header: unknown %q\n", w.Name, tab, cells[idx])
		errors++
	}
	// missing cells
	for ; idx < len(expect); idx++ {
		log.Printf("%s: %s: header: missing %q\n", w.Name, tab, expect[idx])
		errors++
	}

	if errors != 0 {
		return fmt.Errorf("invalid worksheet header")
	}
	return nil
}
