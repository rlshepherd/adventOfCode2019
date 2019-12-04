package fileutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// ReadIntegerFile reads file and returns a slice of integers.
func ReadIntegerFile(f string, sep string) (integers []int) {

	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	strings := strings.Split(string(data), sep)

	for _, j := range strings {
		k, err := strconv.Atoi(j)
		if err != nil {
			log.Fatal(err)
		}
		integers = append(integers, k)
	}

	return
}

// ReadTokenFile accepts a filename and returns LinesOfInstructions
func ReadTokenFile(f string) [][]string {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("File reading error", err)
	}

	lines := strings.Split(string(data), "\n")
	instructions := make([][]string, len(lines))

	for i, lineContents := range lines {
		tokens := strings.Split(lineContents, ",")
		instructions[i] = make([]string, len(tokens))
		for j, token := range tokens {
			instructions[i][j] = token
		}
	}
	return instructions
}

//Equal tests if two slices of ints have the same content
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

//EqualStringSplices tests if two slices of string have the same content
func EqualStringSplices(a, b []string) bool {
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
