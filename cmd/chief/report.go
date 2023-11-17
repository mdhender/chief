// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights rereportd.

package main

import (
	"github.com/mdhender/chief/internal/docconv"
	"github.com/mdhender/chief/internal/turnrpt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// reportCmd implements the report command.
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "parse a turn report",
	Long:  `Load a turn report and create the data from it.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := "0138.899-12.Turn-Report.docx"
		log.Printf("[report] filename %s\n", filename)
		fp, err := os.Open(filename)
		if err != nil {
			log.Printf("[report] open %v\n", err)
			log.Fatal(err)
		}
		body, meta, err := docconv.ConvertDocx(fp)
		if err != nil {
			log.Printf("[report] convert %v\n", err)
			log.Fatal(err)
		}
		log.Printf("[report] parsing %d bytes\n", len(body))

		sections, err := turnrpt.Parse([]byte(body), meta)
		log.Printf("[report] sections %d\n", len(sections))
		if err != nil {
			log.Printf("[report] error %v\n", err)
		}
		for _, section := range sections {
			log.Printf("[report] section %+v\n", *section)
		}
	},
}
