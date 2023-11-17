// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

var cube_direction_vectors = []Cube{
	Cube{Q: +1, R: 0, S: -1},
	Cube{Q: +1, R: -1, S: 0},
	Cube{Q: 0, R: -1, S: +1},
	Cube{Q: -1, R: 0, S: +1},
	Cube{Q: -1, R: +1, S: 0},
	Cube{Q: 0, R: +1, S: -1},
}

func cube_direction(direction int) Cube {
	direction = (6 + (direction % 6)) % 6 // mod(direction, 6)
	return cube_direction_vectors[direction]
}

func cube_add(c Cube, vec Cube) Cube {
	return Cube{c.Q + vec.Q, c.R + vec.R, c.S + vec.S}
}

func cube_neighbor(c Cube, direction int) Cube {
	return cube_add(c, cube_direction(direction))
}

var cube_diagonal_vectors = []Cube{
	Cube{Q: +2, R: -1, S: -1},
	Cube{Q: +1, R: -2, S: +1},
	Cube{Q: -1, R: -1, S: +2},
	Cube{Q: -2, R: +1, S: +1},
	Cube{Q: -1, R: +2, S: -1},
	Cube{Q: +1, R: +1, S: -2},
}

func cube_diagonal_neighbor(c Cube, direction int) Cube {
	return cube_add(c, cube_diagonal_vectors[direction])
}
