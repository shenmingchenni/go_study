package main

import "fmt"

func main() {
	var nums = make([]int, 3, 5)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(nums), cap(nums), nums)
	nums = append(nums, 1)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(nums), cap(nums), nums)
	nums = append(nums, 2)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(nums), cap(nums), nums)
	nums = append(nums, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(nums), cap(nums), nums)

}
