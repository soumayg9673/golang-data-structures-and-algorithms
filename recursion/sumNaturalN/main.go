package main

import "fmt"

func sumNaturalNo(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumNaturalNo(n-1)
}

func main() {
	fmt.Println(sumNaturalNo(4))
	fmt.Println(sumNaturalNo(5))
}
