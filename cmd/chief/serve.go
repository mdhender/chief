// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/spf13/cobra"
	"log"
	"net"
	"net/http"
	"time"
)

// serveCmd implements the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start web server",
	Long:  `Start the web server.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// create a new http server with good values for timeouts and transports
		s := &Server{
			Server: http.Server{
				Addr:         net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
				IdleTimeout:  cfg.Server.IdleTimeout,
				ReadTimeout:  cfg.Server.ReadTimeout,
				WriteTimeout: cfg.Server.WriteTimeout,
			},
		}

		if s.starting, err = LoadStarting("genericStartup.json"); err != nil {
			log.Fatal(err)
		}
		s.Handler = s.routes()

		defer func(started time.Time) {
			log.Printf("[main] elapsed time %v\n", time.Now().Sub(started))
		}(time.Now())

		if err := s.Serve(cfg.Games); err != nil {
			log.Fatal(err)
		}
	},
}
