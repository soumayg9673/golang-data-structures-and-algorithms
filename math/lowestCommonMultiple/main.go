package main

import "fmt"

func lcm1(a, b int) int {
	// make a variable smaller than b
	if a > b {
		a, b = b, a
	}
	if b%a == 0 {
		return b
	}
	res := 1
	for i := 2; i <= a; i += 1 {
		if a%i == 0 && b%i == 0 {
			res *= i
			a /= i
			b /= i
			i = 2
		}
	}
	return res * a * b
}

func lcm2(a, b int) int {
	// make a variable smaller than b
	if a > b {
		a, b = b, a
	}
	res := 1
	for i := a; i > 1; i -= 1 {
		if a%i == 0 && b%i == 0 {
			res *= i
			a /= i
			b /= i
			i = a
		}
	}
	return res * a * b
}

func lcm3(a, b int) int {
	// make a variable smaller than b
	if a > b {
		a, b = b, a
	}
	res := b
	for true {
		if res%a == 0 && res%b == 0 {
			return res
		}
		res += 1
	}
	return res
}

func EuclideanAlgorithmRecusive(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return EuclideanAlgorithmRecusive(b, a%b)
	}
}

func lcm(a, b int) int {
	return (a * b) / EuclideanAlgorithmRecusive(a, b)
}

func main() {
	// naive approach 1
	fmt.Println(lcm1(4, 6))   // 12
	fmt.Println(lcm1(12, 15)) // 60
	fmt.Println(lcm1(2, 8))   // 8
	fmt.Println(lcm1(3, 7))   // 21

	// naive approach 2
	fmt.Println(lcm2(4, 6))   // 12
	fmt.Println(lcm2(12, 15)) // 60
	fmt.Println(lcm2(2, 8))   // 8
	fmt.Println(lcm2(3, 7))   // 21

	// naive approach 3
	fmt.Println(lcm3(4, 6))   // 12
	fmt.Println(lcm3(12, 15)) // 60
	fmt.Println(lcm3(2, 8))   // 8
	fmt.Println(lcm3(3, 7))   // 21

	// optimized solution
	fmt.Println(lcm(4, 6))   // 12
	fmt.Println(lcm(12, 15)) // 60
	fmt.Println(lcm(2, 8))   // 8
	fmt.Println(lcm(3, 7))   // 21
}
