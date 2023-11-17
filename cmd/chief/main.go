// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements a player aid for TribeNet.
package main

import (
	"github.com/mdhender/chief/internal/config"
	"github.com/mdhender/chief/internal/dot"
	"log"
	"os"
	"path/filepath"
)

// globals. ugh
var (
	cfg *config.Config = config.Default()
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	if err := cfg.Load("chief.json"); err != nil {
		log.Fatal(err)
	}
	for _, game := range cfg.Games {
		log.Printf("[%s] year %d month %d\n", game.Id, game.Turn.Year, game.Turn.Month)
		for _, clan := range game.Clans {
			if clan.Docs == "" {
				clan.Docs = "docs"
			}
			if docs, err := filepath.Abs(filepath.Join(clan.Root, clan.Docs)); err != nil {
				log.Fatal(err)
			} else {
				clan.Docs = docs
			}
			log.Printf("[%s] clan %q: docs %q\n", game.Id, clan.Id, clan.Docs)
		}
	}

	if err := dot.Load("CHIEF", true, true); err != nil {
		log.Fatalf("main: %+v\n", err)
	}
	if val := os.Getenv("CHIEF_ENV"); val != "" {
		cfg.Env = val
	}
	if val := os.Getenv("CHIEF_HOST"); val != "" {
		cfg.Server.Host = val
	}
	if val := os.Getenv("CHIEF_PORT"); val != "" {
		cfg.Server.Port = val
	}

	Execute()
}
