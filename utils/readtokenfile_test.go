package fileutils

import (
	"fmt"
	"testing"
)

var testResult1 = [][]string{
	{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
	{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
}
var testResult2 = [][]string{
	{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
	{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
}

func TestReadTokenFile(t *testing.T) {
	cases := []struct {
		input string
		want  [][]string
	}{
		{"testData/readtokenfile_input_test1.txt", testResult1},
		{"testData/readtokenfile_input_test2.txt", testResult2},
	}
	for _, c := range cases {
		got := ReadTokenFile(c.input)
		for i, line := range got {
			if !EqualStringSplices(line, c.want[i]) {
				t.Errorf(fmt.Sprintf("ReadTokenFile(%q) == %q, want %q", c.input, got, c.want[i]))
			}
		}
	}
}
