package main

import (
	"fmt"
	"sort"
	"testing"

	fileutils "github.com/rlshepherd/adventOfCode2019/utils"
)

func TestFindDistance(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"test_input1.txt", 610},
		{"test_input2.txt", 410},
		{"day3_part1_input.txt", 43258},
	}
	for _, c := range cases {
		instructions := fileutils.ParseIntrusctionFile(c.in)
		hPaths, vPaths := NewPathsByOrientation(instructions)
		p := make(Intersections, 0)
		p = append(p, FindIntersections(hPaths[0], vPaths[1])...)
		p = append(p, FindIntersections(hPaths[1], vPaths[0])...)
		sort.Sort(p)
		got := p[0].steps
		if got != c.want {
			t.Errorf(fmt.Sprintf("Steps = %d, want %d", got, c.want))
		}
	}
}
