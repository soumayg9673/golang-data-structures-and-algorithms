package main

import "fmt"

func largestElement(arr []int) int {
	l := arr[0]
	for _, v := range arr {
		if v > l {
			l = v
		}
	}
	return l
}

func main() {
	arr1 := []int{15, 5, 20}
	fmt.Println(largestElement(arr1))
}
