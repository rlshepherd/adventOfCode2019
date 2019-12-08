package main

import (
	"fmt"
	"testing"
)

var h1 = LineSegment{
	Point{0, 0},
	Point{10, 0},
	false,
}
var h2 = LineSegment{
	Point{20, 0},
	Point{10, 0},
	false,
}
var v1 = LineSegment{
	Point{5, -3},
	Point{5, 3},
	false,
}
var v2 = LineSegment{
	Point{15, -3},
	Point{15, 3},
	false,
}
var v3 = LineSegment{
	Point{5, 3},
	Point{5, 10},
	false,
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
		want Point
	}{
		{h1, v1, Point{5, 0}},
		{h2, v2, Point{15, 0}},
	}
	for _, c := range cases {
		got := newIntersection(c.h, c.v)
		if got != c.want {
			t.Errorf(fmt.Sprintf("newIntersection(%v, %v) = %d, want %d", c.h, c.v, got, c.want))
		}
	}
}

func TestFindIntersection(t *testing.T) {
	cases := []struct {
		h    Path
		v    Path
		want Points
	}{
		{hPath, vPath, Points{Point{5, 0}, Point{15, 0}}},
	}
	for _, c := range cases {
		got := FindIntersections(c.h, c.v)
		if !got.Equal(c.want) {
			t.Errorf(fmt.Sprintf("FindIntersection(%+v, %+v) = %v, want %v", c.h, c.v, got, c.want))
		}
	}
}
