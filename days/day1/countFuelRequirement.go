//Package main contains functions for solving advent of code day 1
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

//ConvertMassToFuel returns an int divided by three, round down, subtract 2
func ConvertMassToFuel(m int) int {
	f := math.Floor(float64(m)/3.0) - 2.0
	return int(f)
}

//Sum retruns the total converted fuel requirements.
func Sum(a []string) (sum int) {
	for _, mass := range a {
		massInt, err := strconv.Atoi(mass)

		if err != nil {
			log.Fatal(err)
		}

		sum += ConvertMassToFuel(massInt)
	}
	return
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	weights := strings.Split(string(data), "\n")
	totalFuelRequirement := Sum(weights)
	fmt.Println("Total fuel requirement:", totalFuelRequirement)
}
