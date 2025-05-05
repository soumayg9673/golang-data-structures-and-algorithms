package main

import "fmt"

func primeFactorsNa(n int) {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
			n /= i
			i = 1
		}
	}
}

func primeFactorsEa1(n int) {
	if n <= 1 {
		return
	}
	for i := 2; i*i <= n; i += 1 {
		for n%i == 0 {
			fmt.Println(i)
			n /= i
		}
	}
	if n > 1 {
		fmt.Println(n)
	}
}

func primeFactorEa2(n int) {
	if n <= 1 {
		return
	}
	for n%2 == 0 {
		fmt.Println(2)
		n /= 2
	}
	for n%3 == 0 {
		fmt.Println(3)
		n /= 3
	}
	for i := 5; i*i <= n; i += 6 {
		for n%i == 0 {
			fmt.Println(i)
			n /= i
		}
		for n%(i+2) == 0 {
			fmt.Println(i + 2)
			n /= i + 2
		}
	}
	if n > 3 {
		fmt.Println(n)
	}
}

func main() {
	// naive approach 1
	primeFactorsNa(12)  // 2,2,3
	primeFactorsNa(13)  // 13
	primeFactorsNa(315) //3,3,5,7

	// efficient approach 1
	primeFactorsEa1(12)  // 2,2,3
	primeFactorsEa1(13)  // 13
	primeFactorsEa1(315) //3,3,5,7

	// efficient approach 2
	primeFactorEa2(12)  // 2,2,3
	primeFactorEa2(13)  // 13
	primeFactorEa2(315) //3,3,5,7
}
