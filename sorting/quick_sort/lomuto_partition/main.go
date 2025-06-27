package main

import "fmt"

/*
	Considering pivot the last element of the array
*/

func lomutoParition(arr []int, l, h int) {
	p := arr[h]
	i := l - 1
	for j := l; j <= h-1; j++ {
		if arr[j] < p {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[h] = arr[h], arr[i+1]
	fmt.Println(arr)
}

func main() {
	arr1 := []int{10, 80, 30, 90, 40, 50, 70}

	lomutoParition(arr1, 0, len(arr1)-1)
}
