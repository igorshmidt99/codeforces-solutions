package tests

import (
	"solutions/codeforces/sol"
	"testing"
)

func TestAft(t *testing.T) {
	testCases := []struct {
		n   int
		m   int
		ss  []string
		res int
	}{
		{3, 1, []string{"a", "b", "c"}, 1},
		{2, 9, []string{"alpha", "beta"}, 2},
		{4, 12, []string{"hello", "world", "and", "codeforces"}, 2},
		{3, 2, []string{"ab", "c", "d"}, 1},
		{3, 2, []string{"abc", "ab", "c"}, 0},
	}

	for i, tCase := range testCases {
		res := sol.Aft(tCase.n, tCase.m, tCase.ss)
		if res != tCase.res {
			t.Fatalf("Test case %d fail", i)
		}
	}

}
