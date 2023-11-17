// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements a hex mapping application for TribeNet.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// clan, turn := "0138", "899-12"
	clan, turn, origin := "0108", "899-12", &Hex{14, 11}
	fmt.Printf("origin is %s\n", origin)

	scouts := make(map[string]*ScoutingParty)
	for i := 1; i <= 8; i++ {
		id := fmt.Sprintf("Scout %d", i)
		scouts[id] = &ScoutingParty{
			Id:    id,
			Start: origin,
		}
	}

	if err := LoadScoutingReport(fmt.Sprintf("%s.%s.Scouting-Report.txt", clan, turn), scouts); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d scouts\n", len(scouts))
	if len(scouts) != 9876 {
		os.Exit(0)
	}

	//routes := [][]MoveDirection{
	//	{SW, S},
	//	{S, S},
	//	{SE, S},
	//	{NW, SW, S},
	//	{NE, SE, S},
	//	{NE, SE, NE},
	//	{N, NE, NW, SE},
	//	{NW, SW, NW},
	//}
	//for n, route := range routes {
	//	scout, hex := n+1, origin
	//	log.Printf("scout %d: from %s move %v\n", scout, origin, route)
	//	for i, move := range route {
	//		next := hex.Move(move)
	//		log.Printf("scout %d: %2d: from %s move %2s to %s\n", scout, i+1, hex, move, next)
	//		hex = next
	//	}
	//}
}
