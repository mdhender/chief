// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package terrain

import (
	"bytes"
	"fmt"
)

type Terrain int

const (
	Clear Terrain = iota
	Delta
	Desert
	Forest
	Gravel
	Ice
	Mountain
	Ocean
	Plain
	Rock
	Rough
	SacredMountain
	SaltMarsh
	Sea
	Steppe
	Swamp
)

// ToFill returns the background fill for a terrain type.
/*
#Delta {color: darkolivegreen; }
#Desert {color: khaki ;}
#Forest {color: forestgreen ;}
#Gravel {color: darkgray ;}
#Ice {color: whitesmoke ;}
#Mountain {color: sienna ;}
#Ocean {color: deepskyblue ;}
#Plain {color: wheat ;}
#Rock {color: gainsboro ;}
#Rough {color: lightgray ;}
#SacredMountain {color: gold ;}
#SaltMarsh {color: yellowgreen ;}
#Sea {color: lightblue ;}
#Steppe {color: linen ;}
#Swamp {color: mediumseagreen ;}
*/
func (t Terrain) ToFill() string {
	switch t {
	case Clear:
		return "hsl(53, 100%, 94%)"
	case Delta:
		return "darkolivegreen"
	case Desert:
		return "khaki"
	case Forest:
		return "forestgreen"
	case Gravel:
		return "darkgray"
	case Ice:
		return "black" // "whitesmoke"
	case Mountain:
		return "chocolate"
	case Ocean:
		return "deepskyblue" //"hsl(197, 78%, 85%)"
	case Plain:
		return "wheat"
	case Rock:
		return "gainsboro"
	case Rough:
		return "lightgray"
	case SacredMountain:
		return "gold"
	case SaltMarsh:
		return "yellowgreen"
	case Sea:
		return "lightblue"
	case Steppe:
		return "linen"
	case Swamp:
		return "mediumseagreen"
	}
	panic(fmt.Sprintf("assert(t != %d)", t))
}

func (t Terrain) MarshalJSON() ([]byte, error) {
	switch t {
	case Clear:
		return []byte(`"clear"`), nil
	case Delta:
		return []byte(`"delta"`), nil
	case Desert:
		return []byte(`"desert"`), nil
	case Forest:
		return []byte(`"forest"`), nil
	case Gravel:
		return []byte(`"gravel"`), nil
	case Ice:
		return []byte(`"ice"`), nil
	case Mountain:
		return []byte(`"mountain"`), nil
	case Ocean:
		return []byte(`"ocean"`), nil
	case Plain:
		return []byte(`"plain"`), nil
	case Rock:
		return []byte(`"rock"`), nil
	case Rough:
		return []byte(`"rough"`), nil
	case SacredMountain:
		return []byte(`"sacred-mountain"`), nil
	case SaltMarsh:
		return []byte(`"salt-marsh"`), nil
	case Sea:
		return []byte(`"sea"`), nil
	case Steppe:
		return []byte(`"steppe"`), nil
	case Swamp:
		return []byte(`"swamp"`), nil
	}
	return nil, fmt.Errorf("invalid value")
}

func (t *Terrain) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`"clear"`)) {
		*t = Clear
		return nil
	} else if bytes.Equal(b, []byte(`"delta"`)) {
		*t = Delta
		return nil
	} else if bytes.Equal(b, []byte(`"desert"`)) {
		*t = Desert
		return nil
	} else if bytes.Equal(b, []byte(`"forest"`)) {
		*t = Forest
		return nil
	} else if bytes.Equal(b, []byte(`"gravel"`)) {
		*t = Gravel
		return nil
	} else if bytes.Equal(b, []byte(`"ice"`)) {
		*t = Ice
		return nil
	} else if bytes.Equal(b, []byte(`"mountain"`)) {
		*t = Mountain
		return nil
	} else if bytes.Equal(b, []byte(`"ocean"`)) {
		*t = Ocean
		return nil
	} else if bytes.Equal(b, []byte(`"plain"`)) {
		*t = Plain
		return nil
	} else if bytes.Equal(b, []byte(`"rock"`)) {
		*t = Rock
		return nil
	} else if bytes.Equal(b, []byte(`"rough"`)) {
		*t = Rough
		return nil
	} else if bytes.Equal(b, []byte(`"sacred-mountain"`)) {
		*t = SacredMountain
		return nil
	} else if bytes.Equal(b, []byte(`"salt-marsh"`)) {
		*t = SaltMarsh
		return nil
	} else if bytes.Equal(b, []byte(`"sea"`)) {
		*t = Sea
		return nil
	} else if bytes.Equal(b, []byte(`"steppe"`)) {
		*t = Steppe
		return nil
	} else if bytes.Equal(b, []byte(`"swamp"`)) {
		*t = Swamp
		return nil
	}
	*t = Clear
	return fmt.Errorf("invalid terrain: %q", string(b))
}

func (t Terrain) String() string {
	switch t {
	case Clear:
		return "clear"
	case Delta:
		return "delta"
	case Desert:
		return "desert"
	case Forest:
		return "forest"
	case Gravel:
		return "gravel"
	case Ice:
		return "ice"
	case Mountain:
		return "mountain"
	case Ocean:
		return "ocean"
	case Plain:
		return "plain"
	case Rock:
		return "rock"
	case Rough:
		return "rough"
	case SacredMountain:
		return "sacred-mountain"
	case SaltMarsh:
		return "salt-marsh"
	case Sea:
		return "sea"
	case Steppe:
		return "steppe"
	case Swamp:
		return "swamp"
	}
	return "clear"
}
