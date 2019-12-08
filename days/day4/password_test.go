package main

import (
	"fmt"
	"testing"
)

func TestAscending(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1234, true},
		{4321, false},
		{123450, false},
	}
	for _, c := range cases {
		a := intToSlice(c.in, []int{})
		got := ascending(a)
		if got != c.want {
			t.Errorf(fmt.Sprintf("ascending(%d) = %t, want %t", c.in, got, c.want))
		}
	}
}

func TestHasPair(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1234, false},
		{4321, false},
		{211, true},
		{2121, false},
		{111111, true},
	}
	for _, c := range cases {
		a := intToSlice(c.in, []int{})
		got := hasPair(a)
		if got != c.want {
			t.Errorf(fmt.Sprintf("hasPair(%d) = %t, want %t", c.in, got, c.want))
		}
	}
}
