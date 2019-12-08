package main

import (
	"fmt"
	"sort"
	"testing"
)

var test1 = Path{
	{Point{10, 10}, Point{20, 10}, true, 10},
	{Point{1, 1}, Point{2, 1}, true, 11},
}
var test2 = Path{
	{Point{10, 10}, Point{20, 10}, true, 10},
	{Point{1, 1}, Point{2, 1}, true, 10},
}

var test3 = Path{
	{Point{1, 1}, Point{2, 1}, true, 10},
	{Point{10, 10}, Point{20, 10}, true, 10},
}

var test4 = Path{
	{Point{20, 30}, Point{30, 30}, true, 10},
	{Point{10, 10}, Point{20, 10}, true, 10},
	{Point{1, 1}, Point{2, 1}, true, 10},
}

var test5 = Path{
	{Point{1, 1}, Point{2, 1}, true, 10},
	{Point{10, 10}, Point{20, 10}, true, 10},
	{Point{20, 30}, Point{30, 30}, true, 10},
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
		if c.a.Equal(c.b) != c.want {
			t.Errorf(fmt.Sprintf("EqualPaths(%v,%v) = %t, want %t", c.a, c.b, test1.Equal(test3), c.want))
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
		if !c.in.Equal(c.want) {
			t.Errorf(fmt.Sprintf("Sorted = %v, want %v", c.in, c.want))
		}
	}
}
