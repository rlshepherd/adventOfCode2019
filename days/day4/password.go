package main

func ascending(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func hasPair(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			return true
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

func main() {
	var passwords = make([]int, 0)
	for i := 264360; i <= 746325; i++ {
		a := intToSlice(i, []int{})
		if ascending(a) && hasPair(a) {
			passwords = append(a, passwords...)
		}
	}
	println("number of passwords: %d", len(passwords))
}
