package main

import "fmt"

func naiveApproach(n int) []int {
	if n < 1 {
		return []int{}
	}
	arr := []int{1}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func optimalApproachNonSorted(n int) []int {
	if n < 1 {
		return []int{}
	}
	arr := []int{1, n}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
			if n/i != i {
				arr = append(arr, n/i)
			}
		}
	}
	return arr
}

func optimalApproachSorted(n int) []int {
	if n < 1 {
		return []int{}
	}
	arr := []int{}
	i := 1
	for ; i*i <= n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}
	for ; i >= 1; i-- {
		if n%i == 0 {
			arr = append(arr, n/i)
		}
	}
	return arr
}

func main() {
	// naive approach
	fmt.Println(naiveApproach(7))
	fmt.Println(naiveApproach(15))
	fmt.Println(naiveApproach(100))

	// efficient approach 1 - not sorted array
	fmt.Println(optimalApproachNonSorted(7))
	fmt.Println(optimalApproachNonSorted(15))
	fmt.Println(optimalApproachNonSorted(100))

	// efficient approach 2 - sorted array
	fmt.Println(optimalApproachSorted(7))
	fmt.Println(optimalApproachSorted(15))
	fmt.Println(optimalApproachSorted(100))
}
