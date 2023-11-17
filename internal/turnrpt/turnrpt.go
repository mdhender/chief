// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package turnrpt implements a parser for Tribenet turn reports.
package turnrpt

import (
	"fmt"
	"github.com/mdhender/chief/internal/docconv"
	"log"
	"net/http"
	"os"
)

func ParseToResponse(filename string, w http.ResponseWriter) {
	log.Printf("parse: filename %s\n", filename)
	fp, err := os.Open(filename)
	if err != nil {
		log.Printf("[ptr] open %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	body, meta, err := docconv.ConvertDocx(fp)
	if err != nil {
		log.Printf("[ptr] convert %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "<h1>Turn Report Parser</h1>")
	_, _ = fmt.Fprintf(w, "<p>opened %q</p>\n", filename)

	_, _ = fmt.Fprintf(w, "<h2>Document Meta-Data</h2>")
	_, _ = fmt.Fprintf(w, "<ul>\n")
	for k, v := range meta {
		_, _ = fmt.Fprintf(w, "<li>key %q value %q</li>\n", k, v)
	}
	_, _ = fmt.Fprintf(w, "</ul>\n")

	_, _ = fmt.Fprintf(w, "<h2>Body</h2>\n")
	_, _ = fmt.Fprintf(w, "<pre>\n")
	_, _ = fmt.Fprintf(w, "<code>\n")
	_, _ = fmt.Fprintf(w, "%s\n", body)
	_, _ = fmt.Fprintf(w, "</code>\n")
	_, _ = fmt.Fprintf(w, "</pre>\n")

	reports, err := Parse([]byte(body), make(map[string]string))
	if err != nil {
		log.Printf("[ptr] parse %v\n", err)
	}
	for _, report := range reports {
		fmt.Printf("%+v\n", *report)
	}

	return
}
