// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package tiles

import (
	"testing"
)

func TestModulo(t *testing.T) {
	for _, tc := range []struct {
		id     int
		x, n   int
		expect int
	}{
		{1, 0, 6, 0},
		{2, 1, 6, 1},
		{3, 2, 6, 2},
		{4, 3, 6, 3},
		{5, 4, 6, 4},
		{6, 5, 6, 5},
		{7, 6, 6, 0},
		{8, 15, 6, 3},
		{9, 24, 6, 0},
		{10, -1, 6, 5},
		{11, -2, 6, 4},
		{12, -3, 6, 3},
		{13, -4, 6, 2},
		{14, -5, 6, 1},
		{15, -6, 6, 0},
		{16, -15, 6, 3},
		{17, -24, 6, 0},
	} {
		if got := modulo(tc.x, tc.n); got != tc.expect {
			t.Errorf("%d: %d modulo %d: want %d, got %d\n", tc.id, tc.x, tc.n, tc.expect, got)
		}
	}
}
