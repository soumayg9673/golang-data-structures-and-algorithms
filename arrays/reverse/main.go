package main

import (
	"fmt"
	"slices"
)

// naive approach - use temp array
func reverseArrNa(arr []int) []int {
	temp := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		temp = append(temp, arr[i])
	}
	return temp
}

func reverseArr(arr []int) []int {
	l := len(arr)
	for i := range l / 2 {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return arr
}

func main() {
	arr1 := []int{15}
	arr2 := []int{5, 18, 20}
	arr3 := []int{20, 20}
	arr4 := []int{20, 10}

	fmt.Println(reverseArrNa(arr1))
	fmt.Println(reverseArrNa(arr2))
	fmt.Println(reverseArrNa(arr3))
	fmt.Println(reverseArrNa(arr4))

	fmt.Println(reverseArr(arr1))
	fmt.Println(reverseArr(arr2))
	fmt.Println(reverseArr(arr3))
	fmt.Println(reverseArr(arr4))

	// use in-built function
	slices.Reverse(arr1)
	slices.Reverse(arr2)
	slices.Reverse(arr3)
	slices.Reverse(arr4)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

}
