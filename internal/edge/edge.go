// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package edge

import (
	"fmt"
	"slices"
	"strings"
)

type Edge int

// the enums for Edge must be sorted by the Code string value.
// if they aren't, the UnmarshalJSON will fail.
const (
	Unknown Edge = iota
	ConiferHills
	GrassyHills
	OceanCoast
	Prairie
	River
	RiverFord
	RockyHills
	Swamp
	endOfCodes // used as a sentinel value
)

var (
	Codes = []string{
		ConiferHills: "CH",
		GrassyHills:  "GH",
		OceanCoast:   "OC",
		Prairie:      "PR",
		River:        "R",
		RiverFord:    "RF",
		RockyHills:   "RH",
		Swamp:        "SW",
	}
	Description = []string{
		ConiferHills: "Conifer Hills",
		GrassyHills:  "Grassy Hills",
		OceanCoast:   "Ocean Coast",
		Prairie:      "Prairie",
		River:        "River",
		RiverFord:    "River Fords",
		RockyHills:   "Rocky Hills",
		Swamp:        "Swamp",
	}
	LongDescription = []string{
		ConiferHills: "Conifer hills.",
		GrassyHills:  "Grassy hills.",
		OceanCoast:   "Ocean coastline.",
		Prairie:      "Prairie",
		River:        "River with no crossings.",
		RiverFord:    "Shallow spots that are ways across rivers.",
		RockyHills:   "Rocky hills.",
		Swamp:        "Swamp",
	}
)

func (e Edge) Code() string {
	if e < Unknown || endOfCodes <= e {
		panic(fmt.Sprintf("assert(edge != %d)", e))
	}
	return Codes[e]
}

// Description returns a description of the code.
func (e Edge) Description() string {
	if e < Unknown || endOfCodes <= e {
		panic(fmt.Sprintf("assert(edge != %d)", e))
	}
	return Description[e]
}

// LongDescription returns an expanded description of the code.
func (e Edge) LongDescription() string {
	if e < Unknown || endOfCodes <= e {
		panic(fmt.Sprintf("assert(edge != %d)", e))
	}
	return LongDescription[e]
}

// MarshalJSON implements the json.Marshaler interface
func (e Edge) MarshalJSON() ([]byte, error) {
	if e < Unknown || endOfCodes <= e {
		return nil, fmt.Errorf("unknown edge \"%d\"", e)
	} else if e == Unknown {
		return nil, nil
	}
	return []byte("\"" + Codes[e] + "\""), nil
}

// String implements the string.Stringer interface
func (e Edge) String() string {
	if e < Unknown || endOfCodes <= e {
		panic(fmt.Sprintf("assert(edge != %d)", e))
	}
	return Codes[e]
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (e *Edge) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("invalid edge %q", string(b))
	}
	var ok bool
	*e, ok = unmarshalCode(strings.ToUpper(string(b[1 : len(b)-1])))
	if !ok {
		return fmt.Errorf("unknown edge %s", string(b))
	}
	return nil
}

func unmarshalCode(s string) (Edge, bool) {
	if e, ok := slices.BinarySearch(Codes, s); ok {
		return Edge(e), true
	}
	switch s {
	case "CONIFER HILLS":
		return ConiferHills, true
	case "FORD":
		return RiverFord, true
	case "GRASSY HILLS":
		return GrassyHills, true
	case "PRAIRIE":
		return Prairie, true
	case "RIVER":
		return River, true
	case "ROCKY HILLS":
		return RockyHills, true
	case "O", "OCEAN":
		return OceanCoast, true
	case "SWAMP":
		return Swamp, true
	}
	return Unknown, false
}
