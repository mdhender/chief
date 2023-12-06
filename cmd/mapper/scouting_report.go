// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func LoadScoutingReport(filename string, scouts map[string]*ScoutingParty) error {
	rpt, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	terrain := make(map[string]string)
	terrain["1108"] = "PR"

	for no, line := range strings.Split(string(rpt), "\n") {
		id, orders, ok := strings.Cut(strings.TrimSpace(line), ":")
		if !ok {
			log.Fatalf("scouting report: %d: invalid input: missing ':'", no+1)
		}
		// fmt.Printf("%d: id %q: orders %q\n", no+1, id, move)
		scout, ok := scouts[id]
		if !ok {
			log.Fatalf("scouting report: %d: scout %q: unknown scout\n", no+1, id)
		} else if len(scout.RawText) != 0 {
			log.Fatalf("scouting report: %d: scout %q: has multiple records\n", no+1, id)
		}
		scout.RawText = line
		action, args, ok := strings.Cut(strings.TrimSpace(orders), " ")
		if !ok {
			log.Fatalf("scouting report: %d: scout %q: missing ' '\n", no+1, id)
		} else if action != "Scout" {
			log.Fatalf("scouting report: %d: scout %q: unknown action %q\n", no+1, id, action)
		}
		// fmt.Printf("args are %q\n", args)
		// move is something like
		//    DIRECTION-TERRAIN,COMMENT,COMMENT...
		// or COMMENT-ABOUT-FAILURE
		//move := &Move{}
		from := scout.Start
		for n, arg := range strings.Split(strings.TrimSpace(args), "\\") {
			arg = strings.TrimSpace(arg)
			// fmt.Printf("arg: %2d: %q\n", n, arg)

			move := Move{From: from, Result: &MoveResult{}}

			// simplest thing that could possibly work
			if strings.HasSuffix(arg, ", Nothing of interest found") {
				move.Result.Found = "Nothing of interest"
				arg = strings.TrimSuffix(arg, ", Nothing of interest found")
				// fmt.Printf("ar.: %2d: %q\n", n, arg)
			}
			notEnough := false
			if strings.HasPrefix(arg, "Can't Move on Ocean to ") {
				move.Result.Failed, move.Result.Blocked.Coastline, arg = true, true, arg[23:]
			} else if strings.HasPrefix(arg, "No Ford on River to ") {
				move.Result.Failed, move.Result.Blocked.NoFord, arg = true, true, arg[20:]
			} else if strings.HasPrefix(arg, "Not enough M.P's to move to ") {
				move.Result.Failed, notEnough, arg = true, true, arg[28:]
			}
			if move.Result.Failed {
				// fmt.Printf("...: %2d: %q\n", n, arg)
			}
			if strings.HasPrefix(arg, "N-") || strings.HasPrefix(arg, "N ") {
				move.Direction = N
				arg = arg[2:]
			} else if strings.HasPrefix(arg, "NE-") || strings.HasPrefix(arg, "NE ") {
				move.Direction = NE
				arg = arg[3:]
			} else if strings.HasPrefix(arg, "SE-") || strings.HasPrefix(arg, "SE ") {
				move.Direction = SE
				arg = arg[3:]
			} else if strings.HasPrefix(arg, "S-") || strings.HasPrefix(arg, "S ") {
				move.Direction = S
				arg = arg[2:]
			} else if strings.HasPrefix(arg, "SW-") || strings.HasPrefix(arg, "SW ") {
				move.Direction = SW
				arg = arg[3:]
			} else if strings.HasPrefix(arg, "NW-") || strings.HasPrefix(arg, "NW ") {
				move.Direction = NW
				arg = arg[3:]
			} else {
				log.Fatalf("scouting report: %d: scout %q: move %d: unknown value %q\n", no+1, id, n+1, arg)
			}

			to := move.From.Move(move.Direction)
			move.To = &to

			//fmt.Printf(".rg: %2d: %q\n", n, arg)
			for _, result := range strings.Split(arg, ",") {
				result = strings.TrimSpace(result)
				switch result {
				case "GH", "PR", "RH":
					terrain[to.String()] = result
				case "into CONIFER HILLS":
					terrain[to.String()] = "CH"
				case "into GRASSY HILLS":
					terrain[to.String()] = "GH"
				case "into ROCKY HILLS":
					terrain[to.String()] = "RH"
				case "of HEX": // ignore
				case "O S":
					// flag an Ocean to the hex south of here?
				case "River NE SE S":
				case "River SE":
				case "River SE S":
				case "River S":
				default:
					// fmt.Printf(".rg: %2d: %2d: %q\n", n, rno+1, result)
				}
			}

			if strings.HasPrefix(arg, "GH") {
			}

			if notEnough {
				fmt.Printf("%s: turn %2d: from %s %s move %2s to %s %s (failed)\n", scout.Id, n+1, move.From, terrain[move.From.String()], move.Direction, move.To, terrain[move.To.String()])
			} else {
				fmt.Printf("%s: turn %2d: from %s %s move %2s to %s %s %q\n", scout.Id, n+1, move.From, terrain[move.From.String()], move.Direction, move.To, terrain[move.To.String()], arg)
			}

			from = move.To
		}
	}

	return nil
}
