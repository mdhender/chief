// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

// abs is a helper function to get the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// mod is a helper function to get the modulus of an integer
// (as opposed to %, which is the remainder operator)
func mod(a, b int) int {
	// you can check for b == 0 separately and do what you want
	if b < 0 {
		return -mod(-a, -b)
	}
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

// hex_direction converts 0...5 to a hex offset
func hex_direction(direction int) Hex {
	// direction = mod(direction, 6)
	direction = (6 + (direction % 6)) % 6
	return hex_directions[direction]
}
