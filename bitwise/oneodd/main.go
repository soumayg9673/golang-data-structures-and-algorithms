package main

import "fmt"

func oneOddNa(arr []int) int {
	m := map[int]int{}
	for _, v := range arr {
		m[v] = m[v] + 1
	}
	for k, v := range m {
		if v%2 == 1 {
			return k
		}
	}
	return 0
}

func OneOdd(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res = res ^ arr[i]
	}
	return res
}

func main() {
	// naive solution
	fmt.Println(oneOddNa([]int{4, 3, 4, 4, 4, 5, 5}))
	fmt.Println(oneOddNa([]int{8, 7, 7, 8, 8}))

	// optimal solution
	fmt.Println(OneOdd([]int{4, 3, 4, 4, 4, 5, 5}))
	fmt.Println(OneOdd([]int{8, 7, 7, 8, 8}))
}
