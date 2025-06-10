package main

import "fmt"

/*
	Stable algorithm
	For ascending order: arr[j] > arr[j+1]
	For descending order: arr[j] < arr[j+1]
*/

func bubbleSortNa(arr []int) []int {
	for i := range len(arr) - 1 {
		for j := range len(arr) - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func bubbleSort(arr []int) []int {
	for i := range len(arr) - 1 {
		swp := false
		for j := range len(arr) - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swp = true
			}
		}
		if !swp {
			break
		}
	}
	return arr
}

func main() {
	arr1 := []int{2, 10, 8, 7}
	arr2 := []int{1, 2, 3, 4}

	fmt.Println(bubbleSortNa(arr1))
	fmt.Println(bubbleSort(arr2))
}
