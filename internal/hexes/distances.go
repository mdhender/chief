// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

func cube_subtract(a, b Cube) Cube {
	return Cube{Q: a.Q - b.Q, R: a.R - b.R, S: a.S - b.S}
}

func cube_distance(a, b Cube) int {
	var vec = cube_subtract(a, b)
	return (abs(vec.Q) + abs(vec.R) + abs(vec.S)) / 2
}
