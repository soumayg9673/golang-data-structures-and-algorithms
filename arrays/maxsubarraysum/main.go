package main

import (
	"fmt"
	"math"
)

func maxSumNa(arr []int) int {
	m := arr[0]
	for i := range arr {
		c := 0
		for j := i; j < len(arr); j++ {
			c += arr[j]
			m = int(math.Max(float64(m), float64(c)))
		}
	}
	return m
}

func maxSum(arr []int) int {
	tarr := []int{}
	tarr = append(tarr, arr[0])
	for i := 1; i < len(arr); i++ {
		ci := tarr[len(tarr)-1] + arr[i]
		tarr = append(tarr, int(math.Max(float64(arr[i]), float64(ci))))
	}
	m := tarr[0]
	for _, v := range tarr {
		m = int(math.Max(float64(v), float64(m)))
	}
	return m
}

func main() {
	arr1 := []int{2, 3, -8, 7, -1, 2, 3}
	arr2 := []int{5, 8, 3}
	arr3 := []int{-6, -1, -8}
	arr4 := []int{-5, 1, -2, 3, -1, 2, -2}

	fmt.Println("Naive Approach")
	fmt.Println(maxSumNa(arr1))
	fmt.Println(maxSumNa(arr2))
	fmt.Println(maxSumNa(arr3))
	fmt.Println(maxSumNa(arr4))

	fmt.Println("Optimal Approach")
	fmt.Println(maxSum(arr1))
	fmt.Println(maxSum(arr2))
	fmt.Println(maxSum(arr3))
	fmt.Println(maxSum(arr4))

}
