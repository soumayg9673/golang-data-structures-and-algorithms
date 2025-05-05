package main

import "fmt"

func countDigits(n int) int {
	digits := 0
	for n > 0 {
		n /= 10
		digits += 1
	}
	return digits
}

func main() {
	num := 9235
	fmt.Println(countDigits((num)))
}
