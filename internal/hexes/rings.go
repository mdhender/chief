// chief - a TribeNet player aid
// Copyright (c) 2022-2023 Michael D Henderson. All rights reserved.

package hexes

func cube_scale(c Cube, factor int) Cube {
	return Cube{Q: c.Q * factor, R: c.R * factor, S: c.S * factor}
}

// this code doesn't work for radius == 0; can you see why?
func cube_ring(center Cube, radius int) []Cube {
	var results []Cube
	var hex = cube_add(center, cube_scale(cube_direction(4), radius))
	for i := 0; i < 6; i++ {
		for j := 0; j < radius; j++ {
			results = append(results, hex)
			hex = cube_neighbor(hex, i)
		}
	}
	return results
}

// Area returns number of cubes in the radius
func Area(radius int) int {
	if radius < 0 {
		return 0
	}
	return 1 + 3*radius*(radius+1)
}
