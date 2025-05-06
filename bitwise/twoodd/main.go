package main

import "fmt"

func twoOddNa(arr []int) []int {
	m := map[int]int{}
	for _, v := range arr {
		m[v] = m[v] + 1
	}
	num := []int{}
	for k, v := range m {
		if v%2 == 1 {
			num = append(num, k)
		}
	}
	return num[:2]
}

func main() {
	fmt.Println(twoOddNa([]int{3, 4, 3, 4, 5, 4, 4, 6, 7, 7}))
	fmt.Println(twoOddNa([]int{1, 3, 2, 3, 3, 1}))
}
