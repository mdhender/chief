// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements a hex mapping application for TribeNet.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mdhender/chief/internal/stores/json/scouting"
	"github.com/mdhender/chief/internal/terrain"
	"github.com/mdhender/chief/internal/tiles"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "."
	flag.StringVar(&root, "root", root, "path to data files")
	flag.Parse()

	clan := "0138"
	var turns []*scouting.Results
	for _, turn := range []string{"899-12", "900-01", "900-02", "900-03"} {
		input := filepath.Join(root, turn, fmt.Sprintf("%s.%s.Scouting-Report.json", clan, turn))
		j, err := scouting.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}

		output := fmt.Sprintf("%s.%s.Scouting-Store.json", clan, turn)
		if data, err := json.MarshalIndent(&j, "", "  "); err != nil {
			log.Fatal(err)
		} else if err = os.WriteFile(output, data, 0644); err != nil {
			log.Fatal(err)
		}

		turns = append(turns, j)
	}

	// assume that "##" is actually "DA"
	hashValue := "DA"
	maps := tiles.New(hashValue)
	for _, turn := range turns {
		for _, unit := range turn.Units {
			if strings.HasPrefix(unit.Location.StartedIn, "##") {
				unit.Location.StartedIn = hashValue + unit.Location.StartedIn[2:]
			}
			if strings.HasPrefix(unit.Location.Current, "##") {
				unit.Location.Current = hashValue + unit.Location.Current[2:]
			}
			if unit.Check != nil && strings.HasPrefix(unit.Check.Hex, "##") {
				unit.Check.Hex = hashValue + unit.Check.Hex[2:]
			}
		}
	}

	// convert all successful moves into tiles
	for _, turn := range turns {
		for _, unit := range turn.Units {
			starting := maps.MakeTile(unit.Location.StartedIn)
			log.Printf("turn %s unit %s starting %s %s\n", turn.Turn, unit.Id, unit.Location.StartedIn, starting)
			current := mapMovement(maps, unit.Movement, starting)
			if unit.Scouts != nil {
				for _, id := range []string{"1", "2", "3", "4", "5", "6", "7", "8"} {
					scout := unit.Scouts[id]
					if scout == nil {
						continue
					}
					mapMovement(maps, scout.Scout, current)
				}
			}
			ending := maps.MakeTile(unit.Location.Current)
			log.Printf("turn %s unit %s ending   %s %s\n", turn.Turn, unit.Id, unit.Location.Current, ending)
			if unit.Check != nil {
				if len(unit.Check.Hex) != 7 {
					panic(fmt.Sprintf("turn %q unit %q hex %q", turn.Turn, unit.Id, unit.Check.Hex))
				}
				check := maps.MakeTile(unit.Check.Hex)
				log.Printf("turn %s unit %s check    %s %s\n", turn.Turn, unit.Id, unit.Check.Hex, check)
				if current.Id() != ending.Id() {
					log.Printf("turn %q unit %q: current: id != ending (%q, %q)\n", turn.Turn, unit.Id, current.Id(), ending.Id())
				}
				if ending.Id() != check.Id() {
					log.Printf("turn %q unit %q ending.id != check (%q, %q)\n", turn.Turn, unit.Id, ending.Id(), check.Id())
				} else {
					if check.Terrain != terrain.Unknown && ending.Terrain != check.Terrain {
						log.Printf("turn %q unit %q ending.terrain != check (%q, %q)\n", turn.Turn, unit.Id, ending.Terrain, check.Terrain)
					}
				}
			}
		}
	}

	s := tiles.NewSVG(true)
	for _, tile := range maps.Tiles() {
		log.Printf("dump tile %s %s\n", tile.Id(), tile.Terrain.String())
		s.AddTile(tile)
	}
	log.Printf("%s\n", string(s.Bytes()))
}

func mapMovement(maps *tiles.Map, moves []*scouting.Movement, current *tiles.Tile) *tiles.Tile {
	for n, move := range moves {
		if move.Result.Failed != nil {
			continue
		}
		id := current.Id()
		switch move.Direction {
		case "N":
			current = maps.Neighbor(current, tiles.N)
		case "NE":
			current = maps.Neighbor(current, tiles.NE)
		case "SE":
			current = maps.Neighbor(current, tiles.SE)
		case "S":
			current = maps.Neighbor(current, tiles.S)
		case "SW":
			current = maps.Neighbor(current, tiles.SW)
		case "NW":
			current = maps.Neighbor(current, tiles.NW)
		default:
			panic(fmt.Sprintf("assert(direction != %q)", move.Direction))
		}
		current.Terrain = move.Result.Terrain
		log.Printf("move %d from %s %-2s to %s (%s)\n", n+1, id, move.Direction, current.Id(), current.Terrain.String())
	}
	return current
}
