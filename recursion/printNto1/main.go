package main

import "fmt"

func printNto1(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, ",")
	printNto1(n - 1)
}

func main() {
	printNto1(5)
}
