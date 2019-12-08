package main

import (
	"strconv"
	"strings"
)

func ascending(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func hasAtleastOne2Pair(a []int) bool {
	pairIndexes := make([]int, 0)
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			pairIndexes = append(pairIndexes, i)
		}
	}
	for i := 0; i < len(pairIndexes); i++ {
		l := len(a) - 2
		if pairIndexes[i] >= l && a[pairIndexes[i]] != a[pairIndexes[i]-1] {
			return true
		}
		if pairIndexes[i] > 0 {
			if a[pairIndexes[i]] != a[pairIndexes[i]-1] && a[pairIndexes[i]] != a[pairIndexes[i]+2] {
				return true
			}
		} else {
			if a[pairIndexes[i]] != a[pairIndexes[i]+2] {
				return true
			}
		}
	}
	return false
}

func intToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return intToSlice(n/10, sequence)
	}
	return sequence
}

func sliceToInt(a []int) (int, error) {
	valuesText := make([]string, 0)
	for i := range a {
		number := a[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}
	return strconv.Atoi(strings.Join(valuesText, ""))
}

func findPasswords(low, high int) []int {
	var passwords = make([]int, 0)
	var a = make([]int, 0)
	for i := low; i <= high; i++ {
		a = intToSlice(i, []int{})
		if ascending(a) && hasAtleastOne2Pair(a) {
			aInt, err := sliceToInt(a)
			if err == nil {
				passwords = append(passwords, aInt)
			}
		}
	}
	return passwords
}

func main() {
	passwords := findPasswords(264360, 746325)
	println("number of passwords: %d", len(passwords))
}
