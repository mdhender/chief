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

	Turn        string `json:"turn,omitempty"`
	CurrentHex  string `json:"current-hex,omitempty"`
	StartingHex string `json:"starting-hex,omitempty"`

	GoodsTribe         string              `json:"goods-tribe,omitempty"`
	DesiredCommodities *DesiredCommodities `json:"desired-commodities,omitempty"`
	GMNotes            string              `json:"gm-notes,omitempty"`
	TribeActivities    *TribeActivities    `json:"tribe-activities,omitempty"`
	FinalActivities    *FinalActivities    `json:"final-activities,omitempty"`
	TribeMovement      *TribeMovement      `json:"tribe-movement,omitempty"`
	ScoutActions       *ScoutActions       `json:"scout-actions,omitempty"`
	UnitStatus         *UnitStatus         `json:"unit-status,omitempty"`
	People             *People             `json:"people,omitempty"`
	Humans             *Humans             `json:"humans,omitempty"`
	Possessions        *Possessions        `json:"possessions,omitempty"`
	Skills             *Skills             `json:"skills,omitempty"`
	Morale             *Morale             `json:"morale,omitempty"`
	Weight             *Weight             `json:"weight,omitempty"`
	Truces             *Truces             `json:"truces,omitempty"`

	Bleet  string  `json:"junk,omitempty"`
	Errors []error `json:"errors,omitempty"`
}

type Animals struct {
	Bleet string
}

type CommonHeading struct {
	Turn        string
	StartingHex string
	CurrentHex  string
}

type DesiredCommodities struct {
	Bleet string `json:"bleet,omitempty"`
}

type FinalActivities struct {
	Bleet string `json:"bleet,omitempty"`
}

type FinishedGoods struct {
	Bleet string `json:"bleet,omitempty"`
}

type Humans struct {
	Bleet string `json:"bleet,omitempty"`
}

type Minerals struct {
	Bleet string `json:"bleet,omitempty"`
}

type Morale struct {
	Bleet string `json:"bleet,omitempty"`
}

type Movement struct {
	Direction string  `json:"direction,omitempty"`
	Terrain   string  `json:"terrain,omitempty"`
	Stay      bool    `json:"stay,omitempty"`
	Failed    bool    `json:"failed,omitempty"`
	Info      string  `json:"info,omitempty"`
	Errors    []error `json:"errors,omitempty"`
}

type People struct {
	Warriors int `json:"warriors,omitempty"`
	Active   int `json:"active,omitempty"`
	Inactive int `json:"inactive,omitempty"`
}

type Possessions struct {
	Animals       *Animals       `json:"animals,omitempty"`
	Minerals      *Minerals      `json:"minerals,omitempty"`
	WarEquipment  *WarEquipment  `json:"war-equipment,omitempty"`
	FinishedGoods *FinishedGoods `json:"finished-goods,omitempty"`
	RawMaterials  *RawMaterials  `json:"raw-materials,omitempty"`
	Ships         *Ships         `json:"ships,omitempty"`
}

type RawMaterials struct {
	Bleet string `json:"bleet,omitempty"`
}

type ScoutActions struct {
	Movements []*ScoutMovement `json:"movements,omitempty"`
	Errors    []error          `json:"errors,omitempty"`
}

type ScoutMovement struct {
	Id       int         `json:"id,omitempty"`
	Movement []*Movement `json:"moves,omitempty"`
	Errors   []error     `json:"errors,omitempty"`
	Bleet    string      `json:"bleet,omitempty"`
}

type Settlements struct {
	Bleet string `json:"bleet,omitempty"`
}

type Ships struct {
	Bleet string `json:"bleet,omitempty"`
}

type Skills struct {
	Bleet string `json:"bleet,omitempty"`
}

type Transfers struct {
	Bleet string `json:"bleet,omitempty"`
}

type TribeActivities struct {
	Bleet string `json:"bleet,omitempty"`
}

type TribeMovement struct {
	Follows  string      `json:"follows,omitempty"`
	Movement []*Movement `json:"moves,omitempty"`
	Bleet    string      `json:"bleet,omitempty"`
}

type Truces struct {
	Bleet string `json:"bleet,omitempty"`
}

type UnitStatus struct {
	Id      string   `json:"id,omitempty"`
	Terrain string   `json:"terrain,omitempty"`
	Edges   []string `json:"edges,omitempty"`
	Units   []string `json:"units,omitempty"`
	Bleet   string   `json:"bleet,omitempty"`
	Errors  []error  `json:"errors,omitempty"`
}

type WarEquipment struct {
	Bleet string `json:"bleet,omitempty"`
}

type Weight struct {
	Bleet string `json:"bleet,omitempty"`
}
