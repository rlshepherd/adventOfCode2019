package main

import "math"

// Intersection is a point with an added distance property
type Intersection struct {
	p     Point
	steps int
}

// Intersections implements the sort interface for intersection
type Intersections []Intersection

func (a Intersections) Len() int      { return len(a) }
func (a Intersections) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Intersections) Less(i, j int) bool {
	return a[i].steps < a[j].steps
}

// Equal compares two Intersections (slices of intersection).
func (a Intersections) Equal(b Intersections) bool {
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

// Intersects accepts a horizontal and vertical line segment and returns if they intersect
func Intersects(h LineSegment, v LineSegment) bool {
	hOverlap := (h.start.x <= v.start.x && h.end.x >= v.start.x) || (h.end.x <= v.start.x && h.start.x >= v.start.x)
	vOverlap := (v.start.y <= h.start.y && v.end.y >= h.start.y) || (v.end.y <= h.start.y && v.start.y >= h.start.y)
	return hOverlap && vOverlap
}

func newIntersection(h LineSegment, v LineSegment, d int) Intersection {
	return Intersection{Point{v.start.x, h.start.y}, d}
}

// FindIntersections accepts a set of sorted horizontal and vertical line segments,
// and returns a list of intersection points.
func FindIntersections(h Path, v Path) Intersections {
	p := make(Intersections, 0)
	for _, hl := range h {
		for _, vl := range v {
			if Intersects(hl, vl) {
				if !(hl.start.y == 0 && vl.start.x == 0) {
					d := (hl.steps - int(math.Abs(float64(hl.end.x-vl.end.x)))) + (vl.steps - int(math.Abs(float64(vl.end.y-hl.end.y))))
					p = append(p, newIntersection(hl, vl, d))

				}
			}
		}
	}
	return p

	// TODO implement the Bentley Ottmann algorithm
	// requires implementing a self-balancing BST in go first.
}
