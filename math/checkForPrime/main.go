package main

import "fmt"

func checkPrimeNa1(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func checkPrimeNa2(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func checkPrimeEa(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return false
	}
	for i := 2; i*i < n; i += 1 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i < n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	// naive approach 1
	fmt.Println(checkPrimeNa1(1))
	fmt.Println(checkPrimeNa1(3))
	fmt.Println(checkPrimeNa1(41))

	// naive approach 2
	fmt.Println(checkPrimeNa2(1))
	fmt.Println(checkPrimeNa2(3))
	fmt.Println(checkPrimeNa2(41))

	// efficient approach
	fmt.Println(checkPrimeEa(1))
	fmt.Println(checkPrimeEa(3))
	fmt.Println(checkPrimeEa(41))

	// optimized approach
	fmt.Println(checkPrime(1))
	fmt.Println(checkPrime(3))
	fmt.Println(checkPrime(41))
}
