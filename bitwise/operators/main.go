package main

import "fmt"

func main() {
	x, y := 3, 6
	fmt.Println("AND: ", x&y) // output: 2
	fmt.Println("OR: ", x|y)  //output: 7
	fmt.Println("XOR: ", x^y) // output: 5
	x = 3
	fmt.Println("Left Shift: ", x<<1) //output: 6
	fmt.Println("Left Shift: ", x<<2) //output: 12
	x = 33
	fmt.Println("Right Shift: ", x>>1) //output: 16
	fmt.Println("Right Shift: ", x>>2) //output: 8
}
