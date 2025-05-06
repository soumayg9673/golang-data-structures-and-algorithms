package main

import "fmt"

func checkKth(n, k int) bool {
	if k < 1 {
		return false
	}
	return (n>>(k-1))&1 == 1
}

func main() {
	fmt.Println(checkKth(5, 1))
	fmt.Println(checkKth(8, 2))
	fmt.Println(checkKth(0, 3))
}
