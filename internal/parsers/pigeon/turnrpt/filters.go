// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package parser

import (
	"bytes"
	"regexp"
)

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

// TransformMarkScoutLines modifies a slice of bytes to mark specific lines.
//
// It takes as input a slice of bytes which it scans line by line. A line is identified as
// a scouting action if it matches the regular expression pattern "^Scout \d:Scout", where
// \d is a digit. For every line identified as a scouting action, "$$$" is appended at the
// end of the line.
//
// The resulting line transformations are joined back together to form a new slice of bytes
// which is then returned.
//
// Example:
// Suppose we have the following input:
//
//	[]byte("Scout 1:Scout is a programming assistant\nAnother line\nScout again")
//
// The output will be:
//
//	[]byte("Scout 1:Scout is a programming assistant$$$\nAnother line\nScout again")
//
// Parameters:
//
//	input ([]byte): The original byte slice to be transformed.
//
// Returns:
//
//	[]byte: The transformed byte slice with "$$$" appended to scouting action lines.
func TransformMarkScoutLines(input []byte) []byte {
	// scouting actions look like "Scout #:Scout", where # is a single digit.
	reScoutAction := regexp.MustCompile(`^Scout \d:Scout`)

	lines := bytes.Split(input, []byte{'\n'})
	for i, line := range lines {
		if reScoutAction.Match(line) {
			lines[i] = append(lines[i], '$', '$', '$')
		}
	}

	return bytes.Join(lines, []byte{'\n'})
}
