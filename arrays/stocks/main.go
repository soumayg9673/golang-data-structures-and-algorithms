package main

import "fmt"

func maximizeProfits(arr []int) int {
	p := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			p += arr[i] - arr[i-1]
		}
	}
	return p
}

func main() {
	arr1 := []int{1, 5, 3, 8, 12}
	arr2 := []int{30, 20, 10}
	arr3 := []int{10, 20, 30}
	arr4 := []int{1, 5, 3, 1, 2, 8}

	fmt.Println(maximizeProfits(arr1))
	fmt.Println(maximizeProfits(arr2))
	fmt.Println(maximizeProfits(arr3))
	fmt.Println(maximizeProfits(arr4))
}
