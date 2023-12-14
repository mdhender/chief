// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements a turn report parser.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	parser "github.com/mdhender/chief/internal/parsers/pigeon/turnrpt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	clan := "0138"
	flag.StringVar(&clan, "clan", clan, "clan id")
	grid := "AA"
	flag.StringVar(&grid, "grid", grid, "location of grid (AA..ZZ)")
	root := "."
	flag.StringVar(&root, "root", root, "path to data files")

	// Set custom usage function
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "usage: turnrpt [options] list of turns to parse\n")
		flag.PrintDefaults()
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "notes: if no turns are given, the root path will be scanned for turn folders.\n")
	}

	flag.Parse()

	// turns defaults to the remaining command line arguments.
	// If there are none, then use use the `turnFolders()` function
	// to generate the list of turns from folders in the root directory.
	turns := flag.Args()
	if len(turns) == 0 {
		var err error
		if turns, err = turnFolders(root); err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("parsing %+v\n", turns)
	for _, turn := range turns {
		err := parseReport(root, clan, turn, grid)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("all turn reports parsed\n")
}

func parseReport(root, clan, turn, grid string) error {
	filename := filepath.Join(root, turn, fmt.Sprintf("%s.%s.Turn-Report.txt", clan, turn))
	if sb, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("turn report file does not exist: %s", filename)
		}
		return err
	} else if sb.IsDir() {
		return fmt.Errorf("turn report is not a file: %s", filename)
	}
	log.Printf("parsing %s\n", filename)

	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// apply filters to the input
	if len(grid) == 2 {
		input = parser.FilterDefaultGrid(input, grid)
	} else {
		input = parser.FilterDefaultGrid(input, "LL")
	}
	input = parser.TransformMarkScoutLines(input)

	// parse the turn report
	raw, err := parser.Parse(filename, input)
	if err != nil {
		return err
	}
	rpt := raw.(*parser.Report)
	rpt.FileName = filename
	rpt.Clan = clan
	rpt.Turn = turn
	if len(rpt.Rest) > 35 {
		rpt.Rest = rpt.Rest[:35]
	}

	for _, r := range rpt.T {
		if len(r.Bleet) > 35 {
			r.Bleet = r.Bleet[:35]
		}
		if len(r.GMNotes) > 35 {
			r.GMNotes = r.GMNotes[:35]
		}
	}

	data, err := json.MarshalIndent(rpt, "", "\t")
	if err != nil {
		return err
	}
	filename = filepath.Join(root, turn, fmt.Sprintf("%s.%s.Turn-Report.json", clan, turn))
	if err = os.WriteFile(filename, data, 0644); err != nil {
		return err
	}
	log.Printf("created %s\n", filename)

	return nil
}

// turnFolders reads the directory specified by the path and returns
// a slice of directory names that match a specific pattern.
// The pattern it matches is a three-digit year, followed by a dash,
// and a two-digit month, formatted as "YYY-MM".
// It returns two values: a slice of strings and an error.
// The slice of strings is a list of directory names that match the
// pattern.
// If an error occurs during the reading of the directory or other
// operations, it will be returned as the error value.
//
// Parameters:
//   - path (string): The path to the directory that the function should read.
//
// Returns:
//   - ([]string, error): A slice of strings containing the names of
//     directories that match the pattern and an
//     error if any occurred.
func turnFolders(path string) (turns []string, err error) {
	// turns look like YYY-MM (year and month)
	reTurn := regexp.MustCompile(`^\d{3}-\d{2}$`)

	// fetch all entries from the path
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// find the turn folders
	for _, entry := range entries {
		if entry.IsDir() && reTurn.MatchString(entry.Name()) {
			turns = append(turns, entry.Name())
		}
	}
	return turns, nil
}
