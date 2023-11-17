// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements a program to load an Excel spreadsheet.
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	var filename string
	flag.StringVar(&filename, "input", filename, "xlsx file to load")
	flag.Parse()

	if filename == "" {
		log.Fatal("error: missing input file name\n")
	} else if !strings.HasSuffix(filename, ".xlsx") {
		filename = filename + ".xlsx"
	}

	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
