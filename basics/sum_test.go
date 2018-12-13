package basics

import "testing"

func TestSum(t *testing.T) {
	a, b, exp := 2, 3, 5

	got := Sum(a, b)

	if got != exp {
		t.Fail()
	}
}

func TestSum_Table(t *testing.T) {
	var cases = []struct {
		a   int
		b   int
		exp int
	}{
		{0, 0, 0},
		{2, 3, 5},
		{2, -3, -1},
	}
	for _, v := range cases {
		got := Sum(v.a, v.b)
		if got != v.exp {
			t.Fail()
		}
	}
}
