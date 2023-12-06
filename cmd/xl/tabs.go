// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
)

func (w *workbook) tabCrossCheck(version string) func() error {
	switch version {
	case "1.11":
		return w.tabCrossCheckV1
	}
	panic(fmt.Sprintf("assert(version != %q)", version))
}

func (w *workbook) tabCrossCheckV1() error {
	return w.tabCheck(
		"Instructions",
		"Clan",
		"Comments",
		"GM Actions",
		"Transfers",
		"Tribe_Movement",
		"Scout_Movement",
		"Tribes_Activities",
		"Skill_Attempts",
		"Research_Attempts",
		"Valid_Skills",
		"Clan_Goods",
		"Clan_Research",
		"Valid Units",
		"Valid Activity",
		"Valid_Implements",
		"Valid Goods",
		"Valid_Research",
		"Transfer Codes",
		"References",
	)
}

func (w *workbook) tabCheck(tabs ...string) error {
	tm := make(map[string]bool)
	for _, tab := range tabs {
		tm[tab] = false
	}

	// check the workbook for unknown worksheets
	foundUnknown := false
	for _, tab := range w.f.GetSheetList() {
		if _, ok := tm[tab]; !ok {
			log.Printf("%s: unknown tab %q\n", w.Name, tab)
			foundUnknown = true
		} else {
			tm[tab] = true
		}
	}

	// verify that all the expected tabs are present
	missingExpected := false
	for tab, found := range tm {
		if !found {
			log.Printf("%s: missing tab %q\n", w.Name, tab)
			missingExpected = true
		}
	}

	if foundUnknown && missingExpected {
		return fmt.Errorf("missing/unknown tabs in workbook")
	} else if foundUnknown {
		return fmt.Errorf("unknown tabs in workbook")
	} else if missingExpected {
		return fmt.Errorf("missing tabs in workbook")
	}

	return nil
}
