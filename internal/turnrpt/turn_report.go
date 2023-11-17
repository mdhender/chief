// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package turnrpt

import (
	"fmt"
	"github.com/mdhender/chief/internal/terrain"
)

type Report struct {
	Clan        string
	Tribe       string
	CurrentHex  string
	PreviousHex string
	CurrentTurn struct {
		Turn      string
		Something string
		Season    string
		Weather   string
	}
	NextTurn struct {
		Turn      string
		Something string
		Date      string
	}
	Accounting struct {
		Received int
		Cost     int
		Credit   int
	}
	GoodsTribe         string
	DesiredCommodities string
	ScoutingResults    []ScoutingResult
	Status             struct {
		Clan    string
		Terrain string
		Hex     string
	}
	Humans struct {
		People    int
		Warriors  int
		Actives   int
		Inactives int
	}
	Possessions struct {
		Animals struct {
			Cattle int
			Goat   int
			Horse  int
		}
		FinishedGoods struct {
			Provs int
			Sling int
			Trap  int
			Wagon int
		}
		Minerals struct {
			Brass  int
			Bronze int
			Coal   int
			Iron   int
			Silver int
		}
		RawMaterials struct {
			Bark    int
			Bone    int
			Gut     int
			Leather int
			Log     int
			Skin    int
			Wax     int
		}
		Ships        int
		WarEquipment struct {
			Club   int
			Jerkin int
			Shield int
			Sword  int
		}
	}
	Skills struct {
		Adm  int
		BnW  int
		Bon  int
		Cur  int
		Dip  int
		Eco  int
		Eng  int
		For  int
		Garr int
		Gut  int
		Herd int
		Hunt int
		Int  int
		Ldr  int
		Ltr  int
		Min  int
		Qry  int
		Sct  int
		ShB  int
		ShW  int
		Skn  int
		Tan  int
		Wd   int
	}
	Morale      int
	Weight      int
	Settlements []Settlement
}

type ScoutingResult struct {
	Id              string
	MovementResults []MovementResult
}

type MovementResult struct {
	Failed    bool
	Direction string
	Terrain   string
	Found     []string
	Notes     []string
}

type Settlement struct {
	HexCode string
	Name    string
	Note    string
	Type    string
	SubType string
}

type TribeSection struct {
	Id                 string
	CurrHex            *Hex
	PrevHex            *Hex
	Current            *Turn
	Next               *Turn
	Accounts           *Account
	GoodsTribe         string
	DesiredCommodities *DesiredCommodities
	Scouting           struct {
		Results []*ScoutingResult
	}
	Status *Status
}

type Account struct {
	Received string
	Cost     string
	Credit   string
}

func (a *Account) String() string {
	if a == nil {
		return "nil"
	}
	return fmt.Sprintf("{R %q C %q C %q}", a.Received, a.Cost, a.Credit)
}

type DesiredCommodities struct {
	One string
	Two string
}

type Status struct {
	Clan    string
	Terrain terrain.CODE
	Hex     string
}

type Turn struct {
	No      int
	Year    int
	Month   int
	Season  string
	Weather string
	Date    string
}

func (t *Turn) String() string {
	if t == nil {
		return "nil"
	} else if len(t.Season) != 0 {
		return fmt.Sprintf("{No: %d, Year: %d, Month: %d, Season: %q, Weather: %q}", t.No, t.Year, t.Month, t.Season, t.Weather)
	}
	return fmt.Sprintf("{No: %d, Year: %d, Month: %d, Date: %q}", t.No, t.Year, t.Month, t.Date)
}

type YearMonth struct {
	Year  int
	Month int
}
