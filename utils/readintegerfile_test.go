package fileutils

import (
	"fmt"
	"testing"
)

func TestReadIntegerFile(t *testing.T) {
	cases := []struct {
		file, seperator string
		want            []int
	}{
		{"testData/readintegerfile_test_newline.txt", "\n", []int{1, 2}},
		{"testData/readintegerfile_test_comma.txt", ",", []int{1, 2}},
	}
	for _, c := range cases {
		got := ReadIntegerFile(c.file, c.seperator)
		if !Equal(got, c.want) {
			t.Errorf(fmt.Sprintf("ReadIntegerFile(%q) == %q, want %d", c.file, got, c.want))
		}
	}
}
