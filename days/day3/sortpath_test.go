package main

import (
	"fmt"
	"sort"
	"testing"
)

var test1 = Path{
	{Point{10, 10}, Point{20, 10}, true},
	{Point{1, 1}, Point{2, 1}, true},
}
var test2 = Path{
	{Point{10, 10}, Point{20, 10}, true},
	{Point{1, 1}, Point{2, 1}, true},
}

var test3 = Path{
	{Point{1, 1}, Point{2, 1}, true},
	{Point{10, 10}, Point{20, 10}, true},
}

var test4 = Path{
	{Point{20, 30}, Point{30, 30}, true},
	{Point{10, 10}, Point{20, 10}, true},
	{Point{1, 1}, Point{2, 1}, true},
}

var test5 = Path{
	{Point{1, 1}, Point{2, 1}, true},
	{Point{10, 10}, Point{20, 10}, true},
	{Point{20, 30}, Point{30, 30}, true},
}

func EqualPaths(a Path, b Path) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestEqualPath(t *testing.T) {
	cases := []struct {
		a    Path
		b    Path
		want bool
	}{
		{test1, test2, true},
		{test1, test3, false},
	}
	for _, c := range cases {
		if EqualPaths(c.a, c.b) != c.want {
			t.Errorf(fmt.Sprintf("EqualPaths(%v,%v) = %t, want %t", c.a, c.b, EqualPaths(test1, test3), c.want))
		}
	}
}

func TestSortPath(t *testing.T) {
	cases := []struct {
		in   Path
		want Path
	}{
		{test1, test3},
		{test4, test5},
	}
	for _, c := range cases {
		sort.Sort(c.in)
		if !EqualPaths(c.in, c.want) {
			t.Errorf(fmt.Sprintf("Sorted = %v, want %v", c.in, c.want))
		}
	}
}
