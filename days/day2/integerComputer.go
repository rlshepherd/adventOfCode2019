// Package main contains functions for solving IntComputer challenge
package main

import (
	"fmt"

	fileutils "github.com/rlshepherd/adventOfCode2019/utils"
)

var integers []int

func add(integers []int, pos int) {
	integers[integers[pos+3]] = integers[integers[pos+1]] + integers[integers[pos+2]]
}

func multiply(integers []int, pos int) {
	integers[integers[pos+3]] = integers[integers[pos+1]] * integers[integers[pos+2]]
}

func integerComputer(ints []int) {
Loop:
	for i := 0; i < len(ints); i += 4 {
		switch {
		case ints[i] == 1:
			add(ints, i)
		case ints[i] == 2:
			multiply(ints, i)
		case ints[i] == 99:
			break Loop
		}
	}
}

func main() {
	integers := fileutils.ReadIntegerFile("input.txt", ",")
	integers[1] = 12
	integers[2] = 2
	integerComputer(integers)
	fmt.Println("0th integer: ", integers[0])
}
