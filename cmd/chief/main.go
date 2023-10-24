// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements a player aid for TribeNet.
package main

import (
	"github.com/mdhender/chief/internal/dot"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	if err := dot.Load("CHIEF", true, true); err != nil {
		log.Fatalf("main: %+v\n", err)
	}
	host := ""
	if val := os.Getenv("CHIEF_HOST"); val != "" {
		host = val
	}
	port := "8080"
	if val := os.Getenv("CHIEF_PORT"); val != "" {
		port = val
	}

	var err error

	// create a new http server with good values for timeouts and transports
	s := &Server{}
	s.Addr = net.JoinHostPort(host, port)
	s.IdleTimeout = 10 * time.Second
	s.ReadTimeout = 2 * time.Second
	s.WriteTimeout = 2 * time.Second

	if s.starting, err = LoadStarting("genericStartup.json"); err != nil {
		log.Fatal(err)
	}
	s.Handler = s.routes()

	defer func(started time.Time) {
		log.Printf("[main] elapsed time %v\n", time.Now().Sub(started))
	}(time.Now())

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
