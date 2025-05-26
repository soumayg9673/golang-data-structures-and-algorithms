package main

import "fmt"

func insertInArray(arr []int, x, pos int) []int {
	pos-- // position in slices/array start from 0
	// insert last element at last
	if len(arr) != pos {
		arr = append(arr, arr[len(arr)-1])
	}
	for i := len(arr) - 1; i > pos; i-- {
		arr[i] = arr[i-1]
	}
	if len(arr) != pos {
		arr[pos] = x
	} else {
		arr = append(arr, x)
	}
	return arr
}

func main() {
	arr1 := []int{5, 15, 20}
	fmt.Println(insertInArray(arr1, 10, 2))

	arr2 := []int{1, 2, 3, 4}
	fmt.Println(insertInArray(arr2, 5, 5))

	arr3 := []int{1, 2, 3, 4}
	fmt.Println(insertInArray(arr3, 0, 1))
}
