package main

import "fmt"

func main() {

	a1 := [5]int{} // fixed size array
	fmt.Println(a1, len(a1))

	a2 := [5]int{0: 1, 2: 3} // fixed size array and initialize indexes with values
	fmt.Println(a2, len(a2))

	a3 := [5]int{1, 2, 3, 4, 5} // fixed size array with values
	fmt.Println(a3, len(a3))

	a4 := make([]int, 5)
	fmt.Println(a4, len(a4))

	a5 := []int{}
	fmt.Println(a5, len(a5))

	a6 := []int{1, 2}
	fmt.Println(a6, len(a6))

	a7 := make([]int, 3)
	fmt.Println(a7, len(a7))

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	myslice = arr1[4:]
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

}
