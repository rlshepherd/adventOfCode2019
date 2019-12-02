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
