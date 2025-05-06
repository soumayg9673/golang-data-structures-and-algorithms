package main

import (
	"fmt"
	"slices"
)

func getIndex(arr []int, n int) int {
	return slices.Index(arr, n)
}

func main() {
	arr := []int{20, 5, 7, 25}
	fmt.Println(getIndex(arr, 7))
	fmt.Println(getIndex(arr, 15))
}
