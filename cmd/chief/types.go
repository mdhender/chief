// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import "sync"

type Starting struct {
	sync.Mutex
	Game                  string   `json:"Game"`
	Clan                  string   `json:"Clan"`
	DesiredCommodities    string   `json:"DesiredCommodities"`
	Status                Status   `json:"Status"`
	Skills                Skills   `json:"Skills"`
	Morale                int      `json:"Morale"`
	Notes                 []string `json:"Notes"`
	RulesTweaksMoreToCome []string `json:"Rules tweaks (more to come)"`
}

type Animals struct {
	Cattle int `json:"Cattle"`
	Goat   int `json:"Goat"`
	Horse  int `json:"Horse"`
}

type FinishedGoods struct {
	Provs int `json:"Provs"`
	Sling int `json:"Sling"`
	Trap  int `json:"Trap"`
	Wagon int `json:"Wagon"`
}

type Humans struct {
	People    int `json:"People"`
	Warriors  int `json:"Warriors"`
	Actives   int `json:"Actives"`
	Inactives int `json:"Inactives"`
}

type Minerals struct {
	Brass  int `json:"Brass"`
	Bronze int `json:"Bronze"`
	Coal   int `json:"Coal"`
	Iron   int `json:"Iron"`
	Silver int `json:"Silver"`
}

type RawMaterials struct {
	Bark    int `json:"Bark"`
	Bone    int `json:"Bone"`
	Gut     int `json:"Gut"`
	Leather int `json:"Leather"`
	Logs    int `json:"Logs"`
	Skin    int `json:"Skin"`
	Wax     int `json:"Wax"`
}

type Skills struct {
	Adm  int `json:"Adm"`
	BnW  int `json:"BnW"`
	Bon  int `json:"Bon"`
	Cur  int `json:"Cur"`
	Dip  int `json:"Dip"`
	Eco  int `json:"Eco"`
	Eng  int `json:"Eng"`
	For  int `json:"For"`
	Garr int `json:"Garr"`
	Gut  int `json:"Gut"`
	Herd int `json:"Herd"`
	Hunt int `json:"Hunt"`
	Ldr  int `json:"Ldr"`
	Ltr  int `json:"Ltr"`
	Min  int `json:"Min"`
	Qry  int `json:"Qry"`
	Sct  int `json:"Sct"`
	Skn  int `json:"Skn"`
	Tan  int `json:"Tan"`
	Wd   int `json:"Wd"`
}

type Status struct {
	Humans        Humans        `json:"Humans"`
	Animals       Animals       `json:"Animals"`
	Minerals      Minerals      `json:"Minerals"`
	WarEquipment  WarEquipment  `json:"WarEquipment"`
	FinishedGoods FinishedGoods `json:"FinishedGoods"`
	RawMaterials  RawMaterials  `json:"RawMaterials"`
}

type WarEquipment struct {
	Club   int `json:"Club"`
	Jerkin int `json:"Jerkin"`
	Shield int `json:"Shield"`
	Sword  int `json:"Sword"`
}
