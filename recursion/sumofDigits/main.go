package main

import "fmt"

func sumDigits(n int) int {
	if n <= 9 {
		return n
	}
	return n%10 + sumDigits(n/10)
}

func main() {
	fmt.Println(sumDigits(253))
	fmt.Println(sumDigits(10))
	fmt.Println(sumDigits(9987))
	fmt.Println(sumDigits(7))
}
