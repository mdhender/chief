// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

//go:generate pigeon -o parser.go grammar.peg

// Package parser implements a parser for turn reports.
package parser

type Report struct {
	FileName string
	Clan     string
	Turn     string

	T []*TribeReport

	// Rest is all input after we hit our first error?
	Rest string
}

type TribeReport struct {
	Id string

	GMNotes string

	TribeActivities *TribeActivities
	FinalActivities *FinalActivities
	TribeMovement   *TribeMovement

	Bleet  string
	Errors []error
}

type FinalActivities struct {
	Bleet string
}

type TribeMovement struct {
	Bleet string
}

type TribeActivities struct {
	Bleet string
}
