// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import "fmt"

// abs returns the absolute value of an integer.
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// modulo is not the remainder ("%") operator!
func modulo(x, n int) int {
	return (x%n + n) % n
}

// gxyScale scales TribeNet's ## XXYY coordinates using 21 rows and 30 columns per grid.
// Panics if the caller passes in '#' for the grid.
func gxyScale(gxy string) (x, y int) {
	const rowsPerGrid, columnsPerGrid = 21, 30
	gridY, gridX := int(gxy[0]), int(gxy[1])
	if !('A' <= gridX && gridX <= 'Z' && 'A' <= gridY && gridY <= 'Z') {
		panic("assert(grid.xy is alpha)")
	}
	gridX, gridY = gridX-'A', gridY-'A'
	x = gridX*columnsPerGrid + (int(gxy[3])-'0')*10 + (int(gxy[4]) - '0')
	y = gridY*rowsPerGrid + (int(gxy[5])-'0')*10 + (int(gxy[6]) - '0')
	return x, y
}

// xlatToGXY translates a scaled X, Y to TribeNet's ## XXYY  using 21 rows and 30 columns per grid.
func xlatToGXY(x, y int) string {
	const rowsPerGrid, columnsPerGrid = 21, 30
	gridX, gridY := x/columnsPerGrid, y/rowsPerGrid
	if !(0 <= gridX && gridX <= 25 && 0 <= gridY && gridY <= 25) {
		panic("assert(grid.xy is out of range)")
	}
	x, y = x-gridX*columnsPerGrid, y-gridY*rowsPerGrid
	return fmt.Sprintf("%c%c %02d%02d", byte(gridY)+'A', byte(gridX)+'A', x, y)
}
