package main

import "fmt"

func checkTripletNa(arr []int, x int) bool {
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == x {
					return true
				}
			}
		}
	}
	return false
}

func checkPair(arr []int, x int) bool {
	l, h := 0, len(arr)-1
	for l < h {
		if arr[l]+arr[h] == x {
			return true
		} else if arr[l]+arr[h] > x {
			h--
		} else {
			l++
		}
	}
	return false
}

func checkTriplet(arr []int, x int) bool {
	for i := range len(arr) - 3 {
		if checkPair(arr[i+1:], x-arr[i]) {
			return true
		}
	}
	return false
}

func main() {
	arr1 := []int{2, 3, 4, 8, 9, 20, 40}
	arr2 := []int{1, 2, 5, 6}

	fmt.Println("Naive Approach")
	fmt.Println(checkTripletNa(arr1, 32))
	fmt.Println(checkTripletNa(arr2, 14))

	fmt.Println(checkTriplet(arr1, 32))
	fmt.Println(checkTriplet(arr2, 14))
}
