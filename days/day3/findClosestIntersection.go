package main

import (
	"sort"

	fileutils "github.com/rlshepherd/adventOfCode2019/utils"
)

func main() {

	// Read instructions
	instructions := fileutils.ParseIntrusctionFile("day3_part1_input.txt")

	// Create empty paths
	hPaths, vPaths := NewPathsByOrientation(instructions)
	// for _, path := range hPaths {
	// 	sort.Sort(path)
	// }
	// for _, path := range vPaths {
	// 	sort.Sort(path)
	// }

	// Find intersection points
	p := make(Points, 0)
	p = append(p, FindIntersections(hPaths[0], vPaths[1])...)
	p = append(p, FindIntersections(hPaths[1], vPaths[0])...)

	sort.Sort(p)
	println("Closest point: %w", p[0].manhattanDistance())
}
