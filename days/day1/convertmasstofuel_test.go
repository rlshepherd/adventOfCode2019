package main

import "testing"

func TestConvertMassToFuel(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range cases {
		got := ConvertMassToFuel(c.in)
		if got != c.want {
			t.Errorf("ConvertMassToFuel(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
