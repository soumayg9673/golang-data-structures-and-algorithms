package main

import "fmt"

func sieveEratosthenesNa(n int) []int {
	if n <= 1 {
		return []int{}
	}
	arr := []int{}
	for i := 2; i <= n; i++ {
		for j := 2; j <= i; j++ {
			if j == i {
				arr = append(arr, i)
			} else if i%j == 0 {
				break
			}
		}
	}
	return arr
}

func createArray(n int) []bool {
	arr := []bool{}
	for i := 0; i <= n; i++ {
		arr = append(arr, true)
	}
	return arr
}

func sieveEratosthenesEa(n int) []int {
	if n <= 1 {
		return []int{}
	}
	isPrime := createArray(n)
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := 2 * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	arr := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			arr = append(arr, i)
		}
	}
	return arr
}

func sieveEratosthenesOa1(n int) []int {
	if n <= 1 {
		return []int{}
	}
	isPrime := createArray(n)
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	arr := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			arr = append(arr, i)
		}
	}
	return arr
}

// time complexity: O(n log log n)
func sieveEratosthenesOa2(n int) []int {
	if n <= 1 {
		return []int{}
	}
	isPrime := createArray(n)
	arr := []int{}

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			arr = append(arr, i)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return arr
}

func main() {
	// naive approach 1
	fmt.Println(sieveEratosthenesNa(10))
	fmt.Println(sieveEratosthenesNa(23))

	// using array of boolean
	fmt.Println(sieveEratosthenesEa(10))
	fmt.Println(sieveEratosthenesEa(23))

	// optimized approach
	fmt.Println(sieveEratosthenesOa1(10))
	fmt.Println(sieveEratosthenesOa1(23))

	// optimized approach
	fmt.Println(sieveEratosthenesOa2(10))
	fmt.Println(sieveEratosthenesOa2(23))
}
