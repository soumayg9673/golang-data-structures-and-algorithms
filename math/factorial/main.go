package main

import "fmt"

func FactorialIterative(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	fac := 1
	for n > 0 {
		fac *= n
		n -= 1
	}
	return fac
}

func FactorialRecursive(n int) int {
	if n > 0 {
		return n * FactorialRecursive(n-1)
	}
	return 1
}

func main() {
	// using iterative approach
	fmt.Println(FactorialIterative(0))
	fmt.Println(FactorialIterative(3))
	fmt.Println(FactorialIterative(4))

	// using recursive approach
	fmt.Println(FactorialRecursive(0))
	fmt.Println(FactorialRecursive(3))
	fmt.Println(FactorialRecursive(4))
}
