package main

import (
	"fmt"
	"slices"
)

func findMedianSortedArraysNa(nums1 []int, nums2 []int) float64 {
	nums := slices.Concat(nums1, nums2)
	slices.Sort(nums)
	l := len(nums)
	if l%2 == 0 {
		s := nums[(l/2)-1] + nums[l/2]
		f := float64(s)
		return float64(f / 2)
	} else {
		return float64(nums[l/2])
	}
}

func main() {
	arr1 := []int{10, 20, 30, 40, 50}
	arr2 := []int{5, 15, 25, 35, 45}
	fmt.Println(findMedianSortedArraysNa(arr1, arr2))

	arr1 = []int{1, 2, 3, 4, 5, 6}
	arr2 = []int{10, 20, 30, 40, 50}
	fmt.Println(findMedianSortedArraysNa(arr1, arr2))

	arr1 = []int{10, 20, 30, 40, 50, 60}
	arr2 = []int{1, 2, 3, 4, 5}
	fmt.Println(findMedianSortedArraysNa(arr1, arr2))
}
