// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
)

func (w *workbook) setupLoaders() error {
	// Get value from cell by given worksheet name and axis.
	cell, err := w.f.GetCellValue("Instructions", "B1")
	if err != nil {
		return fmt.Errorf("get version: %w", err)
	}
	w.Version = cell
	log.Printf("%s: version %s\n", w.Name, w.Version)

	switch w.Version {
	case "1.11": // 2023/08/16
		w.loaders = append(w.loaders, loader{id: "tabCheck.v1.11", load: w.tabCrossCheck("1.11")})
		w.loaders = append(w.loaders, loader{id: "clan.v1.11", load: w.loadClan("1.11")})
		w.loaders = append(w.loaders, loader{id: "comments.v1.11", load: w.loadComments("1.11")})
		w.loaders = append(w.loaders, loader{id: "gmActions.v1.11", load: w.loadGMActions("1.11")})
		w.loaders = append(w.loaders, loader{id: "transfers.v1.11", load: w.loadTransfers("1.11")})
		w.loaders = append(w.loaders, loader{id: "tribeMovement.v1.11", load: w.loadTribeMovement("1.11")})
		w.loaders = append(w.loaders, loader{id: "scoutMovement.v1.11", load: w.loadScoutMovement("1.11")})
	}
	return nil
}
