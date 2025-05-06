package main

import "fmt"

func binaryRep(n int) {
	if n == 0 {
		return
	}
	binaryRep(n / 2)
	fmt.Print(n % 2)
}

func main() {
	binaryRep(7)
	fmt.Println()
	binaryRep(8)
}
