package main

import "fmt"

func greatestCommonDivisor(a int, b int) int {
	if a == 1 || b == 1 {
		return 1
	}
	if a > b {
		a, b = b, a // keep a as smallest number
	}
	if b%a == 0 {
		return a
	}
	for i := a; i >= 1; i -= 1 {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

func EuclideanAlgorithm(a int, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func EuclideanAlgorithm1(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return EuclideanAlgorithm1(b, a%b)
	}
}

func main() {
	// naive approach
	fmt.Println(greatestCommonDivisor(4, 6))
	fmt.Println(greatestCommonDivisor(100, 200))

	// interative approach
	fmt.Println(EuclideanAlgorithm(4, 6))
	fmt.Println(EuclideanAlgorithm(100, 200))

	// recursive approach
	fmt.Println(EuclideanAlgorithm1(4, 6))
	fmt.Println(EuclideanAlgorithm1(100, 200))
}
