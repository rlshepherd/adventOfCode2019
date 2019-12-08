package main

import (
	"fmt"
	"testing"
)

func TestNewLineSegment(t *testing.T) {
	cases := []struct {
		point       Point
		instruction string
		steps       int
		want        LineSegment
	}{
		{Point{0, 0}, "U1", 0, LineSegment{Point{0, 0}, Point{0, 1}, true, 1}},
		{Point{1, 1}, "R1", 5, LineSegment{Point{1, 1}, Point{2, 1}, false, 6}},
		{Point{1, 1}, "D10", 1, LineSegment{Point{1, 1}, Point{1, -9}, true, 11}},
	}
	for _, c := range cases {
		got, err := NewLineSegment(c.point, c.instruction, c.steps)
		if err != nil {
			t.Errorf(fmt.Sprintf("Error: %w", err))
		}
		if got != c.want {
			t.Errorf(fmt.Sprintf("NewLineSegment(%d, %q) == %v, want %v", c.point, c.instruction, got, c.want))
		}
	}
}
