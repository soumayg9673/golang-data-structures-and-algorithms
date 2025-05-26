package main

import (
	"fmt"
	"math"
)

func maximumWaterCapacity(arr []int) int {
	n := len(arr)
	c := 0

	lmax := []int{}
	rmax := []int{}

	lmax = append(lmax, arr...)
	rmax = append(rmax, arr...)

	lmax[0] = arr[0]
	for i := 1; i < n; i++ {
		lmax[i] = int(math.Max(float64(arr[i]), float64(lmax[i-1])))
	}
	rmax[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		rmax[i] = int(math.Max(float64(arr[i]), float64(rmax[i+1])))
	}

	for i := 1; i < n-1; i++ {
		c += int(math.Min(float64(lmax[i]), float64(rmax[i]))) - arr[i]
	}
	return c
}

func main() {
	arr1 := []int{2, 0, 2}
	arr2 := []int{3, 0, 1, 2, 5}
	arr3 := []int{3, 2, 1}
	arr4 := []int{1, 2, 5}
	arr5 := []int{5, 0, 6, 2, 3}

	fmt.Println(maximumWaterCapacity(arr1))
	fmt.Println(maximumWaterCapacity(arr2))
	fmt.Println(maximumWaterCapacity(arr3))
	fmt.Println(maximumWaterCapacity(arr4))
	fmt.Println(maximumWaterCapacity(arr5))
}
