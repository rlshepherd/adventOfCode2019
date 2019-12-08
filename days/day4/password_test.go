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
		{1111, true},
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
		{11122, true},
		{1112, false},
		{12344, true},
		{123444, false},
	}
	for _, c := range cases {
		a := intToSlice(c.in, []int{})
		got := hasAtleastOne2Pair(a)
		if got != c.want {
			t.Errorf(fmt.Sprintf("hasPair(%d) = %t, want %t", c.in, got, c.want))
		}
	}
}

func TestSliceToInt(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 123},
		{[]int{2, 1, 5, 5}, 2155},
	}
	for _, c := range cases {
		got, err := sliceToInt(c.in)
		if err != nil {
			println("err")
		}
		if got != c.want {
			t.Errorf(fmt.Sprintf("sliceToInt(%d) = %d, want %d", c.in, got, c.want))
		}
	}
}

func TestFindPasswords(t *testing.T) {
	cases := []struct {
		low  int
		high int
		want int
	}{
		{10, 12, 1},
	}
	for _, c := range cases {
		passwords := findPasswords(c.low, c.high)
		got := len(passwords)
		if got != c.want {
			t.Errorf(fmt.Sprintf("findPasswords(%d, %d) = %d, want %d. passwords found: %v", c.low, c.high, got, c.want, passwords))
		}
	}
}
