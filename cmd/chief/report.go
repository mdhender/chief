// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights rereportd.

package main

import (
	"fmt"
	"github.com/mdhender/chief/internal/docconv"
	"github.com/mdhender/chief/internal/lexer"
	parser "github.com/mdhender/chief/internal/parsers/turnrpt"
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
		doScanner := len(args) < 0
		doLexer := len(args) < 0
		doParser := len(args) >= 0
		for _, turn := range []string{"899-12", "900-01", "900-02"} {
			if turn != "899-12" {
				continue
			}
			filename := fmt.Sprintf("0138.%s.Turn-Report.docx", turn)
			log.Printf("[report] filename %s\n", filename)

			if doLexer {
				log.Printf("[lexer] filename %s\n", filename)
				fp, err := os.Open(filename)
				if err != nil {
					log.Printf("[lexer] open %v\n", err)
				} else {
					if body, _, err := docconv.ConvertDocx(fp); err != nil {
						log.Printf("[lexer] convert %v\n", err)
					} else {
						fib, err := lexer.Parse(filename, body, "{{", "}}", make(map[string]any))
						log.Println(filename, fib, err)
					}
				}
			}
			if doParser {
				log.Printf("[parser] filename %s\n", filename)
				rpt, err := parser.Parse(filename)
				if err != nil {
					log.Printf("[parser] %v\n", err)
				} else if rpt != nil {
					//log.Printf("[parser] report %s\n", rpt)
				}
			}
			if doScanner {
				sections, err := turnrpt.ParseDocument(filename)
				if err != nil {
					log.Println(err)
				} else {
					log.Printf("[scanner] sections %d\n", len(sections))
					for _, section := range sections {
						log.Printf("[scanner] section %+v\n", *section)
					}
				}
			}
			log.Printf("[report] completed.\n\n\n")
		}

		//fp, err := os.Open(filename)
		//if err != nil {
		//	log.Printf("[report] open %v\n", err)
		//	log.Fatal(err)
		//}
		//body, meta, err := docconv.ConvertDocx(fp)
		//if err != nil {
		//	log.Printf("[report] convert %v\n", err)
		//	log.Fatal(err)
		//}
		//log.Printf("[report] parsing %d bytes\n", len(body))
		//
		//sections, err := turnrpt.Parse([]byte(body), meta)

	},
}
