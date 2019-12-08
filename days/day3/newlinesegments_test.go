package main

import (
	"fmt"
	"testing"
)

func TestNewLineSegment(t *testing.T) {
	cases := []struct {
		point       Point
		instruction string
		want        LineSegment
	}{
		{Point{0, 0}, "U1", LineSegment{Point{0, 0}, Point{0, 1}, true}},
		{Point{1, 1}, "R1", LineSegment{Point{1, 1}, Point{2, 1}, false}},
		{Point{1, 1}, "D10", LineSegment{Point{1, 1}, Point{1, -9}, true}},
	}
	for _, c := range cases {
		got, err := NewLineSegment(c.point, c.instruction)
		if err != nil {
			t.Errorf(fmt.Sprintf("Error: %w", err))
		}
		if got != c.want {
			t.Errorf(fmt.Sprintf("NewLineSegment(%d, %q) == %v, want %v", c.point, c.instruction, got, c.want))
		}
	}
}
