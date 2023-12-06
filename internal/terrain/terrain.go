// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package terrain

import (
	"fmt"
	"slices"
	"strings"
)

type Terrain int

const (
	Unknown Terrain = iota
	ALPS
	AR
	BH
	BR
	CH
	DE
	DF
	DH
	FORDS
	GH
	HSM
	JG
	JH
	L
	LCM
	LJM
	LSM
	O
	PI
	PR
	R
	RH
	SH
	SW
	TU
	endOfCodes // used as a sentinel value
)

var (
	Codes = []string{
		ALPS:  "ALPS",
		AR:    "AR",
		BH:    "BH",
		BR:    "BR",
		CH:    "CH",
		DE:    "DE",
		DF:    "DF",
		DH:    "DH",
		FORDS: "FORDS",
		GH:    "GH",
		HSM:   "HSM",
		JG:    "JG",
		JH:    "JH",
		L:     "L",
		LCM:   "LCM",
		LJM:   "LJM",
		LSM:   "LSM",
		O:     "O",
		PI:    "PI",
		PR:    "PR",
		R:     "R",
		RH:    "RH",
		SH:    "SH",
		SW:    "SW",
		TU:    "TU",
	}
	Description = []string{
		ALPS:  "Alpine",
		AR:    "Arid:",
		BH:    "Brush Hill",
		BR:    "Brush",
		CH:    "Conifer Hills",
		DE:    "Desert",
		DF:    "Deciduous Forest",
		DH:    "Deciduous Hill",
		FORDS: "Fords",
		GH:    "Grassy Hills",
		HSM:   "High Mountains",
		JG:    "Jungle",
		JH:    "Jungle Hill",
		L:     "Lake",
		LCM:   "Low Conifer Mountains",
		LJM:   "Low Jungle Mountain",
		LSM:   "Low Snowy Mountains",
		O:     "Ocean",
		PI:    "Polar Ice",
		PR:    "Prairie",
		R:     "River",
		RH:    "Rocky Hills",
		SH:    "Snowy Hills",
		SW:    "Swamp",
		TU:    "Tundra",
	}
	LongDescription = [...]string{
		ALPS:  "a bigger version of HSM.",
		AR:    "Arid: tundra without water.",
		BH:    "Brush Hill: Hill covered with brush.",
		BR:    "Brush: Conifer forest with fewer trees more bushes (Forestry not possible here).",
		CH:    "Conifer Hill: Hill covered with conifer forest.",
		DE:    "Desert: Arid without grass.",
		DF:    "Deciduous: Seasonal forest.",
		DH:    "Deciduous Hill: Hill covered with deciduous forest.",
		FORDS: "Shallow spots that are ways across rivers.",
		GH:    "Grassy Hill: Hill covered with grass.",
		HSM:   "High Mountains: cannot be entered, except through a pass.",
		JG:    "Jungle: Wet forest.",
		JH:    "Jungle Hill: Hill covered with jungle.",
		L:     "Lake a body of water.",
		LCM:   "Low Conifer Mountains: Hills but higher, difficult to enter covered with conifer forest.",
		LJM:   "Low Jungle Mountain.",
		LSM:   "Low Snowy Mountains: Hills but higher, very difficult to enter.",
		O:     "Ocean.",
		PI:    "Polar Ice: Permanent ice and difficult to move through.",
		PR:    "Prairie: Grassland.",
		R:     "Rivers: Large moving bodies of water impossible to cross unless through a ford or by boat.",
		RH:    "Rocky Hill: Hill covered with rocks.",
		SH:    "Snow Hill: colder than GH, snow rather than grass.",
		SW:    "Swamp: very wet grassland.",
		TU:    "Tundra: not very good grassland.",
	}
)

func (c Terrain) Code() string {
	if c < Unknown || endOfCodes <= c {
		panic(fmt.Sprintf("assert(terrain != %d)", c))
	}
	return Codes[c]
}

// Description returns a description of the code.
func (c Terrain) Description() string {
	if c < Unknown || endOfCodes <= c {
		panic(fmt.Sprintf("assert(terrain != %d)", c))
	}
	return Description[c]
}

// LongDescription returns an expanded description of the terrain code.
func (c Terrain) LongDescription() string {
	if c < Unknown || endOfCodes <= c {
		panic(fmt.Sprintf("assert(terrain != %d)", c))
	}
	return LongDescription[c]
}

// MarshalJSON implements the json.Marshaler interface
func (c Terrain) MarshalJSON() ([]byte, error) {
	if c < Unknown || endOfCodes <= c {
		return nil, fmt.Errorf("unknown terrain")
	} else if c == Unknown {
		return nil, nil
	}
	return []byte("\"" + Codes[c] + "\""), nil
}

// String implements the string.Stringer interface
func (c Terrain) String() string {
	if c < Unknown || endOfCodes <= c {
		panic(fmt.Sprintf("assert(terrain != %d)", c))
	}
	return Codes[c]
}

func (c Terrain) ToFill() string {
	return "**fill**"
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *Terrain) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("invalid terrain %q", string(b))
	}
	var ok bool
	*c, ok = unmarshalCode(strings.ToUpper(string(b[1 : len(b)-1])))
	if !ok {
		return fmt.Errorf("unknown terrain %s", string(b))
	}
	return nil
}

func unmarshalCode(s string) (Terrain, bool) {
	if c, ok := slices.BinarySearch(Codes, s); ok {
		return Terrain(c), true
	}
	switch s {
	case "CONIFER HILLS":
		return CH, true
	case "FORD":
		return FORDS, true
	case "GRASSY HILLS":
		return GH, true
	case "OCEAN":
		return O, true
	case "PRAIRIE":
		return PR, true
	case "RIVER":
		return R, true
	case "ROCKY HILLS":
		return RH, true
	}
	return Unknown, true
}
