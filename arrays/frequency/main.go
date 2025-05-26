package main

import "fmt"

func frequency(arr []int) map[int]int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v] += 1
	}
	return m
}

func main() {
	arr1 := []int{10, 10, 10, 25, 30, 30}
	arr2 := []int{10, 10, 10, 10}
	arr3 := []int{10, 20, 30}

	fmt.Println(frequency(arr1))
	fmt.Println(frequency(arr2))
	fmt.Println(frequency(arr3))
}
