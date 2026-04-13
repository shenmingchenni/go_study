package main

import "fmt"

func main() {
	//slice1 := []int{1, 2, 3}

	//var slice1 []int
	//slice1 = make([]int, 3)

	var slice1 []int = make([]int, 3)

	//slice1 := make([]int, 3)
	fmt.Printf("len=%d, slice=%v", len(slice1), slice1)
}
