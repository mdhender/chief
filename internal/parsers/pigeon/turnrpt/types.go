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
	Rest string `json:"rest,omitempty"`
}

type TribeReport struct {
	Id string `json:"id,omitempty"`

	GMNotes string `json:"GMNotes,omitempty"`

	TribeActivities *TribeActivities `json:"tribe-activities,omitempty"`
	FinalActivities *FinalActivities `json:"final-activities,omitempty"`
	TribeMovement   *TribeMovement   `json:"tribe-movement,omitempty"`
	ScoutMovement   *ScoutMovement   `json:"scout-movement,omitempty"`
	UnitStatus      *UnitStatus      `json:"unit-status,omitempty"`
	Humans          *Humans          `json:"humans,omitempty"`
	Animals         *Animals         `json:"animals,omitempty"`
	Minerals        *Minerals        `json:"minerals,omitempty"`
	WarEquipment    *WarEquipment    `json:"war-equipment,omitempty"`
	FinishedGoods   *FinishedGoods   `json:"finished-goods,omitempty"`
	RawMaterials    *RawMaterials    `json:"raw-materials,omitempty"`
	Ships           *Ships           `json:"ships,omitempty"`
	Skills          *Skills          `json:"skills,omitempty"`
	Morale          *Morale          `json:"morale,omitempty"`
	Weight          *Weight          `json:"weight,omitempty"`
	Truces          *Truces          `json:"truces,omitempty"`

	Bleet  string  `json:"junk,omitempty"`
	Errors []error `json:"errors,omitempty"`
}

type Animals struct {
	Bleet string
}

type FinalActivities struct {
	Bleet string
}

type FinishedGoods struct {
	Bleet string
}

type Humans struct {
	Bleet string
}

type Minerals struct {
	Bleet string
}

type Morale struct {
	Bleet string
}

type RawMaterials struct {
	Bleet string
}

type ScoutMovement struct {
	Bleet string
}

type Settlements struct {
	Bleet string
}

type Ships struct {
	Bleet string
}

type Skills struct {
	Bleet string
}

type Transfers struct {
	Bleet string
}

type TribeMovement struct {
	Bleet string
}

type TribeActivities struct {
	Bleet string
}

type Truces struct {
	Bleet string
}

type UnitStatus struct {
	Bleet string
}

type WarEquipment struct {
	Bleet string
}

type Weight struct {
	Bleet string
}
