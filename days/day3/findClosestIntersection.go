package main

import (
	"sort"
	"strconv"

	fileutils "github.com/rlshepherd/adventOfCode2019/utils"
)

// The LinesOfInstructions are returned by ParseInstructionFile
type LinesOfInstructions [][]string

// Point has an x and y
type Point struct {
	x, y int
}

// A LineSegment contains start and end points, as well as orientation.
type LineSegment struct {
	start    Point
	end      Point
	vertical bool
}

// Path implements the sort interface a slice of LineSegments.
type Path []LineSegment

func (a Path) Len() int           { return len(a) }
func (a Path) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Path) Less(i, j int) bool { return a[i].start.x < a[j].start.x }

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

// TracePaths accepts LinesOFInstructions, and returns a set of horizontal and vertical paths.
func TracePaths(instructions LinesOfInstructions) ([]Path, []Path) {
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
			p = s.start
		}
	}
	return h, v
}

func main() {

	// Read instructions
	instructions := fileutils.ParseIntrusctionFile("input.txt")

	// Create empty paths
	hPaths, vPaths := TracePaths(instructions)
	for _, path := range hPaths {
		sort.Sort(path)
	}
	for _, path := range vPaths {
		sort.Sort(path)
	}
	println(hPaths, vPaths)

	// Find intersection points

	// Calculate manhattan distance for each intersection

	// Find intersection points closest to origin
}
