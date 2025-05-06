package main

import "fmt"

func powerTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}

func main() {
	fmt.Println(powerTwo(4))  // true
	fmt.Println(powerTwo(6))  // false
	fmt.Println(powerTwo(1))  // true
	fmt.Println(powerTwo(0))  // false
	fmt.Println(powerTwo(12)) // false
	fmt.Println(powerTwo(16)) // true
	fmt.Println(powerTwo(11)) // false
}
