package main

import "fmt"

func print1toN(n int) {
	if n == 0 {
		return
	}
	print1toN(n - 1)
	fmt.Print(n, ",")
}

func main() {
	print1toN(5)
}
