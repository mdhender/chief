// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package terrain

import (
	"bytes"
	"fmt"
)

type CODE int

const (
	NONE CODE = iota
	ALPS
	AR
	BH
	BR
	CH
	DE
	DF
	DH
	Fords
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
)

// MarshalJSON implements the json.Marshaler interface
func (c CODE) MarshalJSON() ([]byte, error) {
	switch c {
	case NONE:
		return []byte("NONE"), nil
	case ALPS:
		return []byte("ALPS"), nil
	case AR:
		return []byte("AR"), nil
	case BH:
		return []byte("BH"), nil
	case BR:
		return []byte("BR"), nil
	case CH:
		return []byte("CH"), nil
	case DE:
		return []byte("DE"), nil
	case DF:
		return []byte("DF"), nil
	case DH:
		return []byte("DH"), nil
	case Fords:
		return []byte("Fords"), nil
	case GH:
		return []byte("GH"), nil
	case HSM:
		return []byte("HSM"), nil
	case JG:
		return []byte("JG"), nil
	case JH:
		return []byte("JH"), nil
	case L:
		return []byte("L"), nil
	case LCM:
		return []byte("LCM"), nil
	case LJM:
		return []byte("LJM"), nil
	case LSM:
		return []byte("LSM"), nil
	case O:
		return []byte("O"), nil
	case PI:
		return []byte("PI"), nil
	case PR:
		return []byte("PR"), nil
	case R:
		return []byte("R"), nil
	case RH:
		return []byte("RH"), nil
	case SH:
		return []byte("SH"), nil
	case SW:
		return []byte("SW"), nil
	case TU:
		return []byte("TU"), nil
	}
	return nil, fmt.Errorf("invalid terrain code")
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *CODE) UnmarshalJSON(b []byte) error {
	*c = bytesToCode(b)
	if *c == NONE {
		return fmt.Errorf("unknown code")
	}
	return nil
}

// Description returns a description of the code.
func (c CODE) Description() string {
	switch c {
	case NONE:
		return "unknown terrain"
	case ALPS:
		return "a bigger version of HSM"
	case AR:
		return "Arid: tundra without water."
	case BH:
		return "Brush Hill: Hill covered with brush."
	case BR:
		return "Brush: Conifer forest with fewer trees more bushes (Forestry not possible here)."
	case CH:
		return "Conifer Hill: Hill covered with conifer forest."
	case DE:
		return "Desert: Arid without grass."
	case DF:
		return "Deciduous: Seasonal forest."
	case DH:
		return "Deciduous Hill: Hill covered with deciduous forest."
	case Fords:
		return "Shallow spots that are ways across rivers."
	case GH:
		return "Grassy Hill: Hill covered with grass."
	case HSM:
		return "High Mountains: cannot be entered, except through a pass."
	case JG:
		return "Jungle: Wet forest."
	case JH:
		return "Jungle Hill: Hill covered with jungle."
	case L:
		return "Lake a body of water"
	case LCM:
		return "Low Conifer Mountains: Hills but higher, difficult to enter covered with conifer forest."
	case LJM:
		return "Low Jungle Mountain"
	case LSM:
		return "Low Snowy Mountains: Hills but higher, very difficult to enter."
	case O:
		return "Ocean"
	case PI:
		return "Polar Ice: Permanent ice and difficult to move through."
	case PR:
		return "Prairie: Grassland."
	case R:
		return "Rivers: Large moving bodies of water impossible to cross unless through a ford or by boat."
	case RH:
		return "Rocky Hill: Hill covered with rocks."
	case SH:
		return "Snow Hill: colder than GH, snow rather than grass."
	case SW:
		return "Swamp: very wet grassland."
	case TU:
		return "Tundra: not very good grassland."
	}
	panic(fmt.Sprintf("assert(code != %d)", c))
}

// String implements the string.Stringer interface
func (c CODE) String() string {
	switch c {
	case NONE:
		return "NONE"
	case ALPS:
		return "ALPS"
	case AR:
		return "AR"
	case BH:
		return "BH"
	case BR:
		return "BR"
	case CH:
		return "CH"
	case DE:
		return "DE"
	case DF:
		return "DF"
	case DH:
		return "DH"
	case Fords:
		return "Fords"
	case GH:
		return "GH"
	case HSM:
		return "HSM"
	case JG:
		return "JG"
	case JH:
		return "JH"
	case L:
		return "L"
	case LCM:
		return "LCM"
	case LJM:
		return "LJM"
	case LSM:
		return "LSM"
	case O:
		return "O"
	case PI:
		return "PI"
	case PR:
		return "PR"
	case R:
		return "R"
	case RH:
		return "RH"
	case SH:
		return "SH"
	case SW:
		return "SW"
	case TU:
		return "TU"
	}
	panic(fmt.Sprintf("assert(code != %d)", c))
}

func bytesToCode(b []byte) CODE {
	if bytes.Equal(b, []byte("ALPS")) {
		return ALPS
	} else if bytes.Equal(b, []byte("AR")) {
		return AR
	} else if bytes.Equal(b, []byte("BH")) {
		return BH
	} else if bytes.Equal(b, []byte("BR")) {
		return BR
	} else if bytes.Equal(b, []byte("CH")) || bytes.Equal(b, []byte("CONIFER HILLS")) {
		return CH
	} else if bytes.Equal(b, []byte("DE")) {
		return DE
	} else if bytes.Equal(b, []byte("DF")) {
		return DF
	} else if bytes.Equal(b, []byte("DH")) {
		return DH
	} else if bytes.Equal(b, []byte("Fords")) {
		return Fords
	} else if bytes.Equal(b, []byte("GH")) {
		return GH
	} else if bytes.Equal(b, []byte("HSM")) {
		return HSM
	} else if bytes.Equal(b, []byte("JG")) {
		return JG
	} else if bytes.Equal(b, []byte("JH")) {
		return JH
	} else if bytes.Equal(b, []byte("L")) {
		return L
	} else if bytes.Equal(b, []byte("LCM")) {
		return LCM
	} else if bytes.Equal(b, []byte("LJM")) {
		return LJM
	} else if bytes.Equal(b, []byte("LSM")) {
		return LSM
	} else if bytes.Equal(b, []byte("O")) {
		return O
	} else if bytes.Equal(b, []byte("PI")) {
		return PI
	} else if bytes.Equal(b, []byte("PR")) || bytes.Equal(b, []byte("PRAIRIE")) {
		return PR
	} else if bytes.Equal(b, []byte("R")) {
		return R
	} else if bytes.Equal(b, []byte("RH")) {
		return RH
	} else if bytes.Equal(b, []byte("SH")) {
		return SH
	} else if bytes.Equal(b, []byte("SW")) {
		return SW
	} else if bytes.Equal(b, []byte("TU")) {
		return TU
	}
	return NONE
}
