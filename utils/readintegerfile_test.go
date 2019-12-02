package fileutils

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestReadIntegerFile(t *testing.T) {
	cases := []struct {
		file, seperator string
		want            []int
	}{
		{"test_newline.txt", "\n", []int{1, 2}},
		{"test_comma.txt", ",", []int{1, 2}},
	}
	for _, c := range cases {
		got := ReadIntegerFile(c.file, c.seperator)
		if !Equal(got, c.want) {
			t.Errorf(fmt.Sprintf("ReadIntegerFile(%q) == %q, want %d", c.file, got, c.want))
		}
	}
}
