package main

import "fmt"

func leftShiftOne(arr []int) []int {
	e := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1], arr[i] = arr[i], arr[i-1]
	}
	arr[len(arr)-1] = e
	return arr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{30, 5, 20}

	fmt.Println(leftShiftOne(arr1))
	fmt.Println(leftShiftOne(arr2))
}
