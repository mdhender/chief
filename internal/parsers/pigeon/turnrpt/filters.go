// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package parser

import "bytes"

// FilterDefaultGrid searches the provided input []byte for all occurrences of ' ## '
// and replaces each occurrence with the defaultHex string surrounded by spaces (' ').
//
// For instance, if you provide a grid of "DA", the function will replace all ' ## '
// strings in the input with ' DA '.
//
// Parameters:
//   - input ([]byte): The input byte slice in which replacements are to be made.
//   - grid (string): A string to replace ' ## ' occurrences in the input.
//     Note: it should be exactly 2 characters long.
//
// Returns:
//   - ([]byte): The original slice with every occurrence of ' ## ' replaced by grid.
func FilterDefaultGrid(input []byte, grid string) []byte {
	return bytes.ReplaceAll(input, []byte{' ', '#', '#', ' '}, []byte{' ', grid[0], grid[1], ' '})
}
