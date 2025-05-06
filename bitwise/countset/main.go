package main

import "fmt"

func countSetNa(n int) int {
	total := 0
	for n > 0 {
		if n&1 == 1 { //explanation: n%2 == 1
			total++
		}
		n = n >> 1 //explanation: n = n/2
	}
	return total
}

func countSetEa(n int) int {
	total := 0
	for n > 0 {
		n = n & (n - 1)
		total++
	}
	return total
}

func initialize() [256]int {
	n := [256]int{}
	n[0] = 0
	for i := 1; i < 256; i++ {
		n[i] = n[i&(i-1)] + 1
	}
	return n
}
func countSet(n int, arr [256]int) int {
	return arr[n&255] + arr[(n>>8)&255] + arr[(n>>16)&255] + arr[(n>>24)]
}

func main() {
	// naive approach
	fmt.Println(countSetNa(5))    //output: 2
	fmt.Println(countSetNa(7))    //output: 3
	fmt.Println(countSetNa(13))   //output: 3
	fmt.Println(countSetNa(1200)) //output: 4

	// efficient approach
	fmt.Println(countSetEa(5))    //output: 2
	fmt.Println(countSetEa(7))    //output: 3
	fmt.Println(countSetEa(13))   //output: 3
	fmt.Println(countSetEa(1200)) //output: 4

	// optimal approach
	arr := initialize()
	fmt.Println(countSet(5, arr)) //output: 2
	arr = initialize()
	fmt.Println(countSet(7, arr)) //output: 3
	arr = initialize()
	fmt.Println(countSet(13, arr)) //output: 3
	arr = initialize()
	fmt.Println(countSet(1200, arr)) //output: 4
}
