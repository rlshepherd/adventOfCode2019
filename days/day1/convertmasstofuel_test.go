package main

import (
	"fmt"
	"testing"
)

func TestConvertMassToFuel(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, c := range cases {
		got := ConvertMassToFuel(c.in)
		if got != c.want {
			t.Errorf(fmt.Sprintf("ConvertMassToFuel(%d) == %d, want %d", c.in, got, c.want))
		}
	}
}
