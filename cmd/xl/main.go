// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements a program to load an Excel spreadsheet.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	var clan string
	flag.StringVar(&clan, "clan", clan, "clan")
	var doIssued bool = false
	flag.BoolVar(&doIssued, "issued", doIssued, "report on orders as issued")
	var doReceived bool = true
	flag.BoolVar(&doReceived, "received", doReceived, "report on orders as received")
	var filename string
	flag.StringVar(&filename, "input", filename, "xlsx file to load")
	var turn string
	flag.StringVar(&turn, "turn", turn, "turn to load")
	flag.Parse()

	if clan == "" {
		clan = "0138"
	}
	if filename != "" {
		log.Printf("ignoring %q on command line\n", filename)
		filename = ""
	}

	turns := []string{"899-12", "900-01", "900-02", "900-03", "900-04"}
	if turn != "" {
		turns = []string{turn}
	}
	for _, turn := range turns {
		if doReceived {
			if err := run(clan, turn, doReceived); err != nil {
				log.Printf("%s: %s: error %v\n", clan, turn, err)
			}
		}
		if doIssued {
			if err := run(clan, turn, !doIssued); err != nil {
				log.Printf("%s: %s: error %v\n", clan, turn, err)
			}
		}
	}
}

func run(clan, turn string, received bool) error {
	var filename, jsonFilename string
	revision := 0 // updated only for orders-issued
	if received {
		filename = fmt.Sprintf("%s.%s.Orders.xlsx", clan, turn)
		jsonFilename = fmt.Sprintf("output/%s.%s.received.json", clan, turn)
	} else { // orders issued
		maxRevision := -1
		if fnams, err := filepath.Glob(fmt.Sprintf("%s.%s.Orders-Issued.v*.xlsx", clan, turn)); err == nil {
			for _, fnam := range fnams {
				vers := strings.TrimPrefix(strings.TrimSuffix(fnam, ".xlsx"), fmt.Sprintf("%s.%s.Orders-Issued.v", clan, turn))
				if i, err := strconv.Atoi(vers); err == nil {
					if i > maxRevision {
						maxRevision = i
					}
				}
			}
		}
		if maxRevision < 0 {
			filename = fmt.Sprintf("%s.%s.Orders-Issued.xlsx", clan, turn)
		} else {
			revision = maxRevision
			filename = fmt.Sprintf("%s.%s.Orders-Issued.v%d.xlsx", clan, turn, revision)
		}
		jsonFilename = fmt.Sprintf("output/%s.%s.issued.json", clan, turn)
	}

	f, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}
	w := &workbook{Name: filename, Revision: revision, f: f, Units: make(map[string]*Unit)}

	foundInstructions := false
	for _, tab := range w.f.GetSheetList() {
		if tab == "Instructions" {
			foundInstructions = true
		}
	}
	if !foundInstructions {
		log.Printf("%s: %s: invalid workbook: missing tab %q\n", clan, turn, "Instructions")
		return fmt.Errorf("missing tab %q", "Instructions")
	}

	if err := w.setupLoaders(); err != nil {
		return fmt.Errorf("instructions: %w", err)
	}

	for _, loader := range w.loaders {
		err := loader.load()
		if err != nil {
			return fmt.Errorf("loader: %s: %w", loader.id, err)
		}
		log.Printf("%s: %s: loaded %s\n", clan, turn, loader.id)
	}

	if w.ClanId != clan {
		log.Printf("%s: %s: warning: expected clan %q, got %q\n", clan, turn, clan, w.ClanId)
	}

	data, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return fmt.Errorf("json: %w", err)
	} else if err = os.WriteFile(jsonFilename, data, 0644); err != nil {
		return fmt.Errorf("save: %w", err)
	}
	log.Printf("%s: %s: created %s\n", clan, turn, jsonFilename)

	return nil
}

type workbook struct {
	Name      string           `json:"name"`
	Version   string           `json:"version"`            // version of the workbook
	Revision  int              `json:"revision,omitempty"` // revision number for orders issued
	ClanId    string           `json:"clanId"`
	Units     map[string]*Unit `json:"units,omitempty"`
	Transfers *Transfers       `json:"transfers,omitempty"`
	f         *excelize.File
	loaders   []loader
}

type loader struct {
	id   string
	load func() error
}

type Unit struct {
	Unit        string    `json:"unit,omitempty"`
	Kind        string    `json:"kind,omitempty"`
	GT          string    `json:"goods-tribe,omitempty"`
	Warrior     int       `json:"warrior,omitempty"`
	Active      int       `json:"active,omitempty"`
	Inactive    int       `json:"inactive,omitempty"`
	Slave       int       `json:"slave,omitempty"`
	Eaters      int       `json:"eaters,omitempty"`
	Provs       int       `json:"provs,omitempty"`
	Months      float64   `json:"months,omitempty"`
	Hirelings   int       `json:"hirelings,omitempty"`
	Mercs       int       `json:"mercs,omitempty"`
	Locals      int       `json:"locals,omitempty"`
	Auxiliaries int       `json:"auxiliaries,omitempty"`
	Workers     int       `json:"workers,omitempty"`
	Used        int       `json:"used,omitempty"`
	Remains     int       `json:"remains,omitempty"`
	Cattle      int       `json:"cattle,omitempty"`
	Dog         int       `json:"dog,omitempty"`
	Elephant    int       `json:"elephant,omitempty"`
	Goat        int       `json:"goat,omitempty"`
	Horse       int       `json:"horse,omitempty"`
	Camel       int       `json:"camel,omitempty"`
	Herders     int       `json:"herders,omitempty"`
	Comments    []string  `json:"comments,omitempty"`
	GMRequests  []string  `json:"gm-requests,omitempty"`
	Movement    *Movement `json:"movement,omitempty"`
	Scouts      []*Scout  `json:"scouts,omitempty"`
	New         bool      `json:"new,omitempty"` // true if unit was created on the fly
}

func NewUnit(id string) *Unit {
	u := &Unit{Unit: id, New: true}
	u.DeriveKind()
	return u
}

func (u *Unit) DeriveKind() {
	switch len(u.Unit) {
	case 4:
		u.Kind = "Tribe"
	case 6:
		switch u.Unit[4] {
		case 'c':
			u.Kind = "Courier"
		case 'e':
			u.Kind = "Element"
		default:
			u.Kind = "Unknown"
		}
	default:
		u.Kind = "Unknown"
	}
}

type Transfers struct {
	BeforeMovement []*Transfer `json:"before-movement,omitempty"`
	AfterMovement  []*Transfer `json:"after-movement,omitempty"`
}

type Transfer struct {
	From           string `json:"from"` // id of source unit
	To             string `json:"to"`   // id of destination unit
	Item           string `json:"item"`
	Quantity       int    `json:"quantity,omitempty"`
	TransferTiming string `json:"transfer-timing,omitempty"`
	Notes          string `json:"notes,omitempty"`
	Processed      bool   `json:"processed,omitempty"`
}

type Movement struct {
	Hex       string  `json:"hex,omitempty"`    // if set, starting hex of tribe
	Follow    string  `json:"follow,omitempty"` // if set, follow this tribe
	Moves     []*Move `json:"moves,omitempty"`  // empty when following
	Processed bool    `json:"processed,omitempty"`
}

type Move struct {
	Direction string `json:"direction,omitempty"`
	Still     bool   `json:"still,omitempty"` // if set, tribe will not move
	ToLimit   bool   `json:"to-limit,omitempty"`
}

type Scout struct {
	Scouts    int          `json:"scouts,omitempty"`
	Horses    int          `json:"horses,omitempty"`
	Elephants int          `json:"elephants,omitempty"`
	Camels    int          `json:"camels,omitempty"`
	Mission   string       `json:"mission,omitempty"`
	Moves     []*ScoutMove `json:"moves,omitempty"`
	Processed bool         `json:"processed,omitempty"`
}

type ScoutMove struct {
	Direction string `json:"direction,omitempty"`
	ToLimit   bool   `json:"to-limit,omitempty"`
}
