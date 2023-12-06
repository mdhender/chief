// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"strconv"
	"strings"
)

func cellToBool(row []string, n int) (bool, error) {
	val, _ := cellToString(row, n)
	return val == "Y", nil
}

func cellToFloat(row []string, n int) (float64, error) {
	if val, _ := cellToString(row, n); val != "" {
		return strconv.ParseFloat(val, 64)
	}
	return 0, nil
}

func cellToInt(row []string, n int) (int, error) {
	if val, _ := cellToString(row, n); val != "" {
		return strconv.Atoi(val)
	}
	return 0, nil
}

func cellToString(row []string, n int) (string, error) {
	if n < len(row) {
		return strings.TrimSpace(row[n]), nil
	}
	return "", nil
}
