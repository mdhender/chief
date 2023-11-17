// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"github.com/mdhender/chief/internal/turnrpt"
	"github.com/mdhender/chief/internal/way"
	"net/http"
	"path/filepath"
)

func (s *Server) handleTurnReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		game, ok := s.games[way.Param(r.Context(), "game")]
		if !ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		clan, ok := game.Clans[way.Param(r.Context(), "clan")]
		if !ok {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		year := way.Param(r.Context(), "year")
		month := way.Param(r.Context(), "month")
		turnReportFile := filepath.Join(clan.Docs, fmt.Sprintf("%s.%s-%s.Turn-Report.docx", clan.Id, year, month))
		turnrpt.ParseToResponse(turnReportFile, w)
	}
}
