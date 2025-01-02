package tests

import (
	"solutions/codeforces/sol"
	"testing"
)

func TestB(t *testing.T) {
	cases := []struct {
		n   int
		a   []int
		res string
	}{
		{3, []int{3, 2, 1}, "YES"},
		{3, []int{1, 1, 3}, "NO"},
		{4, []int{1, 2, 5, 4}, "YES"},
		{4, []int{1, 6, 6, 1}, "NO"},
		{5, []int{6, 2, 1, 4, 2}, "YES"},
		{4, []int{1, 4, 2, 1}, "NO"},
		{5, []int{3, 1, 2, 1, 3}, "NO"},
		{3, []int{2, 4, 2}, "NO"},
		// {5, []int{1, 4, 2, 2, 1}, "YES"},
	}

	for i, c := range cases {
		if sol.B(c.n, c.a) != c.res {
			t.Fatalf("Test № %d with %v has failed!", i, c)
		}
		t.Logf("Test № %d with values %v has passed\n", i, c)
	}
}
