package main

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

func TestIntegerComputer(t *testing.T) {
	cases := []struct {
		input, want []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, c := range cases {
		thisInput := c.input
		integerComputer(thisInput)
		got := thisInput
		if !Equal(got, c.want) {
			t.Errorf(fmt.Sprintf("integerComputer(%d) == %d, want %d", c.input, got, c.want))
		}
	}
}
