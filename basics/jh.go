package main

import "fmt"

func swag(a *int, b *int) {
	var temp int = 0
	temp = *a
	*a = *b
	*b = temp

}

func main() {
	var a, b int = 3, 4
	swag(&a, &b)
	fmt.Println("a=", a, "b=", b)
}
