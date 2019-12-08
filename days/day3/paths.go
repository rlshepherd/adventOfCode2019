package main

import (
	"math"
	"strconv"
)

// Point has an x and y.
type Point struct {
	x, y int
}

func (a Point) manhattanDistance() int {
	return int(math.Abs(float64(a.x))) + int(math.Abs(float64(a.y)))
}

// Points implements the sorting interface for point
// Points uses Manhattan distance from zero to implement less.
type Points []Point

func (a Points) Len() int           { return len(a) }
func (a Points) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Points) Less(i, j int) bool { return a[i].manhattanDistance() < a[j].manhattanDistance() }

// Equal is a method for Points.
func (a Points) Equal(b Points) bool {
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

// A LineSegment contains start and end points, as well as orientation.
type LineSegment struct {
	start    Point
	end      Point
	vertical bool
}

// NewLineSegment accepts a Point and an instruction string in the format "U81", and returns a LineSegment
func NewLineSegment(s Point, i string) (LineSegment, error) {
	v := string(i[0])
	n, err := strconv.Atoi(string(i[1:len(i)]))
	if err != nil {
		println("error:", err)
	}
	p := new(Point)

	switch {
	case v == "U":
		p = &Point{s.x, s.y + n}
	case v == "D":
		p = &Point{s.x, s.y - n}
	case v == "R":
		p = &Point{s.x + n, s.y}
	case v == "L":
		p = &Point{s.x - n, s.y}
	default:
		p = &Point{s.x - n, s.y}
	}
	return LineSegment{s, *p, s.x == p.x}, nil
}

// Path implements the sort interface for a slice of LineSegments.
type Path []LineSegment

func (a Path) Len() int           { return len(a) }
func (a Path) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Path) Less(i, j int) bool { return a[i].start.x < a[j].start.x }

// Equal accepts two Paths, and returns true is they are equal, otherwise it returns false.
func (a Path) Equal(b Path) bool {
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

// NewPathsByOrientation accepts 2d array of strings and returns two sets of paths:
// a set of hortizontal paths (h), and a set of vertical paths (v).
// to access all line segments for a given path path i, use h[i] ad v[i].
func NewPathsByOrientation(instructions [][]string) ([]Path, []Path) {
	h := make([]Path, len(instructions))
	v := make([]Path, len(instructions))

	for j, k := range instructions {
		h[j] = make(Path, 0)
		v[j] = make(Path, 0)

		p := Point{0, 0}
		for _, i := range k {
			s, err := NewLineSegment(p, i)
			if err != nil {
				println("error!", err)
			}
			if s.vertical {
				v[j] = append(v[j], s)
			} else {
				h[j] = append(h[j], s)
			}
			p = s.end
		}
	}
	return h, v
}
