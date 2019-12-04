package main

//LinesOfInstructions are read from a file.
type LinesOfInstructions [][]string

// Point has an x and y
type Point struct {
	x, y int
}

// LineSegment contains a start and an end point
type LineSegment struct {
	start    Point
	end      Point
	vertical bool
}

// Path is a slice of line segments.
type Path []LineSegment

func newLineSegment(point Point, instruction string) (segment LineSegment) {
	return
}

func main() {

}
