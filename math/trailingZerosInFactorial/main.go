package main

import "fmt"

func FactorialRecursive(n int) int {
	if n > 0 {
		return n * FactorialRecursive(n-1)
	}
	return 1
}

func trailingZeroNa(n int) int {
	fact := FactorialRecursive(n)
	count := 0
	for fact%10 == 0 {
		count += 1
		fact /= 10
	}
	return count
}

func trailingZeroEa(n int) int {
	count := 0
	for i := 5; i <= n; i *= 5 {
		count += n / i
	}
	return count
}

func main() {
	// naive method - fails for higher value due to bits restrictions
	fmt.Println(trailingZeroNa(15))

	// efficient method
	fmt.Println(trailingZeroEa(15))
}
