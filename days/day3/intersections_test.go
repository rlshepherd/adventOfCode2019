package main

import (
	"fmt"
	"testing"
)

var h1 = LineSegment{
	Point{0, 0},
	Point{10, 0},
	false,
	10,
}
var h2 = LineSegment{
	Point{20, 0},
	Point{10, 0},
	false,
	20,
}
var v1 = LineSegment{
	Point{5, -3},
	Point{5, 3},
	true,
	6,
}
var v2 = LineSegment{
	Point{15, -3},
	Point{15, 3},
	true,
	6,
}
var v3 = LineSegment{
	Point{5, 3},
	Point{5, 10},
	true,
	7,
}

var hPath = Path{h1, h2}
var vPath = Path{v1, v2}

func TestIntersection(t *testing.T) {
	cases := []struct {
		h    LineSegment
		v    LineSegment
		want bool
	}{
		{h1, v1, true},
		{h1, v2, false},
		{h1, v3, false},
		{h2, v1, false},
		{h2, v2, true},
		{h2, v3, false},
	}
	for _, c := range cases {
		got := Intersects(c.h, c.v)
		if got != c.want {
			t.Errorf(fmt.Sprintf("Intersection(%v, %v) = %t, want %t", c.v, c.h, got, c.want))
		}
	}
}

func TestNewIntersection(t *testing.T) {
	cases := []struct {
		h    LineSegment
		v    LineSegment
		want Intersection
	}{
		{h1, v1, Intersection{Point{5, 0}, 5}},
		{h2, v2, Intersection{Point{15, 0}, 5}},
	}
	for _, c := range cases {
		got := newIntersection(c.h, c.v, 0)
		if got != c.want {
			t.Errorf(fmt.Sprintf("newIntersection(%v, %v) = %d, want %d", c.h, c.v, got, c.want))
		}
	}
}

var inters = Intersections{
	Intersection{
		Point{5, 0},
		3,
	},
	Intersection{
		Point{15, 0},
		3,
	},
}

func TestFindIntersection(t *testing.T) {
	cases := []struct {
		h    Path
		v    Path
		want Intersections
	}{
		{hPath, vPath, inters},
	}
	for _, c := range cases {
		got := FindIntersections(c.h, c.v)
		if !got.Equal(c.want) {
			t.Errorf(fmt.Sprintf("FindIntersection(%+v, %+v) = %v, want %v", c.h, c.v, got, c.want))
		}
	}
}
