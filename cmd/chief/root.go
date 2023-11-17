// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/spf13/cobra"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chief",
	Short: "The base command, which does not much",
	Long:  `The base command for chief. It does very little.`,
}

// Execute wires all the commands and sub-commands together.
// It is called only by main().
func Execute() {
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
