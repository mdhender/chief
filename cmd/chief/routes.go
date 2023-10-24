// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"github.com/mdhender/chief/internal/way"
	"net/http"
)

func (s *Server) routes() http.Handler {
	r := way.NewRouter()

	// create routes
	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hi")
	})

	return r
}
