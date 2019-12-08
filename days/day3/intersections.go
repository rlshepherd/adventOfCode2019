package main

// Intersects accepts a horizontal and vertical line segment and returns if they intersect
func Intersects(h LineSegment, v LineSegment) bool {
	hOverlap := (h.start.x <= v.start.x && h.end.x >= v.start.x) || (h.end.x <= v.start.x && h.start.x >= v.start.x)
	vOverlap := (v.start.y <= h.start.y && v.end.y >= h.start.y) || (v.end.y <= h.start.y && v.start.y >= h.start.y)
	return hOverlap && vOverlap
}

func newIntersection(h LineSegment, v LineSegment) Point {
	return Point{v.start.x, h.start.y}
}

// FindIntersections accepts a set of sorted horizontal and vertical line segments,
// and returns a list of intersection points.
func FindIntersections(h Path, v Path) Points {
	p := make(Points, 0)
	for _, hl := range h {
		for _, vl := range v {
			if Intersects(hl, vl) {
				if !(hl.start.y == 0 && vl.start.x == 0) { // ignore origin
					p = append(p, newIntersection(hl, vl))
				}
			}
		}
	}
	return p

	// TODO implement the Bentley Ottmann algorithm
	// requires implementing a self-balancing BST in go first.
}
