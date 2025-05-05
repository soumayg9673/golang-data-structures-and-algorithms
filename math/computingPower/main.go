package main

import "fmt"

// time complexity: O(n)
func computingPower1(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	for i := 2; i < n; i++ {
		x *= x
	}
	return x
}

func computingPower(x, n int) int {
	if n == 0 {
		return 1
	}
	temp := computingPower(x, n/2)
	temp *= temp
	if n%2 == 0 {
		return temp
	} else {
		return temp * x
	}
}

func main() {
	fmt.Println("Naive Approach")
	fmt.Println(computingPower1(2, 3)) // 8
	fmt.Println(computingPower1(3, 4)) // 81
	fmt.Println(computingPower1(5, 0)) // 1
	fmt.Println(computingPower1(5, 1)) // 5

	fmt.Println("Efficient Approach")
	fmt.Println(computingPower(2, 3)) // 8
	fmt.Println(computingPower(3, 4)) // 81
	fmt.Println(computingPower(5, 0)) // 1
	fmt.Println(computingPower(5, 1)) // 5
}
