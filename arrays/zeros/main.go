package main

import "fmt"

func moveZeros1(arr []int) []int {
	for i := range arr {
		if arr[i] == 0 {
			for j := i + 1; j < len(arr); j++ {
				if arr[j] != 0 {
					arr[i], arr[j] = arr[j], arr[i]
					break
				}
			}
		}
	}
	return arr
}

func moveZeros(arr []int) []int {
	zi := -1
	for i := range arr {
		if arr[i] == 0 && zi == -1 {
			zi = i
		} else if zi != -1 && arr[i] != 0 {
			arr[i], arr[zi] = arr[zi], arr[i]
			zi++
		}
	}
	return arr
}

func main() {
	arr1 := []int{8, 5, 0, 10, 0, 20}
	arr2 := []int{0, 0, 0, 10, 0}
	arr3 := []int{10, 20}
	arr4 := []int{10, 20, 0, 15, 0, 30, 0, 0, 25}
	arr5 := []int{10, 0, 20}
	arr6 := []int{10, 20, 0, 30}

	fmt.Println(moveZeros1(arr1))
	fmt.Println(moveZeros1(arr2))
	fmt.Println(moveZeros1(arr3))
	fmt.Println(moveZeros1(arr4))

	fmt.Println(moveZeros(arr1))
	fmt.Println(moveZeros(arr2))
	fmt.Println(moveZeros(arr3))
	fmt.Println(moveZeros(arr4))
	fmt.Println(moveZeros(arr5))
	fmt.Println(moveZeros(arr6))
}
