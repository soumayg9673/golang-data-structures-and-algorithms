package main

import "fmt"

func countRepeatElementsNa(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v] += 1
		if m[v] > 1 {
			return v
		}
	}
	return -1
}

func main() {
	arr1 := []int{0, 2, 1, 3, 2, 2}
	arr2 := []int{1, 2, 3, 0, 3, 4, 5}
	arr3 := []int{0, 0}

	fmt.Println(countRepeatElementsNa(arr1))
	fmt.Println(countRepeatElementsNa(arr2))
	fmt.Println(countRepeatElementsNa(arr3))
}
