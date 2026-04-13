package main

import "fmt"

func printmarry(marry []int) {
	for _, value := range marry {
		fmt.Println("value=", value)
	}
	marry[0] = 100
}

func main() {
	marry := []int{1, 2, 3, 4}
	printmarry(marry)

	fmt.Println("........")

	for _, value := range marry {
		fmt.Println("value=", value)
	}
}
